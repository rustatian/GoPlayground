package main

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/spiral/errors"
	"golang.org/x/crypto/acme"
)

const (
	leProductionURL        = "https://acme-v02.api.letsencrypt.org/directory"
	leStagingURL    string = "https://acme-staging-v02.api.letsencrypt.org/directory"

	keyLen   int    = 4096
	keyType  string = "RSA PRIVATE KEY"
	certType string = "CERTIFICATE"
)

type AcmeConfig struct {
	// directory to save the certificates
	Dir string `mapstructure:"dir"`
}

func newKey(filename string) (crypto.Signer, error) {
	const op = errors.Op("acme_generate_new_key")

	k, err := rsa.GenerateKey(rand.Reader, keyLen)
	if err != nil {
		return nil, errors.E(op, err)
	}

	b := pem.EncodeToMemory(&pem.Block{
		Type:  keyType,
		Bytes: x509.MarshalPKCS1PrivateKey(k),
	})

	err = os.WriteFile(filename, b, 0600)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return k, nil
}

func main() {
	tmpDir := os.TempDir()

	// generate account key
	accKey, err := newKey(path.Join(tmpDir, "account.pem"))
	if err != nil {
		panic(err)
	}

	client := &acme.Client{
		Key:          accKey,
		DirectoryURL: leStagingURL,
	}

	_, err = client.Register(context.Background(), &acme.Account{
		URI:                    "",
		Contact:                []string{"govnomulo@mail.ru"},
		Status:                 "",
		OrdersURL:              "",
		AgreedTerms:            "",
		CurrentTerms:           "",
		Authz:                  "",
		Authorizations:         "",
		Certificates:           "",
		ExternalAccountBinding: nil,
	}, acme.AcceptTOS)
	if err != nil {
		panic(err)
	}

	// https://datatracker.ietf.org/doc/html/rfc8555#section-7.4.1
	auth, err := client.Authorize(context.Background(), "rustatian.me")
	if err != nil {
		panic(err)
	}

	var challenge *acme.Challenge
	for k := range auth.Challenges {
		if auth.Challenges[k].Type == "http-01" {
			challenge = auth.Challenges[k]
		}
	}

	if challenge == nil {
		panic(errors.Str("no http-01 challenge found"))
	}

	// "http-01", "tls-alpn-01", "dns-01".
	challenge.Token = "test-test"
	challengePath := client.HTTP01ChallengePath(challenge.Token)
	challengeResp, err := client.HTTP01ChallengeResponse(challenge.Token)
	if err != nil {
		panic(err)
	}

	go func() {
		serv := http.NewServeMux()
		serv.HandleFunc(challengePath, func(w http.ResponseWriter, r *http.Request) {
			b := []byte(challengeResp)
			w.Header().Set("Content-Length", strconv.Itoa(len(b)))
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(b)
		})

		l, err := net.Listen("tcp", "0.0.0.0:80")
		if err != nil {
			panic(err)
		}

		defer l.Close()

		http.Serve(l, serv)

		// LISTENER SHOULD BE HERE
	}()

	_, err = client.Accept(context.Background(), challenge)
	if err != nil {
		panic(err)
	}

	_, err = client.WaitAuthorization(context.Background(), auth.URI)
	if err != nil {
		panic(err)
	}

	uDomain := strings.ReplaceAll("rustatian.me", ".", "_")
	domainKey, err := newKey(fmt.Sprintf("%s.key", path.Join(tmpDir, uDomain)))
	if err != nil {
		panic(err)
	}

	csr, err := x509.CreateCertificateRequest(rand.Reader, &x509.CertificateRequest{
		Subject: pkix.Name{CommonName: "rustatian.me"},
	}, domainKey)

	if err != nil {
		panic(err)
	}

	domainCert, _, err := client.CreateCert(context.Background(), csr, 90*24*time.Hour, true)
	if err != nil {
		panic(err)
	}

	w, err := os.Create(path.Join(tmpDir, fmt.Sprintf("%s.crt", uDomain)))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = w.Close()
	}()

	for k := range domainCert {
		err = pem.Encode(w, &pem.Block{
			Type:  certType,
			Bytes: domainCert[k],
		})
		if err != nil {
			panic(err)
		}
	}


	fmt.Println("[------] DONE")
}
