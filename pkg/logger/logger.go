package logger

import (
	"fmt"

	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func Log(value string) {
	fmt.Println(value)
}
