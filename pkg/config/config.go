package config

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"io/ioutil"
)

// "dubbo-go-samples-configcenter-nacos-server", dubbo
func PublishDubbogoConfigToNacos(configPath , dataID, group string, ){
	configCenterNacosConfig, err:= ioutil.ReadFile(configPath)
	dynamicConfig, err := config.NewConfigCenterConfig(
		config.WithConfigCenterProtocol("nacos"),
		config.WithConfigCenterAddress("127.0.0.1:8848")).GetDynamicConfiguration()
	if err != nil {
		panic(err)
	}
	if err := dynamicConfig.PublishConfig(dataID,group , string(configCenterNacosConfig)); err != nil {
		panic(err)
	}
}