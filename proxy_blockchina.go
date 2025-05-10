package blockchina

import (
	"io"

	"github.com/caddyserver/caddy/v2"

	"github.com/imgk/caddy-trojan/app"
	"github.com/imgk/caddy-trojan/trojan"
)

func init() {
	caddy.RegisterModule((*BlockChina)(nil))
}

// BlockChina is ...
type BlockChina struct {
	Dialer
}

// CaddyModule is ...
func (*BlockChina) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "trojan.proxy.block_china",
		New: func() caddy.Module { return new(BlockChina) },
	}
}

// Close is ...
func (p *BlockChina) Close() error {
	return nil
}

// Handle is ...
func (p *BlockChina) Handle(r io.Reader, w io.Writer) (int64, int64, error) {
	return trojan.HandleWithDialer(r, w, &p.Dialer)
}

var (
	_ app.Proxy         = (*BlockChina)(nil)
	_ caddy.Provisioner = (*BlockChina)(nil)
)
