package test

import (
	"encoding/json"
	"fmt"
	"iot-platform/define"
	"iot-platform/helper"
	"testing"
)

var userServerAddr = "http://localhost:8888"

func TestUserLogin(t *testing.T) {
	m := define.M{
		"username": "goo",
		"password": "123456",
	}

	data, _ := json.Marshal(m)

	rep, err := helper.HttpPost(userServerAddr+"/user/login", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))

}

func TestUserAdd(t *testing.T) {
	m := define.M{
		"username": "goo",
		"password": "123456",
	}

	data, _ := json.Marshal(m)

	rep, err := helper.HttpPost(userServerAddr+"/user/add", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))

}
