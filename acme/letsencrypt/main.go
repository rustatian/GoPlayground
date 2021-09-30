package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/caddyserver/certmagic"
	"golang.org/x/sys/cpu"
)

func main() {
	cache := certmagic.NewCache(certmagic.CacheOptions{
		GetConfigForCert: func(c certmagic.Certificate) (*certmagic.Config, error) {
			return &certmagic.Config{
				RenewalWindowRatio: 0,
				OnEvent:            nil,
				OnDemand:           nil,
				MustStaple:         false,
				KeySource:          nil,
				CertSelection:      nil,
				OCSP:               certmagic.OCSPConfig{},
				Storage:            &certmagic.FileStorage{Path: "rr_le_certs"},
				Logger:             nil,
			}, nil
		},
		OCSPCheckInterval:  0,
		RenewCheckInterval: 0,
		Capacity:           0,
	})

	cfg := certmagic.New(cache, certmagic.Config{
		RenewalWindowRatio: 0,
		OnEvent:            nil,
		OnDemand:           nil,
		MustStaple:         false,
		Issuers:            nil,
		KeySource:          nil,
		CertSelection:      nil,
		OCSP:               certmagic.OCSPConfig{},
		Storage:            &certmagic.FileStorage{Path: "rr_le_certs"},
	})

	myAcme := certmagic.NewACMEManager(cfg, certmagic.ACMEManager{
		CA:                      certmagic.LetsEncryptProductionCA,
		TestCA:                  certmagic.LetsEncryptStagingCA,
		Email:                   "piashchynski.valery@gmail.com",
		Agreed:                  true,
		DisableHTTPChallenge:    false,
		DisableTLSALPNChallenge: false,
		ListenHost:              "",
		AltHTTPPort:             0,
		AltTLSALPNPort:          0,
		TrustedRoots:            nil,
		CertObtainTimeout:       0,
		Resolver:                "",
		NewAccountFunc:          nil,
		PreferredChains:         certmagic.ChainPreference{},
		Logger:                  nil,
	})

	cfg.Issuers = append(cfg.Issuers, myAcme)

	err := cfg.ObtainCertAsync(context.Background(), "piashchynski.valery@gmail.com")
	if err != nil {
		panic(err)
	}

	err = cfg.ManageSync([]string{"testtesttest.club"})
	if err != nil {
		panic(err)
	}

	var topCipherSuites []uint16
	var defaultCipherSuitesTLS13 []uint16

	hasGCMAsmAMD64 := cpu.X86.HasAES && cpu.X86.HasPCLMULQDQ
	hasGCMAsmARM64 := cpu.ARM64.HasAES && cpu.ARM64.HasPMULL
	// Keep in sync with crypto/aes/cipher_s390x.go.
	hasGCMAsmS390X := cpu.S390X.HasAES && cpu.S390X.HasAESCBC && cpu.S390X.HasAESCTR && (cpu.S390X.HasGHASH || cpu.S390X.HasAESGCM)

	hasGCMAsm := hasGCMAsmAMD64 || hasGCMAsmARM64 || hasGCMAsmS390X

	if hasGCMAsm {
		// If AES-GCM hardware is provided then priorities AES-GCM
		// cipher suites.
		topCipherSuites = []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		}
		defaultCipherSuitesTLS13 = []uint16{
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
		}
	} else {
		// Without AES-GCM hardware, we put the ChaCha20-Poly1305
		// cipher suites first.
		topCipherSuites = []uint16{
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		}
		defaultCipherSuitesTLS13 = []uint16{
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
		}
	}

	DefaultCipherSuites := make([]uint16, 0, 22)
	DefaultCipherSuites = append(DefaultCipherSuites, topCipherSuites...)
	DefaultCipherSuites = append(DefaultCipherSuites, defaultCipherSuitesTLS13...)

	sslServer := &http.Server{
		Addr: "0.0.0.0:443",
		TLSConfig: &tls.Config{
			CurvePreferences: []tls.CurveID{
				tls.X25519,
				tls.CurveP256,
				tls.CurveP384,
				tls.CurveP521,
			},
			CipherSuites:             DefaultCipherSuites,
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
			GetCertificate:           cfg.GetCertificate,
		},
	}

	log.Fatal(sslServer.ListenAndServeTLS("", ""))
}
