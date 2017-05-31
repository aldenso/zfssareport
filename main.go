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

	"github.com/aldenso/zfssareport/model"
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
	chversion := make(chan *model.Version)
	chcluster := make(chan *model.Cluster)
	chchassis := make(chan *model.ChassisAll)
	chproblems := make(chan *model.Problems)
	chusers := make(chan *model.Users)
	go GetZFSSAVersion(chversion)
	go GetClusterInfo(chcluster)
	go GetChassis(chchassis)
	go GetProblems(chproblems)
	go GetUsers(chusers)
	version := <-chversion
	cluster := <-chcluster
	chassis := <-chchassis
	problems := <-chproblems
	users := <-chusers
	PrintZFSSAVersion(version, Fs)
	PrintZFSSACluster(cluster, Fs)
	PrintChassis(chassis, Fs)
	PrintProblems(problems, Fs)
	PrintUsers(users, Fs)
	chnetinterfaces := make(chan *model.NetInterfaces)
	chnetdatalinks := make(chan *model.NetDatalinks)
	chnetdevices := make(chan *model.NetDevices)
	chroutes := make(chan *model.Routes)
	go GetNetInterfaces(chnetinterfaces)
	go GetNetDatalinks(chnetdatalinks)
	go GetNetDevices(chnetdevices)
	go GetRoutes(chroutes)
	interfaces := <-chnetinterfaces
	datalinks := <-chnetdatalinks
	devices := <-chnetdevices
	routes := <-chroutes
	PrintNetInterfaces(interfaces, Fs)
	PrintNetDatalinks(datalinks, Fs)
	PrintNetDevices(devices, Fs)
	PrintRoutes(routes, Fs)
	chpools := make(chan *model.Pools)
	chprojects := make(chan *model.Projects)
	chfilesystems := make(chan *model.Filesystems)
	chluns := make(chan *model.LUNS)
	go GetPools(chpools)
	go GetProjects(chprojects)
	go GetFilesystems(chfilesystems)
	go GetLUNS(chluns)
	pools := <-chpools
	projects := <-chprojects
	filesystems := <-chfilesystems
	luns := <-chluns
	PrintPools(pools, Fs)
	PrintProjects(projects, Fs)
	PrintFilesystems(filesystems, Fs)
	PrintLUNS(luns, Fs)
	chfcinitiators := make(chan *model.FCInitiators)
	chfcinitiatorsgroups := make(chan *model.FCInitiatorGroups)
	chfctargets := make(chan *model.FCTargets)
	go GetFCInitiators(chfcinitiators)
	go GetFCInitiatorGroups(chfcinitiatorsgroups)
	go GetFCTargets(chfctargets)
	fcinitiators := <-chfcinitiators
	fcinitiatorsgroups := <-chfcinitiatorsgroups
	fctargets := <-chfctargets
	PrintFCInitiators(fcinitiators, Fs)
	PrintFCInitiatorGroups(fcinitiatorsgroups, Fs)
	PrintFCTargets(fctargets, Fs)
	chiscsiIs := make(chan *model.IscsiInitiators)
	chiscsiIGs := make(chan *model.IscsiInitiatorGroups)
	go GetIscsiInitiators(chiscsiIs)
	go GetIscsiInitiatorGroups(chiscsiIGs)
	iscsiIs := <-chiscsiIs
	iscsiIGs := <-chiscsiIGs
	PrintIscsiInitiators(iscsiIs, Fs)
	PrintIscsiInitiatorGroups(iscsiIGs, Fs)
	if err := utils.ZipDir(Fs, dirname); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n+++ results file '%s' created +++\n", dirname)
	fmt.Printf("\n############# DONE in %s #############\n", time.Since(NOW).String())
}
