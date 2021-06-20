package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

//var defaultConf = []byte(`
//server:
//    ip: 0.0.0.0
//    port: 8088
//upload:
//    dir: ./upload
//    tmpdir: ./upload/tmp
//`)

type ConfYaml struct {
	Server ServerYaml `yaml:"server"`
	Upload UploadYaml `yaml:"upload"`
}
type ServerYaml struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}
type UploadYaml struct {
	Dir    string `yaml:"dir"`
	Tmpdir string `yaml:"tmpdir"`
}

func InitConfig() (ConfYaml, error) {

	var config ConfYaml
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	//todo set default values,the default struct already init.

	config.Server.Ip = viper.GetString("server.ip")
	config.Server.Port = viper.GetString("server.port")
	config.Upload.Tmpdir = viper.GetString("upload.tmpdir")
	config.Upload.Dir = viper.GetString("upload.dir")

	return config, nil
}
