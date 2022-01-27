package core

import (
	"flag"
	"fmt"
	"os"

	"github.com/dolica/mds/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var command_config string
	if len(path) == 0 {
		flag.StringVar(&command_config, "c", "", "config file path")
		flag.Parse()
		if command_config == "" {
			if configEnv := os.Getenv(global.ConfigEnvName); configEnv == "" {
				path = append(path, global.ConfigFile)
				fmt.Printf("config file is not setting, use the default config [%v].\n", global.ConfigFile)
			} else {
				path = append(path, configEnv)
				fmt.Printf("config file is setting by env, use config file [%v].\n", configEnv)
			}
		} else {
			path = append(path, command_config)
			fmt.Printf("config file is setting by command line, use config file [%v].]\n", command_config)
		}
	}
	runtime_config := viper.New()
	for _, config := range path {
		runtime_config.SetConfigFile(config)
	}
	//add env config to viper
	runtime_config.AutomaticEnv()
	runtime_config.SetConfigType("yaml")
	err := runtime_config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config file occur err, %v", err))
	}

	runtime_config.WatchConfig()
	runtime_config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config file changed,%s", e.Name)
		if err := runtime_config.Unmarshal(&global.HTF_Config); err != nil {
			fmt.Println(err)
		}
	})

	if err := runtime_config.Unmarshal(&global.HTF_Config); err != nil {
		fmt.Println(err)
	}

	return runtime_config
}
