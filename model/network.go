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

//NetDatalinks struct for network datalinks.
type NetDatalinks struct {
	List []NetDatalink `json:"datalinks"`
}

//NetDatalink struct for network datalink.
type NetDatalink struct {
	Class    string   `json:"class"`
	Datalink string   `json:"datalink"`
	Duplex   string   `json:"duplex"`
	HREF     string   `json:"href"`
	Label    string   `json:"label"`
	Links    []string `json:"links"`
	MAC      string   `json:"mac"`
	MTU      int      `json:"mtu"`
	Speed    string   `json:"speed"`
}

//PrintNetDatalinkInfo method to print some datalink info.
func (datalink *NetDatalink) PrintNetDatalinkInfo() {
	fmt.Printf("%-10s %-15s %-15s %-15s %-20s %-10d %s\n",
		datalink.Class, datalink.Datalink, datalink.Label, strings.Join(datalink.Links, ","), datalink.MAC,
		datalink.MTU, datalink.Speed)
}

//Routes struct for network routes.
type Routes struct {
	List []Route `json:"routes"`
}

//Route struct for network route.
type Route struct {
	Status      string `json:"status"`
	Family      string `json:"family"`
	Destination string `json:"destination"`
	Mask        int    `json:"mask"`
	HREF        string `json:"href"`
	Interface   string `json:"interface"`
	Type        string `json:"type"`
	Gateway     string `json:"gateway"`
}

//PrintRouteInfo method to print some datalink info.
func (route *Route) PrintRouteInfo() {
	fmt.Printf("%-18s %-18s %-10s %-8s %-6d %-10s %s\n",
		route.Destination, route.Gateway, route.Interface, route.Status, route.Mask, route.Type, route.Family)
}
