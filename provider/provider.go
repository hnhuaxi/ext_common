package provider

import "github.com/hysios/utils"

type Provider[Iface any, C any, A Ctor[Iface, C]] struct {
	registers utils.Map[string, A]
}

type Ctor[Iface any, C any] func(C) Iface

func (p *Provider[Iface, C, A]) Register(name string, ctor A) {
	p.registers.Store(name, ctor)
}

func (p *Provider[Iface, C, A]) Unregister(name string) {
	p.registers.Delete(name)
}

// Lookup returns the provider registered with the given name.
func (p *Provider[Iface, C, A]) Lookup(name string) (A, bool) {
	return p.registers.Load(name)
}
