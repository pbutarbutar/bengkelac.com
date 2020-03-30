package libs

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	LayoutPath  string
	IncludePath string
	Port        int64
}

func Loadconfig(fileName string) (Configuration, error) {

	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Println("error:", err)
	}
	//log.Println("layout path: ", configuration.LayoutPath)
	//log.Println("include path: ", configuration.IncludePath)
	return configuration, err
}
