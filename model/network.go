package model

import (
	"fmt"
	"strings"
)

//NetInterfaces struct for network interfaces.
type NetInterfaces struct {
	List []NetInterface `json:"interfaces"`
}

//NetInterface struct for network interface.
type NetInterface struct {
	Admin        bool     `json:"admin"`
	Class        string   `json:"class"`
	CurAddrs     []string `json:"curaddrs"`
	Enable       bool     `json:"enable"`
	HREF         string   `json:"href"`
	Interface    string   `json:"interface"`
	Label        string   `json:"label"`
	Links        []string `json:"links"`
	State        string   `json:"state"`
	V4Addrs      []string `json:"v4addrs"`
	V4DHCP       bool     `json:"V4dhcp"`
	V4DirectNets []string `json:"v4directnets"`
	V6Addrs      []string `json:"v6addrs"`
	V6DHCP       bool     `json:"v6dhcp"`
	V6DirectNets []string `json:"v6directnets"`
}

//PrintNetInterfaceInfo method to print some interface info.
func (interf *NetInterface) PrintNetInterfaceInfo() {
	fmt.Printf("%-15s %-8s %-25s %-20s %-25s %4s\n",
		interf.Interface, interf.Class, strings.Join(interf.Links, ","), interf.Label,
		strings.Join(interf.V4Addrs, ","), interf.State)
}
