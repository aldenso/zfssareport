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

//GetPools get all pools.
func GetPools() model.Pools {
	var pools model.Pools
	HTTPClientCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: HTTPClientCfg, Timeout: 30 * time.Second}
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
	HTTPClientCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: HTTPClientCfg, Timeout: 60 * time.Second}
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
	HTTPClientCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: HTTPClientCfg, Timeout: 60 * time.Second}
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
	HTTPClientCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: HTTPClientCfg, Timeout: 60 * time.Second}
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
	HTTPClientCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: HTTPClientCfg, Timeout: 60 * time.Second}
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
	version.PrintVersionInfo()
	version.WriteCSV(Fs, dirname)
}

func getNetInterfaces() *model.NetInterfaces {
	interfaces := &model.NetInterfaces{}
	HTTPClientCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: HTTPClientCfg, Timeout: 60 * time.Second}
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
	HTTPClientCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: HTTPClientCfg, Timeout: 60 * time.Second}
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
	HTTPClientCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: HTTPClientCfg, Timeout: 60 * time.Second}
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
