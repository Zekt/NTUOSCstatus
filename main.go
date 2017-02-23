package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	body := strings.NewReader(`oid_4=IGD_LANDevice_i_ConnectedAddress_i_&inst_4=1100&ccp_act=get&num_inst=24`)
	req, err := http.NewRequest("POST", "http://10.0.87.1/get_set.ccp", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	buf := make([]byte, 2048)
	n, _ := resp.Body.Read(buf)
	fmt.Println(string(buf[:n]))
	defer resp.Body.Close()
}
