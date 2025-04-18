package caddy_dns

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/console-dns/client"
	"github.com/console-dns/libdns"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *libdns.ConsoleDnsProvider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "dns.providers.console",
		New: func() caddy.Module {
			return &Provider{&libdns.ConsoleDnsProvider{
				ConsoleDnsClient: &client.ConsoleDnsClient{
					Server: "",
					Token:  "",
				},
			}}
		},
	}
}

func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "server":
				if d.NextArg() {
					p.Server = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "token":
				if d.NextArg() {
					p.Token = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Token == "" || p.Server == "" {
		return d.Err("missing Server or API token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
)
