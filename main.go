package main

import (
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/plugins"
	"github.com/teutat3s/nomad-driver-triton/plugin"
)

func main() {
	plugins.Serve(factory)
}

// factory returns a new instance of the IIS Driver plugin
func factory(log hclog.Logger) interface{} {
	return plugin.NewDriver(log)
}
