package main

import (
	"net/http"
	//"fmt"
	"log"
	"encoding/json"
	"fmt"
)

type Rajni struct {
	Name string
}


func homePage(w http.ResponseWriter, r *http.Request){

	a := json.NewDecoder(r.Body)


	ap := Rajni{}

	a.Decode(&ap)



	fmt.Println(ap.Name)
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9081", nil))
}