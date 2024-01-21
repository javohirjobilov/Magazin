package main

import (
	"fmt"
	"net/http"
	"Magazin/Hendler"

)

func main (){

	fmt.Println("Server is working... :8080")


	http.HandleFunc("/catagory", Hendler.CatagoryHendler)
	http.HandleFunc("/product", Hendler.PostHendler)
	http.HandleFunc("/manage", Hendler.ManageProductHendler )
	

	http.ListenAndServe(":8080", nil)

}