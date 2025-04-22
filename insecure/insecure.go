package insecure

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"

	_ "embed"
)

// // go:embed platform.crt
var certPEM []byte

// //go:embed platform.key
var keyPEM []byte

var (
	// Cert is a self signed certificate
	Cert tls.Certificate
	// CertPool contains the self signed certificate
	CertPool *x509.CertPool
)

func init() {
	var err error
	if len(certPEM) == 0 || len(keyPEM) == 0 {
		certPEM, err = loadFileFromFile(os.Getenv("PLATFORM_CERT_FILE"))
		if err != nil {
			log.Fatalln("Failed to load certificate:", err)
		}
		keyPEM, err = loadFileFromFile(os.Getenv("PLATFORM_KEY_FILE"))
		if err != nil {
			log.Fatalln("Failed to load key:", err)
		}
	}
	Cert, err = tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
	if err != nil {
		log.Fatalln("Failed to parse key pair:", err)
	}
	Cert.Leaf, err = x509.ParseCertificate(Cert.Certificate[0])
	if err != nil {
		log.Fatalln("Failed to parse certificate:", err)
	}

	CertPool = x509.NewCertPool()
	CertPool.AddCert(Cert.Leaf)
}

func loadFileFromFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}
