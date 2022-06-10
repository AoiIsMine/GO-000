package main

import (
	"go-battle/routers"
)

func main() {
	router := routers.RoutersInit()
	router.Run("localhost:3000")
}
