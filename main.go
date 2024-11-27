package main

import (
	"gin-framework/routers"
)

func main() {
	routers.StartServer().Run(":8080")
}
