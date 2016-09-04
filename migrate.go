package main

import (
	"encoding/json"
	"github.com/tarantool/go-tarantool"
	"io/ioutil"
	"log"
	"time"
)

var (
	config Configuration
	tnt    *tarantool.Connection
)

type (
	Configuration struct {
		Tarantool Tarantool `json:"tarantool"`
	}

	Tarantool struct {
		Server string `json:"server"`
		User   string `json:"user"`
		Password  string `json:"password"`
		File   string `json:"file"`
	}
)
func main() {
	// Read configuration
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("[Configuration] Reading error: %+v", err)
	} else {
		log.Println("[Configuration] Read successfully.")
	}
	// Decode configuration
	if err = json.Unmarshal(file, &config); err != nil {
		log.Fatalf("[Configuration] Decoding error: %+v", err)
	}

	// Tarantool connect
	opts := tarantool.Opts{
		Timeout:       500 * time.Millisecond,
		Reconnect:     1 * time.Second,
		MaxReconnects: 3,
		User:          config.Tarantool.User,
		Pass:          config.Tarantool.Password,
	}
	newTnt, err := tarantool.Connect(config.Tarantool.Server, opts)
	if err != nil {
		log.Fatalf("[TNT] Failed to connect: %s", err.Error())
	} else {
		tnt = newTnt
		log.Printf("[TNT] Successful connected to: %s", tnt.Greeting.Version)
	}
	file, ferr := ioutil.ReadFile(config.Tarantool.File)
	if ferr != nil {
		log.Fatalf("[TNT] Migrate file reading error: %+v", ferr)
	} else {
		log.Println("[TNT] Migrate file read successfully.")
		resp, err := tnt.Eval(string(file), []interface{}{})
		log.Println("[TNT] Migrate err: ", err)
		log.Println("[TNT] Migrate Code: ", resp.Code)
		log.Println("[TNT] Migrate Data: ", resp.Data)
	}
}