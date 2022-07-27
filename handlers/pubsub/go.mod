module example.com/go-folder-structure/handlers/pubsub

go 1.16

require (
   example.com/go-folder-structure/model v0.0.0 
	example.com/go-folder-structure/pkg v0.0.0
)

replace example.com/go-folder-structure/model => ../../model 
replace example.com/go-folder-structure/pkg => ../../pkg