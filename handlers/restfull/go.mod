module example.com/go-folder-structure/handlers/restfull

go 1.16

require (
	example.com/go-folder-structure/model v0.0.0
	example.com/go-folder-structure/pkg v0.0.0
	github.com/GoogleCloudPlatform/functions-framework-go v1.5.3 // indirect
)

replace (
	example.com/go-folder-structure/model => ../../model
	example.com/go-folder-structure/pkg => ../../pkg
)
