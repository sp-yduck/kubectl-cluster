package kubeconfig

import (
	"os"
	"testing"
)

func TestGetRawConfig(t *testing.T) {
	os.Setenv("KUBECONFIG", "./test/data/dummy-config.yaml")
	config, err := getRawConfig()
	if err != nil {
		t.Fatalf("failed to get kubeconfig: %v", err)
	}
	t.Logf("get kubeconfig: %v", config)
}

