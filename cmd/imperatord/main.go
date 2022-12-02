package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/imperator-co/imperator-core/app"
	"github.com/imperator-co/imperator-core/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		"imperator",
		"cosmos",
		app.DefaultNodeHome,
		"imperator-1",
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
