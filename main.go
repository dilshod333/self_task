package main

import (
	"conn/Handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	rooter := mux.NewRouter()
	rooter.HandleFunc("/create", Handlers.CreatBook)
	rooter.HandleFunc("/get", Handlers.GetBook)
	rooter.HandleFunc("/update", Handlers.UpdateBook)
	rooter.HandleFunc("/delete", Handlers.DeleteBook)
	fmt.Println("Server running at :8080...")
	http.ListenAndServe(":8080", rooter)

}
