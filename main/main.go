package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
)

type Provider struct {
	Username string
	Password string
	Host     string
}

func main() {
	t1, err := template.ParseFiles("/Users/zk.wang/go/src/github.com/kotf/resource/test.tf")
	if err != nil {
		panic(err)
	}
	//p := Provider{Username: "test", Password: "test", Host: "172.0.0.1"}
	//t1.Execute(os.Stdout, p)

	pers := `{"Username":"test","Password":"123123","Host":"123123"}`
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(pers), &dat); err == nil {
	}

	f, err := os.Create("./tets.tf")
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	fmt.Println("123123")
	fmt.Println(t1)
	err = t1.Execute(f, dat)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
}
