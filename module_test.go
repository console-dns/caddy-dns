package caddy_dns

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

func TestUnmarshalCaddyfile(t *testing.T) {
	provider := Provider{}.CaddyModule().New().(*Provider)
	dispenser := caddyfile.NewTestDispenser(`dns console {
		server https://dns.example.com
		token secret
	}`)

	if err := provider.UnmarshalCaddyfile(dispenser); err != nil {
		t.Fatal(err)
	}
	if provider.Server != "https://dns.example.com" || provider.Token != "secret" {
		t.Fatalf("unexpected provider config: server=%q token=%q", provider.Server, provider.Token)
	}
}

func TestUnmarshalCaddyfileRequiresCredentials(t *testing.T) {
	provider := Provider{}.CaddyModule().New().(*Provider)
	dispenser := caddyfile.NewTestDispenser(`dns console {
		server https://dns.example.com
	}`)

	if err := provider.UnmarshalCaddyfile(dispenser); err == nil {
		t.Fatal("expected missing token error")
	}
}
