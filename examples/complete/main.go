package main

import (
	routercmd "github.com/wundergraph/cosmo/router/cmd"
	// Import your modules here
	_ "github.com/wundergraph/router-examples/complete/myModule"
)

func main() {
	routercmd.Main()
}
