package main

import (
	"./parse"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func listHandler(w http.ResponseWriter, r *http.Request, user, pass string) {
	body := strings.NewReader("username=" + user + "&password=" + pass)
	req, err := http.NewRequest("POST", "http://10.0.87.1/login.ccp", body)
	if err != nil {
		fmt.Println("Error creating request: ", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error logging in: ", err)
	}
	defer resp.Body.Close()
	l, err := parse.GetList()
	fmt.Println(l)
	if err != nil {
		fmt.Println("Error parsing list: ", err)
	}
	ret, err := json.Marshal(l)
	if err != nil {
		fmt.Println("Error generating json: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
}

func main() {
	if len(os.Args) < 3 {
		panic("Not enough arguments!")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./index.html") })
	http.HandleFunc("/list/", func(w http.ResponseWriter, r *http.Request) { listHandler(w, r, os.Args[1], os.Args[2]) })
	http.ListenAndServe(":8080", nil)
}
