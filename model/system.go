package model

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/aldenso/zfssareport/utils"
	"github.com/spf13/afero"
)

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

//WriteCSV method to write version info to csv file.
func (version *Version) WriteCSV(fs afero.Fs, dirname string) {
	file, err := utils.CreateFile(fs, dirname, "version.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	fileheader := []string{"nodename", "product", "version", "os_version", "csn", "navagent", "navname",
		"mkt_product", "sp_version", "http", "ssl", "ak_version", "bios_version", "part", "asn",
		"href", "urn", "install_time", "update_time", "boot_time"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	line := fmt.Sprintf("%s;%s;%s;%s;%s;"+
		"%s;%s;%s;%s;%s;"+
		"%s;%s;%s;%s;%s;"+
		"%s;%s;%s;%s;%s;",
		version.NodeName, version.Product, version.Version, version.OSVersion, version.CSN,
		version.NavAgent, version.NavName, version.MktProduct, version.SPVersion, version.HTTP,
		version.SSL, version.AKVersion, version.BIOSVersion, version.Part, version.ASN,
		version.HREF, version.URN, version.InstallTime, version.UpdateTime, version.BootTime)

	record := strings.Split(line, ";")
	if err := writer.Write(record); err != nil {
		log.Fatal(err)
	}
	writer.Flush()
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
