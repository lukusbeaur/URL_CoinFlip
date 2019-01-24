package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	
}
func handler(w http.ResponseWriter, r *http.Request) {
	guess := strings.Trim(r.URL.Path, "/")
	coin := map[int]string{
		0: "Heads",
		1: "Tails",
		2: "I can not believe this. It landed on the edge",
	}
	coinflip := rand.Intn(3)
	result := coin[coinflip]
	fmt.Fprintf(w, "Drum roll please..... %v .\n", coin[coinflip])

	switch {
	case result == "I can not believe this. It landed on the edge":
		fmt.Fprintf(w, "We go agane bois\n")
		break
	case result == guess:
		fmt.Fprintf(w, "You guessed %v and you were correct!\n", guess)
	case result != guess:
		fmt.Fprintf(w, "You guessed %v and were wrong! loooserrrr\n", guess)

	}

}
