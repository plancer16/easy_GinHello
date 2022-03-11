package main

import "diy_ginHello/initRouter"

func main() {
	router := initRouter.SetupRouter()
	router.Run()
}