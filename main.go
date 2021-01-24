package main

import (
	"fmt"
	"message_api/interface/router"
	"net/http"
	"os"
)

func main() {
	r := router.CreateRouter()
	router.PrintRoutes(r)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err.Error())
	}
}
