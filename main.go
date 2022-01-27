package main

import (
	"fmt"

	"github.com/dolica/mds/core"
	"github.com/dolica/mds/global"
)

func main() {
	core.Viper()
	fmt.Println(global.HTF_Config)
	core.RunServer()
}
