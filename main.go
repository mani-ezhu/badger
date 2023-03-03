package main

import (
	"badger/config"
	"badger/constants"
	"badger/handler"
	"flag"
	"log"
	"net/http"
)

// main ...
func main() {

	// parsing config absolute file path
	var configFilePath string
	flag.StringVar(&configFilePath, constants.ConstConfig, constants.ConstConfigFilePath, constants.ConstConfigDesc)
	flag.Parse()
	config.LoadConfigFromFile(configFilePath, config.ConfigReaderFunc())

	// api endpoints
	http.HandleFunc(constants.SpamNofityEndpoint, handler.SpammerHandler)

	// server will listen on 8000 port
	log.Println("listen on", 8000)
	log.Fatal(http.ListenAndServe(constants.Port, nil))
}
