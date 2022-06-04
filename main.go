package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var dataParse DataIn

type DataIn []struct {
	IP        string `json:"ip"`
	Timestamp string `json:"timestamp"`
	Ports     []struct {
		Port   int    `json:"port"`
		Proto  string `json:"proto"`
		Status string `json:"status"`
		Reason string `json:"reason"`
		TTL    int    `json:"ttl"`
	} `json:"ports"`
}

func main() {
	argsWithProg := os.Args[1:]
	for _, cmd := range argsWithProg {
		data, err := ioutil.ReadFile("Please Add Full Path here" + cmd + "/mscan.xml")
		if err != nil {
			println(err)
		}
		//fmt.Println(string(data))
		err2 := json.Unmarshal(data, &dataParse)
		if err2 != nil {
			fmt.Println()
		}
		for _, listPorts := range dataParse {

			for _, portValue := range listPorts.Ports {
				fmt.Println(portValue.Port)
			}
		}
	}
}
