package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"remote-OHM/utils"
)

func main() {
	conf := utils.ReadConf()
	url := "http://" + conf.IP + ":" + conf.Port + "/data.json"

	if resp, err := http.Get(url); err != nil {
		fmt.Printf("Error while fetching data :\n")
		fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
	} else {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			fmt.Printf("Error while fetching data :\n")
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		} else {
			var ohmdata utils.OhmData
			if err := json.Unmarshal(body, &ohmdata); err == nil {
				utils.SaveSensor(ohmdata)
			} else {
				fmt.Printf("Error while fetching data :\n")
				fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
			}
		}
	}
}
