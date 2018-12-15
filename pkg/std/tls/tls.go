package tls

import (
	"crypto/tls"
	"log"
)

func LoadCertif(certFile, keyFile string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load certificate pair: %s", err)
	}
	return cert
}
