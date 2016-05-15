package config

import (
	"github.com/twitchyliquid64/CNC/logging"
	"crypto/tls"
	"os"
)

var gConfig *Config = nil
var gTls *tls.Config = nil

func Load(fpath string)error{
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
	  logging.Error("config", "Configuration file not found. If you need an example, use example-config.json (cp example-config.json config.json)")
		return err
	}

	conf, err := readConfig(fpath)
	if err == nil{
		gConfig = conf
	} else {
		logging.Error("config", "config.Load() error:", err)
		return err
	}

	tls, err := loadTLS(gConfig.TLS.PrivateKey, gConfig.TLS.Cert)
	if err == nil{
		gTls = tls
	} else {
		logging.Error("config", "config.Load() error:", err)
		return err
	}

	return nil
}

func GetServerName()string{
	checkInitialisedOrPanic()
	return gConfig.Name
}

func TLS()*tls.Config{
	checkInitialisedOrPanic()
	return gTls
}

func All()*Config{
	checkInitialisedOrPanic()
	return gConfig
}

func checkInitialisedOrPanic(){
	if gConfig == nil{
		panic("Config not initialised")
	}
	if gTls == nil{
		panic("TLS not initialised")
	}
}
