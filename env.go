package env

import "github.com/coredns/coredns/plugin"

type EnvPlugin struct {
	Next  plugin.Handler
	Paths []string
}

func (b *EnvPlugin) Name() string { return "env" }
