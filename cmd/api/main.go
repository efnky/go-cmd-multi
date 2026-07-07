package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "go-cmd-multi api up on :%s\n", port)
	})
	fmt.Println("api Listening on :" + port)
	http.ListenAndServe(":"+port, nil)
}
