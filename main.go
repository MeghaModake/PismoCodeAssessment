package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the server package.")

	/*
		 	var nd Doer
			nd = NewMeghasTodo()
			router := mux.NewRouter()
			router.HandleFunc("/login", loginHanlder)
			router.HandleFunc("/create", nd.CreateHandler).Methods("POST")
			//router.HandleFunc("/{id}", GetHandler).Methods("GET")
			router.Handle("/{id}", Authorizarion(nd.GetHandler))
			http.ListenAndServe(":8080", router)
			fmt.Println("Server is running on port 8080")
	*/
}
