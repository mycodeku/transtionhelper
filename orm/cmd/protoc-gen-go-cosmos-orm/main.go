package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/mycodeku/transtionhelper/orm/internal/codegen"
)

func main() {
	protogen.Options{}.Run(codegen.PluginRunner)
}
