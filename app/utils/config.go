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
	Server ServerYaml   `yaml:"server"`
	Upload UploadYaml   `yaml:"upload"`
	Db     DbConfigYaml `yaml:"db"`
}
type ServerYaml struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}
type UploadYaml struct {
	Dir    string `yaml:"dir"`
	Tmpdir string `yaml:"tmpdir"`
}
type DbConfigYaml struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Db      string `yaml:"db"`
	User    string `yaml:"user"`
	Passwd  string `yaml:"passwd"`
	Charset string `yaml:"charset"`
}

func initConfig() (ConfYaml, error) {

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

	config.Db.Host = viper.GetString("database.host")
	config.Db.Port = viper.GetString("database.port")
	config.Db.User = viper.GetString("database.user")
	config.Db.Passwd = viper.GetString("database.passwd")
	config.Db.Db = viper.GetString("database.db")
	config.Db.Charset = viper.GetString("database.charset")
	return config, nil
}

func GetConfig() ConfYaml {
	Cfg, _ := initConfig()
	return Cfg
}
