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
	client = &http.Client{Transport: HTTPClientCfg, Timeout: 60 * time.Second}
)

//GetPools get all pools.
func GetPools() model.Pools {
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
func GetProjects(pool string) model.Projects {
	var projects model.Projects
	fullurl := fmt.Sprintf("%s/%s/%s/%s", URL, "storage/v1/pools", pool, "projects")
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
func GetFilesystems(pool string, project string) *model.Filesystems {
	filesystems := &model.Filesystems{}
	fullurl := fmt.Sprintf("%s/storage/v1/pools/%s/projects/%s/filesystems", URL, pool, project)
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
func GetLUNS(pool string, project string) *model.LUNS {
	luns := &model.LUNS{}
	fullurl := fmt.Sprintf("%s/storage/v1/pools/%s/projects/%s/luns", URL, pool, project)
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
	if !silent {
		version.PrintVersionInfo()
	} else {
		fmt.Println("getting version info.")
	}
	version.WriteCSV(Fs, dirname)
}

func getNetInterfaces() *model.NetInterfaces {
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
