package config

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/Eric-GreenComb/eth-server/bean"
)

// Ethereum Ethereum Config
var Ethereum bean.EthereumConfig

// ServerConfig Server Config
var ServerConfig bean.ServerConfig

const cmdRoot = "core"

var p string

func init() {

	flag.StringVar(&p, "p", "/root/union/config", "set `prefix` path")
	flag.Parse()
	fmt.Println(p)

	// err := loadRemoteConfig(p)
	// if err != nil {
	// 	fmt.Println("load remote config error:", err.Error())
	// 	os.Exit(0)
	// }

	err := loadLocalConfig("./")
	if err != nil {
		fmt.Println("load local config error:", err.Error())
		os.Exit(0)
	}
}

func loadLocalConfig(path string) error {
	local := viper.New()
	local.SetEnvPrefix(cmdRoot)
	local.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	local.SetEnvKeyReplacer(replacer)
	local.SetConfigName(cmdRoot)
	local.AddConfigPath(path)

	err := local.ReadInConfig()
	if err != nil {
		return err
	}

	ServerConfig.Port = strings.Split(local.GetString("server.port"), ",")
	ServerConfig.Mode = local.GetString("server.mode")
	ServerConfig.GormLogMode = local.GetString("server.gorm.LogMode")
	ServerConfig.ViewLimit = local.GetInt("server.view.limit")

	Ethereum.ChainID = local.GetInt64("ethereum.chainID")
	Ethereum.Host = local.GetString("ethereum.host")
	Ethereum.Passphrase = local.GetString("ethereum.passphrase")

	return nil
}

func loadRemoteConfig(path string) error {
	remote := viper.New()
	remote.SetEnvPrefix(cmdRoot)
	remote.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	remote.SetEnvKeyReplacer(replacer)
	remote.SetConfigName(cmdRoot)
	remote.AddConfigPath(path)

	err := remote.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
