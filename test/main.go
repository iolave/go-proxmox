package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	A string `json:"a"`
	B string `json:"b"`
}

type S struct {
	Data any `json:"data"`
}

func main() {
	s := S{
		Data: struct {
			A string `json:"a"`
			B string `json:"b"`
		}{
			A: "a",
			B: "b",
		},
	}

	b, _ := json.Marshal(s)

	newS := S{}
	err := json.Unmarshal(b, &newS)
	if err != nil {
		panic(err)
	}

	myData := MyData{}
	b, err = json.Marshal(newS.Data)
	json.Unmarshal(b, &myData)

	fmt.Println(myData)
}
