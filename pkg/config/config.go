package config

import "github.com/spf13/viper"

type RPCConfig struct {
	URL           string
	UserName      string
	Password      string
	TokenSymbol   string
	AddressLength int
}

var rpcConfig *RPCConfig

func GetRPCConfig() *RPCConfig {
	if rpcConfig == nil {
		rpcConfig = newRPCConfig(viper.Sub("data.http"))
	}

	return rpcConfig
}

func newRPCConfig(conf *viper.Viper) *RPCConfig {
	rpcConfig := &RPCConfig{
		URL:           conf.GetString("node.url"),
		UserName:      conf.GetString("node.username"),
		Password:      conf.GetString("node.password"),
		TokenSymbol:   conf.GetString("node.token-symbol"),
		AddressLength: conf.GetInt("node.address-length"),
	}

	return rpcConfig
}
