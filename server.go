package main

import (
	"net/http"
	"encoding/json"
	"log"
	"os"
	"fmt"
)

type UserData struct {
  UserAgent string `json:"user-agent"`
  Language string  `json:"language"`
	IPAddr string `json:"ip-address"`
	RemoteIPAddr string `json:"remote-ip-address"`
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	return ":"+port
}

func whoami(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	userData := map[string]string{
		"User-Agent": "",
		"Accept-Language": "",
		"IP-Addr": r.Host,
		"Remote IP": r.RemoteAddr,
	}

	fmt.Println(r)

	headers := r.Header
	for k, v := range headers {
		if _, ok := userData[k]; ok {
			userData[k] = v[0]
		}
		if k == "X-Forwarded-For" {
			userData["Remote IP"] = v[0]
		}
	}

	res, _ := json.Marshal(userData)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func faviconHandler(w http.ResponseWriter, r *http.Request)  {
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
