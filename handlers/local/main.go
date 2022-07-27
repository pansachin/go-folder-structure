package main

import (
	"fmt"

	"example.com/go-folder-structure/model"
)

func main() {
	data := model.NewModel("sachin", 25)
	data.GetName()
	fmt.Println(data.GetName())
	fmt.Println(data.GetAge())
}
