package hamster

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Set struct
var Config = struct {
	DB struct {
		Host     string
		Port     string
		User     string
		Passwd   string
		DataName string
	}
	Http struct {
		Port string
	}
	Client struct {
		Path string
	}
}{}

func ConfigInit() {
	// Open config file
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	jsonStruct, err := ioutil.ReadAll(file)

	// Json handle
	err = json.Unmarshal(jsonStruct, &Config)
	if err != nil {
		log.Fatal("LOAD CONFIG ERROR:", err)
	}
}
