package main

import (
	"os"

	"github.com/mycodeku/transtionhelper/server"
	svrcmd "github.com/mycodeku/transtionhelper/server/cmd"
	"github.com/mycodeku/transtionhelper/simapp"
	"github.com/mycodeku/transtionhelper/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
