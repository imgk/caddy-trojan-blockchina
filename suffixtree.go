package blockchina

import "strings"

type node struct {
	next map[string]*node
}

func (nd *node) Put(domain string) {
	const sep = "."

	nd.store(strings.Split(strings.TrimSuffix(domain, sep), sep))
}

func (nd *node) store(ks []string) {
	l := len(ks)
	switch l {
	case 0:
		return
	case 1:
		k := ks[l-1]

		if k == "**" {
			nd.next[k] = (*node)(nil)
			return
		}

		_, ok := nd.next[k]
		if ok {
			return
		}

		nd.next[k] = &node{
			next: map[string]*node{},
		}
	default:
		k := ks[l-1]

		b, ok := nd.next[k]
		if !ok {
			b = &node{
				next: map[string]*node{},
			}
			nd.next[k] = b
		}

		b.store(ks[:l-1])
	}
}

func (nd *node) Get(domain string) bool {
	const sep = "."

	return nd.load(strings.Split(strings.TrimSuffix(domain, sep), sep))
}

func (nd *node) load(ks []string) bool {
	l := len(ks)
	switch l {
	case 0:
		return false
	case 1:
		_, ok := nd.next[ks[l-1]]
		if ok {
			return true
		}

		_, ok = nd.next["*"]
		if ok {
			return true
		}

		_, ok = nd.next["**"]
		if ok {
			return true
		}

		return false
	default:
		b, ok := nd.next[ks[l-1]]
		if ok {
			return b.load(ks[:l-1])
		}

		b, ok = nd.next["*"]
		if ok {
			return b.load(ks[:l-1])
		}

		_, ok = nd.next["**"]
		if ok {
			return true
		}

		return false
	}
}
