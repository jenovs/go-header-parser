package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	return ":" + port
}

func whoami(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	userData := map[string]string{
		"User-Agent":      "",
		"Accept-Language": "",
		"Remote IP":       r.RemoteAddr,
	}

	for k, v := range r.Header {
		if _, ok := userData[k]; ok {
			userData[k] = v[0]
		}
		if k == "X-Forwarded-For" {
			userData["Remote IP"] = v[0]
		}
	}

	res, err := json.Marshal(userData)

	if err != nil {
		fmt.Println("Marshal error: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}

func main() {
	http.HandleFunc("/", whoami)
	http.HandleFunc("/favicon.ico", faviconHandler)
	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
