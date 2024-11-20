package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	name, _ := os.Hostname()
	log.Printf("hello from %s", name)
	fmt.Fprintf(w,
		fmt.Sprintf(`
			{
				"message": "Hello from %s"
			}
		`, name))
}

func main() {
	log.Println("starting server ...")
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("could not start server: ", err)
	}
}
