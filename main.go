package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"time"

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

	template   bool
	configfile string
	dirname    string
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
	dirname = fmt.Sprintf("%s_%s", IP, strings.Replace(NOW.Format(time.RFC3339), ":", "", -1))
	if err := CreateDir(Fs, dirname); err != nil {
		log.Fatal(err)
	}
	pools := GetPools()
	PrintPools(pools, Fs)
	pmap := CreateMapPoolsProjects(pools)
	PrintProjects(pmap, Fs)
	allFS := CreateFSSlice(pmap)
	PrintFilesystems(allFS, Fs)
	PrintLUNS(pmap, Fs)
	fmt.Printf("############# DONE in %s #############\n", time.Since(NOW).String())
}
