package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func loadConfig(file string) ConfigRec {
	log.Println("Loading Config: " + file)
	cnfFile, err := os.Open("data/" + file)
	if err != nil {
		log.Println(err.Error())
	}
	defer cnfFile.Close()

	byteValue, _ := ioutil.ReadAll(cnfFile)
	var loadedConf ConfigRec
	json.Unmarshal(byteValue, &loadedConf)

	return loadedConf
}
