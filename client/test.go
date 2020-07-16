package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	fmt.Println("First test")
	firstClient()
	fmt.Println("")
	fmt.Println("Second test")
	secondClient()

}

func firstClient() {
	host, _ := os.Hostname()

	dataJSON, err := json.Marshal(map[string]string{
		"Host": host,
		"Time": time.Now().String(),
	})
	if err != nil {
		log.Fatal("11", err)
	}

	response, err := http.Post("http://"+os.Getenv("IP_PROXY")+":"+os.Getenv("PORT_PROXY")+"/demo", "application/json", bytes.NewBuffer(dataJSON))
	if err != nil {
		log.Fatal("12", err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("13", err)
	}
	fmt.Println("First client: ", string(body))
}

func secondClient() {

	dataJSON, err := json.Marshal("Test suite for demo2")
	if err != nil {
		log.Fatal("21", err)
	}

	response, err := http.Post("http://"+os.Getenv("IP_PROXY")+":"+os.Getenv("PORT_PROXY")+"/demo2", "application/json", bytes.NewBuffer(dataJSON))
	if err != nil {
		log.Fatal("22", err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("23", err)
	}
	fmt.Println("Second client: ", string(body))
}
