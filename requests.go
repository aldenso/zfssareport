package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aldenso/zfssareport/model"
)

var (
	//HTTPClientCfg to define http client configuration for requests.
	HTTPClientCfg = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//Timeout is high because some ZFSSA storage are complex and takes too much time to retrieve some info.
	client = &http.Client{Transport: HTTPClientCfg, Timeout: 100 * time.Second}
)

//GetPools get all pools.
func GetPools() model.Pools {
	if silent {
		fmt.Println("getting pools info.")
	}
	var pools model.Pools
	req, err := http.NewRequest("GET", URL+"/storage/v1/pools", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&pools)
	if err != nil {
		log.Fatal(err)
	}
	return pools
}

//GetProjects get all projects in a pool.
func GetProjects() *model.Projects {
	if silent {
		fmt.Println("getting projects info.")
	}
	projects := &model.Projects{}
	fullurl := fmt.Sprintf("%s/storage/v1/projects", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		log.Fatal(err)
	}
	return projects
}

//GetFilesystems get all Filesystems in a project.
func GetFilesystems() *model.Filesystems {
	if silent {
		fmt.Println("getting filesystems info.")
	}
	filesystems := &model.Filesystems{}
	fullurl := fmt.Sprintf("%s/storage/v1/filesystems", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&filesystems)
	if err != nil {
		log.Fatal(err)
	}
	return filesystems
}

//GetLUNS get all LUNS in a project.
func GetLUNS() *model.LUNS {
	if silent {
		fmt.Println("getting luns info.")
	}
	luns := &model.LUNS{}
	fullurl := fmt.Sprintf("%s/storage/v1/luns", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&luns)
	if err != nil {
		log.Fatal(err)
	}
	return luns
}

func getZFSSAVersion() {
	if silent {
		fmt.Println("getting version info.")
	}
	version := &model.Version{}
	fullurl := fmt.Sprintf("%s/system/v1/version", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Status: '%s', check your credentials.", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&version)
	if err != nil {
		log.Fatal(err)
	}
	version.WriteCSV(Fs, dirname)
}

func getNetInterfaces() *model.NetInterfaces {
	if silent {
		fmt.Println("getting network interfaces info.")
	}
	interfaces := &model.NetInterfaces{}
	fullurl := fmt.Sprintf("%s/network/v1/interfaces", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&interfaces)
	if err != nil {
		log.Fatal(err)
	}
	return interfaces
}

//GetFCInitiators get all initiators in zfssa.
func GetFCInitiators() *model.FCInitiators {
	if silent {
		fmt.Println("getting FC initiators info.")
	}
	initiators := &model.FCInitiators{}
	fullurl := fmt.Sprintf("%s/san/v1/fc/initiators", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&initiators)
	if err != nil {
		log.Fatal(err)
	}
	return initiators
}

//GetFCInitiatorGroups get all initiators in zfssa.
func GetFCInitiatorGroups() *model.FCInitiatorGroups {
	if silent {
		fmt.Println("getting FC initiators groups info.")
	}
	groups := &model.FCInitiatorGroups{}
	fullurl := fmt.Sprintf("%s/san/v1/fc/initiator-groups", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&groups)
	if err != nil {
		log.Fatal(err)
	}
	return groups
}

//GetFCTargets get all targets in zfssa.
func GetFCTargets() *model.FCTargets {
	if silent {
		fmt.Println("getting FC targets info.")
	}
	targets := &model.FCTargets{}
	fullurl := fmt.Sprintf("%s/san/v1/fc/targets", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&targets)
	if err != nil {
		log.Fatal(err)
	}
	return targets
}

//GetIscsiInitiators get all iscsi initiators in zfssa.
func GetIscsiInitiators() *model.IscsiInitiators {
	if silent {
		fmt.Println("getting ISCSI initiators info.")
	}
	initiators := &model.IscsiInitiators{}
	fullurl := fmt.Sprintf("%s/san/v1/iscsi/initiators", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&initiators)
	if err != nil {
		log.Fatal(err)
	}
	return initiators
}

//GetIscsiInitiatorGroups get all iscsi initiators in zfssa.
func GetIscsiInitiatorGroups() *model.IscsiInitiatorGroups {
	if silent {
		fmt.Println("getting ISCSI initiators groups info.")
	}
	groups := &model.IscsiInitiatorGroups{}
	fullurl := fmt.Sprintf("%s/san/v1/iscsi/initiator-groups", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&groups)
	if err != nil {
		log.Fatal(err)
	}
	return groups
}

func getClusterInfo() {
	cluster := &model.Cluster{}
	fullurl := fmt.Sprintf("%s/hardware/v1/cluster", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Status: '%s', check your credentials.", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&cluster)
	if err != nil {
		log.Fatal(err)
	}
	if !silent {
		cluster.PrintClusterInfo()
	} else {
		fmt.Println("getting cluster info.")
	}
	cluster.WriteCSV(Fs, dirname)
}

//GetChassis get chassis in zfssa.
func GetChassis() *model.ChassisAll {
	if silent {
		fmt.Println("getting chassis info.")
	}
	chassisslice := &model.ChassisAll{}
	fullurl := fmt.Sprintf("%s/hardware/v1/chassis", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&chassisslice)
	if err != nil {
		log.Fatal(err)
	}
	return chassisslice
}

//GetProblems get problems in zfssa.
func GetProblems() *model.Problems {
	if silent {
		fmt.Println("getting problems info.")
	}
	problems := &model.Problems{}
	fullurl := fmt.Sprintf("%s/problem/v1/problems", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&problems)
	if err != nil {
		log.Fatal(err)
	}
	return problems
}

//GetNetDevices get network devices in zfssa.
func GetNetDevices() *model.NetDevices {
	if silent {
		fmt.Println("getting network devices info.")
	}
	devices := &model.NetDevices{}
	fullurl := fmt.Sprintf("%s/network/v1/devices", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&devices)
	if err != nil {
		log.Fatal(err)
	}
	return devices
}

//GetNetDatalinks get network datalinks in zfssa.
func GetNetDatalinks() *model.NetDatalinks {
	if silent {
		fmt.Println("getting network datalinks info.")
	}
	datalinks := &model.NetDatalinks{}
	fullurl := fmt.Sprintf("%s/network/v1/datalinks", URL)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-User", USER)
	req.Header.Add("X-Auth-Key", PASSWORD)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&datalinks)
	if err != nil {
		log.Fatal(err)
	}
	return datalinks
}
