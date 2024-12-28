package main

import (
	"github.com/loft-sh/vcluster-pod-hooks/hooks"
	"github.com/nirvati/vcluster-sdk/plugin"
)

func main() {
	_ = plugin.MustInit()
	plugin.MustRegister(hooks.NewPodHook())
	plugin.MustRegister(hooks.NewServiceHook())
	plugin.MustRegister(hooks.NewSecretHook())
	plugin.MustStart()
}
