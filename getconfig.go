package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gobuffalo/packr/v2"
	"gopkg.in/yaml.v2"
)

type SearchConfig struct {
	Baseurl    string `yaml:"baseurl"`
	Baselimit  string `yaml:"baselimit"`
	Maxresults int    `yaml:"maxresults"`
	Jobs       []struct {
		Job      interface{} `yaml:"job"`
		Keyword  string      `yaml:"keyword"`
		Location []string    `yaml:"location"`
	} `yaml:"jobs"`
}

func GetConfig() {
	if _, err := os.Stat("./config/config.yml"); err == nil { // check if config file exists
		yamlFile, err := ioutil.ReadFile("./config/config.yml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			panic(err)
		}
	} else if os.IsNotExist(err) { // config file not included, use embedded config
		yamlFile := packr.New("configBox", "./config")

		configFile, err := yamlFile.Find("config.yml")
		if err != nil {
			log.Fatal(err)
		}
		// yamlFile, err := Asset("config/config.yml")
		// if err != nil {
		// 	panic(err)
		// }
		err = yaml.Unmarshal(configFile, &config)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Schrodinger: file may or may not exist. See err for details.")
		// panic(err)
	}
}
