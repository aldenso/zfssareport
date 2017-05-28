package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"time"

	"github.com/aldenso/zfssareport/utils"
	"github.com/aldenso/zfssareport/zfssareportfs"
	"github.com/spf13/afero"
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

func cleanexit(Fs afero.Fs, dirname string) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Signal received")
		err := utils.Cleanup(Fs, dirname)
		if err != nil {
			fmt.Printf("failed to remove directory %s: %v\n", dirname, err)
		} else {
			fmt.Printf("removed directory %s\n", dirname)
		}
		os.Exit(1)
	}()
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
	cleanexit(Fs, dirname)
	getZFSSAVersion()
	getClusterInfo()
	allchassis := GetChassis()
	PrintChassis(allchassis, Fs)
	problems := GetProblems()
	PrintProblems(problems, Fs)
	interfaces := getNetInterfaces()
	PrintNetInterfaces(interfaces, Fs)
	datalinks := GetNetDatalinks()
	PrintNetDatalinks(datalinks, Fs)
	devices := GetNetDevices()
	PrintNetDevices(devices, Fs)
	pools := GetPools()
	PrintPools(pools, Fs)
	projects := GetProjects()
	PrintProjects(projects, Fs)
	filesystems := GetFilesystems()
	PrintFilesystems(filesystems, Fs)
	luns := GetLUNS()
	PrintLUNS(luns, Fs)
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
	fmt.Printf("\n+++ results file '%s' created +++\n", dirname)
	fmt.Printf("\n############# DONE in %s #############\n", time.Since(NOW).String())
}
