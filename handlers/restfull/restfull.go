package restfull

import (
	"fmt"
	"net/http"

	"example.com/go-folder-structure/model"
	"example.com/go-folder-structure/pkg/logger"
)

func RestFull(w http.ResponseWriter, r *http.Request) {
	data := model.NewModel("sachin", 25)
	data.GetName()
	fmt.Println(data.GetName())
	fmt.Println(data.GetAge())
	logger.Log("Logger pkg is working.....")
}
