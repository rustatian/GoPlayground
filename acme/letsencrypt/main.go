package main

import (
	"github.com/caddyserver/certmagic"
)

func main() {
	cache := certmagic.NewCache(certmagic.CacheOptions{
		GetConfigForCert: func(c certmagic.Certificate) (*certmagic.Config, error) {
			return &certmagic.Config{
				RenewalWindowRatio: 0,
				OnEvent:            nil,
				DefaultServerName:  "https",
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

	cfg := certmagic.New(cache, certmagic.Config{})

	myAcme := certmagic.NewACMEManager(cfg, certmagic.ACMEManager{
		CA:                      certmagic.LetsEncryptStagingCA,
		TestCA:                  certmagic.LetsEncryptStagingCA,
		Email:                   "govnomulo@mail.ru",
		Agreed:                  true,
		DisableHTTPChallenge:    false,
		DisableTLSALPNChallenge: false,
		ListenHost:              "localhost",
		AltHTTPPort:             0,
		AltTLSALPNPort:          0,
		DNS01Solver:             nil,
		TrustedRoots:            nil,
		CertObtainTimeout:       0,
		Resolver:                "",
		PreferredChains:         certmagic.ChainPreference{},
	})

	cfg.Issuers = append(cfg.Issuers, myAcme)

	err := cfg.ManageSync([]string{"testtesttest.club"})
	if err != nil {
		panic(err)
	}

	tlsConfig := cfg.TLSConfig()
	_ = tlsConfig
}
