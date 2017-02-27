package main

import(
	"fmt"
	"net/http"
	"./parse"
)

func handler(w http.ResponseWriter, r *http.Request) {
	l, err := parse.GetList()
	if err != nil {
		fmt.Fprintf(w, "Error occured:", err)
	}
	fmt.Fprintf(w, "Connecting devices:\n")
	for _, device := range l.Device {
		fmt.Fprintf(w, "%s\n", device.Name.Data)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
