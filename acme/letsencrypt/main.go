package main

import (
	"context"

	"github.com/caddyserver/certmagic"
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
		CA:                      certmagic.LetsEncryptStagingCA,
		TestCA:                  certmagic.LetsEncryptStagingCA,
		Email:                   "govnomulo@mail.ru",
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

	err := cfg.ObtainCertAsync(context.Background(), "testtesttest.club")
	if err != nil {
		panic(err)
	}
}
