package main

import (
	// Register builtin memory and redis engines.
	_ "github.com/stone-payments/centrifugo/libcentrifugo/engine/enginememory"
	_ "github.com/stone-payments/centrifugo/libcentrifugo/engine/engineredis"

	// Register embedded web interface.
	_ "github.com/stone-payments/centrifugo/libcentrifugo/statik"

	"github.com/stone-payments/centrifugo/libcentrifugo/centrifugo"
)

// VERSION of Centrifugo server. Set on build stage.
var VERSION string

func main() {
	centrifugo.Main(VERSION)
}
