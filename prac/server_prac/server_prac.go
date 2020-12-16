package server_prac

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	fmt.Fprint(w, "Hello, hi\nURL = %s\n", req.URL)
}

func server_prac() {
	fmt.Println("connect to loacalhost:7777/hello")
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
