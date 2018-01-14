package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "controllers/items"
    "controllers/transactions"
)

func initializeRoutes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/items", items.GetItems).Methods("GET")
    r.HandleFunc("/transactions", transactions.CreateTransaction).Methods("POST")
    r.HandleFunc("/transactions", transactions.ListTransactions).Methods("GET")
    r.HandleFunc("/transactions/detail/{trxCode}", transactions.ListTransactionDetails).Methods("GET")
    return r
}

func main() {
    routes := initializeRoutes()
    http.ListenAndServe(":8080", routes)
}
