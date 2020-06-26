package main
import (
	"io"
	"net/http"
	"log"
)
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I am Vik")
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}