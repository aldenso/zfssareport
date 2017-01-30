package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"time"

	"github.com/aldenso/zfssareport/utils"
	"github.com/aldenso/zfssareport/zfssareportfs"
)

var (
	//IP for zfssa REST api.
	IP string
	//USER for zfssa authentication.
	USER string
	//PASSWORD for zfssa authentication.
	PASSWORD string
	//URL for zfssa REST api.
	URL string
	// Fs afero fs to help later with testing.
	Fs = zfssareportfs.InitOSFs()

	//NOW for timestamp
	NOW = time.Now()

	template, silent bool
	configfile       = "config.yml"
	dirname          string
)

func init() {
	flag.BoolVar(&template, "template", false, "Create an example config.yml file.")
	flag.StringVar(&configfile, "t", "config.yml", "Specify a config file.")
	flag.BoolVar(&silent, "silent", false, "Do not print info, only create the csv outputs in zip file.")
}

func main() {
	flag.Parse()
	if template {
		msg, err := utils.CreateTemplate(Fs, configfile)
		if err != nil {
			log.Fatalf("Error creating tomfile: %v", err)
		} else {
			fmt.Println(msg)
			os.Exit(0)
		}
	}
	IP, USER, PASSWORD, URL = utils.ReadConfigFile(configfile)
	dirname = fmt.Sprintf("%s_%s", IP, strings.Replace(NOW.Format(time.RFC3339), ":", "", -1))
	if err := utils.CreateDir(Fs, dirname); err != nil {
		log.Fatal(err)
	}
	getZFSSAVersion()
	getClusterInfo()
	allchassis := GetChassis()
	PrintChassis(allchassis, Fs)
	problems := GetProblems()
	PrintProblems(problems, Fs)
	interfaces := getNetInterfaces()
	PrintNetInterfaces(interfaces, Fs)
	devices := GetNetDevices()
	PrintNetDevices(devices, Fs)
	pools := GetPools()
	PrintPools(pools, Fs)
	pmap := CreateMapPoolsProjects(pools)
	PrintProjects(pmap, Fs)
	allFS := CreateFSSlice(pmap)
	PrintFilesystems(allFS, Fs)
	allLUNS := CreateLUNSSlice(pmap)
	PrintLUNS(allLUNS, Fs)
	fcinitiators := GetFCInitiators()
	PrintFCInitiators(fcinitiators, Fs)
	fcinitiatorgroups := GetFCInitiatorGroups()
	PrintFCInitiatorGroups(fcinitiatorgroups, Fs)
	fctargets := GetFCTargets()
	PrintFCTargets(fctargets, Fs)
	iscsiinitiators := GetIscsiInitiators()
	PrintIscsiInitiators(iscsiinitiators, Fs)
	iscsiinitiatorgroups := GetIscsiInitiatorGroups()
	PrintIscsiInitiatorGroups(iscsiinitiatorgroups, Fs)
	if err := utils.ZipDir(Fs, dirname); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("############# DONE in %s #############\n", time.Since(NOW).String())
}
