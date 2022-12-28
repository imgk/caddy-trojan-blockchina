package blockchina

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/caddyserver/caddy/v2"
)

// Dialer is ...
type Dialer struct {
	List []string `json:"block_list,omitempty"`

	sync.RWMutex
	node node
}

// Provision is ...
func (d *Dialer) Provision(ctx caddy.Context) error {
	d.node.next = map[string]*node{}
	for _, domain := range d.List {
		d.node.Put(domain)
	}
	return nil
}

// Dial is ...
func (d *Dialer) Dial(network, addr string) (net.Conn, error) {
	address, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, fmt.Errorf("dial %s error: %w", addr, err)
	}
	if d.Filt(address) {
		return nil, errors.New("blocked domain")
	}
	return net.Dial(network, addr)
}

// ListenPacket is ...
func (d *Dialer) ListenPacket(network, addr string) (net.PacketConn, error) {
	address, _, err := net.SplitHostPort(addr)
	if err != nil {
		if addr != "" {
			return nil, err
		}
		return net.ListenPacket(network, addr)
	}
	if d.Filt(address) {
		return nil, errors.New("blocked domain")
	}
	return net.ListenPacket(network, addr)
}

// Filt is ...
func (d *Dialer) Filt(domain string) bool {
	d.RLock()
	defer d.RUnlock()

	return d.node.Get(domain)
}
