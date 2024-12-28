package main

import (
	"github.com/loft-sh/vcluster-example/syncers"
	"github.com/nirvati/vcluster-sdk/plugin"
)

func main() {
	ctx := plugin.MustInit()
	plugin.MustRegister(syncers.NewCarSyncer(ctx))
	plugin.MustStart()
}
