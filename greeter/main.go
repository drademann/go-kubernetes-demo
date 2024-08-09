package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	name := "Nobody"
	if len(args) > 0 {
		name = args[0]
	}
	http.HandleFunc("/greet", newGreetHandler(name))

	log.Printf("starting %s Server on 8080\n", name)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("couldn't start server: %s\n", err.Error())
	}
}

func newGreetHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("This is %s responding!\n", name)
		log.Printf(msg)
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte(msg))
		if err != nil {
			log.Fatalf("couldn't write response: %s\n", err.Error())
		}
	}
}
