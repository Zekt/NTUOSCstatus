package parse

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)

func GetList() (List, error) {
	body := strings.NewReader(`oid_1=IGD_LANDevice_i_ConnectedAddress_i_&inst_1=1100&ccp_act=get&num_inst=100`)
	req, err := http.NewRequest("POST", "http://10.0.87.1/get_set.ccp", body)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var l List
	err = xml.Unmarshal(buf, &l)
	if err != nil {
		fmt.Println(err)
	}
	for i, device := range l.Device {
		l.Device[i].MAC.Data = ""
		fmt.Println("Device:", device.Name.Data)
	}

	return l, nil
}

type List struct {
	XMLName xml.Name `xml:"root"`
	Device  []Device `xml:"IGD_LANDevice_i_ConnectedAddress_i_"`
}
type Device struct {
	Name name `xml:"igdLanHostStatus_HostName_"`
	IP   ip   `xml:"igdLanHostStatus_HostIPv4Address_"`
	MAC  mac  `xml:"igdLanHostStatus_HostMACAddress_"`
}

type name struct {
	Data string `xml:",cdata"`
}
type ip struct {
	Data string `xml:",cdata"`
}
type mac struct {
	Data string `xml:",cdata"`
}
