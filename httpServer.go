package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/api/inventory/v1/service", func() {}).Methods("PUT")
	//for HTTP
	err := http.ListenAndServe(":8080", router)
	// for HTTPS
	err := http.ListenAndServeTLS(":8080", router)
	if err != nil {
		log.Printf("error %v", err)
	}
	log.Printf("listening")
}

func makeRequest()  {
	tr := &http.Transport{
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		Proxy:                 http.ProxyFromEnvironment,
	}
	client := http.Client{
		Transport: tr,
		Timeout:   30 * time.Second,
	}
	http.NewRequestWithContext(context.b, "GET","http://xyz.com",{})
	response, err = client.Do(req)
}

func context()  {
	
}
