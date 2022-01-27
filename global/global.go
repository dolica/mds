package global

import (
	"github.com/dolica/mds/config"
	"github.com/spf13/viper"
)

const (
	ConfigFile    = "config.yaml"
	ConfigEnvName = "HTF_CONFIG"
)

//defind global object
var (
	HTF_Config config.Server
	HTF_Viper  *viper.Viper
)
