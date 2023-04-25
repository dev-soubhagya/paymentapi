package main

import (
	"fmt"
	"net/http"

	"github.com/dev-soubhagya/paymentapi/services"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Test Run")
	mux := Router()
	http.ListenAndServe(":8000", mux)
}
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/test", test).Methods("GET")
	r.HandleFunc("/send", services.UserToMerchant).Methods("POST")
	r.HandleFunc("/refund", services.MerchantToUser).Methods("POST")
	r.HandleFunc("/merhant-withraw", services.MerchantWithraw).Methods("POST")
	r.HandleFunc("/merhant-txn-history", services.MerchantTransactionCheck).Methods("POST")
	return r
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Payment Application Started (Golang)")
	w.WriteHeader(200)
}
