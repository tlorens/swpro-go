package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() ConfigRec {
	viper.AddConfigPath("./data")
	viper.SetConfigFile("config.json")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("fatal error config file: %s", err))
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	var loadedConf ConfigRec
	err := viper.Unmarshal(&loadedConf)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return loadedConf
}

// func LoadJson(file string) ConfigRec {
// 	log.Println("Loading Config: " + file)
// 	cnfFile, err := os.Open("data/" + file)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	defer cnfFile.Close()

// 	byteValue, _ := ioutil.ReadAll(cnfFile)
// 	var loadedConf ConfigRec
// 	json.Unmarshal(byteValue, &loadedConf)

// 	return loadedConf
// }
