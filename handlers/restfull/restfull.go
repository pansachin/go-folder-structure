package restfull

import (
	"fmt"

	"example.com/go-folder-structure/model"
	"example.com/go-folder-structure/pkg/logger"
)

func RestFull() {
	data := model.NewModel("sachin", 25)
	data.GetName()
	fmt.Println(data.GetName())
	fmt.Println(data.GetAge())
	logger.Log("Logger pkg is working.....")
}
