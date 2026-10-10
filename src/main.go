package main

import (
	"fmt"
	"log"
	"net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello,World")
	
}

func main(){
	http.HandleFunc("/",helloHandler)
	log.Println("Starting server on : 8080")
	log.Fatal(http.ListenAndServe(":8000",nil))
}
