package main

import "net/http" 
import "github.com/gorilla/mux"

// Custom packages
import "./views"

func UrlRegister() {
	router := mux.NewRouter()
	router.HandleFunc("/", views.IndexView)
	router.HandleFunc("/api/block/{block_number:[0-9]+}/total", views.API_GetTotalTransactionsAmountOfEthBlockView).Methods("GET")
	http.Handle("/", router)
}

func StartServer(port string) {
	http.ListenAndServe(":" + port, nil)
}

func main() {
 	UrlRegister()
 	StartServer("8080")
}
