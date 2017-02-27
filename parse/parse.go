package parse

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

func GetList() (List, error) {
	body := strings.NewReader(`oid_4=IGD_LANDevice_i_ConnectedAddress_i_&inst_4=1100&ccp_act=get&num_inst=24`)
	req, err := http.NewRequest("POST", "http://10.0.87.1/get_set.ccp", body)
	if err != nil {
		return List{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return List{}, err
	}
	buf := make([]byte, 2048)
	n, err := resp.Body.Read(buf)
	defer resp.Body.Close()

	var l List
	err = xml.Unmarshal(buf[:n], &l)
	if err != nil {
		if err.Error() != "EOF" {
			return List{}, err
		}
	}
	for _, device := range l.Device {
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
	Ip   ip   `xml:"igdLanHostStatus_HostIPv4Address_"`
	Mac  mac  `xml:"igdLanHostStatus_HostMACAddress_"`
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
