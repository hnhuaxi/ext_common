package insecure

import "testing"

func TestLoadFileFromFile(t *testing.T) {
	t.Logf("certPEM: %s", string(certPEM))
	t.Logf("keyPEM: %s", string(keyPEM))
}
