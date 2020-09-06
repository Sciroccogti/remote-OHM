package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// ReadConf : Read conf.yml
func ReadConf() (conf *Conf) {
	fmt.Println("Reading conf.yml ...")
	conf = new(Conf)
	if _, err := os.Stat("conf.yml"); os.IsNotExist(err) {
		fmt.Println("conf.yml not found, starting with default value ...")
	} else {
		yamlFile, err := ioutil.ReadFile("conf.yml")
		if err != nil {
			fmt.Printf("Error while reading conf.yml :\n")
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		}
		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			fmt.Printf("Error while reading conf.yml :\n")
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		}
	}

	// Set default values
	SetDefault(conf)

	SaveConf(conf)

	return conf
}

// SetDefault : Set default value of the conf
func SetDefault(conf *Conf) {
	if conf.IP == "" {
		conf.IP = "192.168.1.100"
	}

	if conf.Port == "" {
		conf.Port = "8085"
	}
}

// SaveConf : Save the conf.yml
func SaveConf(conf *Conf) {
	fmt.Println("Saving conf.yml ...")
	yamlChanged, err := yaml.Marshal(conf)
	if err != nil {
		fmt.Printf("Error while saving conf.yml :\n")
		fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
	}
	err = ioutil.WriteFile("conf.yml", yamlChanged, 0644)
}
