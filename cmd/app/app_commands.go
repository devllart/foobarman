package main

import (
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"fmt"

	"github.com/urfave/cli"
)

var appCommands = []cli.Command{
	{
		Name:  "bar",
		Usage: "Show product from the bar",
		Action: func(c *cli.Context) {
			fmt.Println(state.Bar)
		},
	},

	{
		Name:  "shop",
		Usage: "Show product from the shop",
		Action: func(c *cli.Context) {
			scenes.Show("Store")
		},
	},
}
