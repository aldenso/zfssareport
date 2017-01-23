package model

import "fmt"
import "github.com/aldenso/zfssareport/utils"

//Version get ZFSSA version info.
type Version struct {
	versionStruct `json:"version"`
}

type versionStruct struct {
	AKVersion   string `json:"ak_version"`
	ASN         string `json:"asn"`
	BIOSVersion string `json:"bios_version"`
	BootTime    string `json:"boot_time"`
	//BootTime    time.Time `json:"boot_time"`
	CSN         string `json:"csn"`
	HREF        string `json:"href"`
	HTTP        string `json:"http"`
	InstallTime string `json:"install_time"`
	//InstallTime time.Time `json:"install_time"`
	MktProduct string `json:"mkt_product"`
	NavAgent   string `json:"navagent"`
	NavName    string `json:"navname"`
	NodeName   string `json:"nodename"`
	OSVersion  string `json:"os_version"`
	Part       string `json:"part"`
	Product    string `json:"product"`
	SPVersion  string `json:"sp_version"`
	SSL        string `json:"ssl"`
	UpdateTime string `json:"update_time"`
	//UpdateTime  time.Time `json:"update_time"`
	URN     string `json:"urn"`
	Version string `json:"version"`
}

//PrintVersionInfo to print some version info.
func (version *Version) PrintVersionInfo() {
	utils.Header("ZFS Storage Appliance Version")
	fmt.Printf("%-15s %-24s %-23s %-11s %40s\n", "nodename", "version", "product", "csn", "boot_time")
	fmt.Println("=====================================================================================================================")
	fmt.Printf("%-15s %-24s %-23s %-11s %40s\n",
		version.NodeName, version.Version, version.Product, version.CSN, version.BootTime)
}

/*func (version *Version) WriteCSV()  {

}*/
