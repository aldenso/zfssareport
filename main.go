package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	//POOLS pools in zfssa.
	POOLS Pools
	//POOLSPROJECTS map for projects in pools.
	POOLSPROJECTS = make(map[string][]string)

	// Fs afero fs to help later with testing.
	Fs = zfssareportfs.InitOSFs()

	template   bool
	configfile string
)

func init() {
	flag.BoolVar(&template, "template", false, "Create an example config.yml file.")
	flag.StringVar(&configfile, "t", "config.yml", "Specify a config file.")
}

func main() {
	flag.Parse()
	if template {
		msg, err := CreateTemplate(Fs)
		if err != nil {
			log.Fatalf("Error creating tomfile: %v", err)
		} else {
			fmt.Println(msg)
			os.Exit(0)
		}
	}
	ReadConfigFile()
	PrintPools()
	PrintProjects()
	PrintFilesystems()
	PrintLUNS()
	fmt.Println("DONE!")
}
