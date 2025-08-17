// Package executor — executor internal logic
package executor

import (
	"context"
	"net/http"
	"testing"

	"github.com/voltgate/voltgate/v6/internal/config"
	voltgateauth "github.com/voltgate/voltgate/v6/sdk/voltgate/auth"
	sdkconfig "github.com/voltgate/voltgate/v6/sdk/config"
)

func TestNewProxyAwareHTTPClientDirectBypassesGlobalProxy(t *testing.T) {
	t.Parallel()

	client := newProxyAwareHTTPClient(
		context.Background(),
		&config.Config{SDKConfig: sdkconfig.SDKConfig{ProxyURL: "http://global-proxy.example.com:8080"}},
		&voltgateauth.Auth{ProxyURL: "direct"},
		0,
	)

	transport, ok := client.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("transport type = %T, want *http.Transport", client.Transport)
	}
	if transport.Proxy != nil {
		t.Fatal("expected direct transport to disable proxy function")
	}
}
