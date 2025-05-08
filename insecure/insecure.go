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

}

func loadFileFromFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func Load() error {
	var err error
	if len(certPEM) == 0 || len(keyPEM) == 0 {
		log.Printf("Loading certificate and key from file %s and %s", os.Getenv("PLATFORM_CERT_FILE"), os.Getenv("PLATFORM_KEY_FILE"))
		certPEM, err = loadFileFromFile(os.Getenv("PLATFORM_CERT_FILE"))
		if err != nil {
			return err
		}
		keyPEM, err = loadFileFromFile(os.Getenv("PLATFORM_KEY_FILE"))
		if err != nil {
			return err
		}
	}
	Cert, err = tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
	if err != nil {
		return err
	}
	Cert.Leaf, err = x509.ParseCertificate(Cert.Certificate[0])
	if err != nil {
		return err
	}

	CertPool = x509.NewCertPool()
	CertPool.AddCert(Cert.Leaf)
	return nil
}
