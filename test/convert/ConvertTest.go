package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

func main() {
	r := Result[*User]{
		Code: 200,
		Msg:  "成功",
		Data: nil,
	}

	resBytes, _ := json.Marshal(&r)

	var m = make(map[string]any)
	err := json.Unmarshal(resBytes, &m)
	if err != nil {
		panic(err)
	}

	var result Result[*User]
	err = mapstructure.Decode(m, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", m)
}

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
