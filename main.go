package main

import (
	"fmt"
	"net/http"

	routes "github.com/sahandPgr/Email-verifier/routes"
)

func main() {
	route := routes.InitializeRoutes()
	fmt.Print("Server is running on port 8080\n")
	http.ListenAndServe(":8080", route)
}
