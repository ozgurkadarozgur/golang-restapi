package main

import (
	"fmt"
	"pro/restapi/app/model"
	"pro/restapi/app/router"
)

func main() {

	fmt.Println("Server running on http://localhost:8082")

	model.Init()
	router.Init()

}
