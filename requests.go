package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//GetPools get all pools.
func GetPools() Pools {
	var pools Pools
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
func GetProjects(pool string) Projects {
	var projects Projects
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
func GetFilesystems(pool string, project string) Filesystems {
	var filesystems Filesystems
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
func GetLUNS(pool string, project string) LUNS {
	var luns LUNS
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
