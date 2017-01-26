package model

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/aldenso/zfssareport/utils"
	"github.com/spf13/afero"
)

//Cluster struct for Cluster.
type Cluster struct {
	clusterStruct `json:"cluster"`
}

type clusterStruct struct {
	Description     string            `json:"description"`
	PeerASN         string            `json:"peer_asn"`
	PeerDescription string            `json:"peer_description"`
	PeerHostname    string            `json:"peer_hostname"`
	PeerState       string            `json:"peer_state"`
	Resources       []clusterResource `json:"resources"`
}

type clusterResource struct {
	Details   []string `json:"details"`
	HREF      string   `json:"href"`
	Owner     string   `json:"owner"`
	Type      string   `json:"type"`
	UserLabel string   `json:"user_label"`
}

//PrintClusterInfo to print some version info.
func (cluster *Cluster) PrintClusterInfo() {
	utils.Header("ZFS Storage Appliance Cluster")
	fmt.Printf("%-30s %-30s %-15s %-15s\n", "description", "peer description", "peer Hostname", "peer state")
	fmt.Printf("%-30s %-30s %-15s %-15s\n", cluster.Description, cluster.PeerDescription, cluster.PeerHostname,
		cluster.PeerState)
	fmt.Println("=====================================================================================================================")
	fmt.Printf("%-15s %-15s %-15s %-15s %15s\n", "resource", "type", "owner", "details", "label")
	for _, resource := range cluster.Resources {
		sliceref := strings.Split(resource.HREF, "/")
		newhref := sliceref[len(sliceref)-1]
		fmt.Printf("%-15s %-15s %-15s %-15s %-15s\n", newhref, resource.Type, resource.Owner, resource.Details,
			resource.UserLabel)
	}
}

//WriteCSV method to write cluster info to csv file.
func (cluster *Cluster) WriteCSV(fs afero.Fs, dirname string) {
	file, err := utils.CreateFile(fs, dirname, "cluster.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	fileheader := []string{"description", "peer_asn", "peer_description", "peer_hostname", "peer_state"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	line := fmt.Sprintf("%s;%s;%s;%s;%s", cluster.Description, cluster.PeerASN, cluster.PeerDescription,
		cluster.PeerHostname, cluster.PeerState)
	record := strings.Split(line, ";")
	if err := writer.Write(record); err != nil {
		log.Fatal(err)
	}

	fileheader = []string{"details", "href", "owner", "type", "user_label"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	for _, resource := range cluster.Resources {

		line = fmt.Sprintf("%s;%s;%s;%s;%s", resource.Details, resource.HREF, resource.Owner, resource.Type,
			resource.UserLabel)
		record := strings.Split(line, ";")
		if err := writer.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	writer.Flush()
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
