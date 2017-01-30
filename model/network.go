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

//NetDevices struct for network devices.
type NetDevices struct {
	List []NetDevice `json:"devices"`
}

//NetDevice struct for network device.
type NetDevice struct {
	Active     bool   `json:"active"`
	Device     string `json:"device"`
	Duplex     string `json:"duplex"`
	FactoryMAC string `json:"factory_mac"`
	HREF       string `json:"href"`
	Media      string `json:"media"`
	Speed      string `json:"Speed"`
	UP         bool   `json:"up"`
}

//PrintNetDeviceInfo method to print some device info.
func (device *NetDevice) PrintNetDeviceInfo() {
	fmt.Printf("%-8s %-8t %-8t %-15s %-15s %-18s %12s\n",
		device.Device, device.Active, device.UP, device.Speed, device.Media, device.FactoryMAC, device.Duplex)
}
