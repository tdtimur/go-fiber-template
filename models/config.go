package models

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Config = setupConfig()

func setupConfig() *viper.Viper {
	conf := viper.New()
	conf.SetConfigName("config")                // name of config file (without extension)
	conf.AddConfigPath("$HOME/.fiber-template") // call multiple times to add many search paths
	conf.AddConfigPath(".")
	conf.SetDefault("MONGODB_HOST", "mongodb://localhost:27017")
	conf.SetDefault("JWT_SECRET", "somethingsupersecret,changemeonproduction")
	conf.WatchConfig()
	conf.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	if err := conf.ReadInConfig(); err != nil {
		log.Fatalf("Viper configuration error: %s", err)
		return nil
	} // Find and read the config file
	return conf
}
