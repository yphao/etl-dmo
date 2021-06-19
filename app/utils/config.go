package utils

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
)

var defaultConf = []byte(`
server:
    ip: 0.0.0.0
    port: 8088
upload:
    dir: ./upload
    tmpdir: ./upload/tmp
`)

type ConfYaml struct {
	server ServerYaml `yaml:"server"`
	upload UploadYaml `yaml:"upload"`
}
type ServerYaml struct {
	ip   string `yaml:"ip"`
	port string `yaml:"port"`
}
type UploadYaml struct {
	dir    string `yaml:"dir"`
	tmpdir string `yaml:"tmpdir"`
}

func InitConfig() (ConfYaml, error) {
	var config ConfYaml

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	//err := viper.ReadInConfig() //读取配置文件信息
	//if err != nil {
	//	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	//}
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
		// 加载默认配置
		config.server.ip = viper.GetString("server.ip")
		config.server.port = viper.GetString("server.port")
		config.upload.tmpdir = viper.GetString("upload.tmpdir")
		config.upload.dir = viper.GetString("upload.dir")
		fmt.Println("DefaultConf", string(defaultConf))
		return config, err
	}

	config.server.ip = viper.GetString("server.ip")
	config.server.port = viper.GetString("server.port")
	config.upload.tmpdir = viper.GetString("upload.tmpdir")
	config.upload.dir = viper.GetString("upload.dir")

	return config, nil
}
