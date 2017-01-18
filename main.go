package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
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
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	IP = viper.GetString("ip")
	USER = viper.GetString("user")
	PASSWORD = viper.GetString("password")
	URL = fmt.Sprintf("https://%s:215/api", IP)
}

func main() {
	PrintPools()
	PrintProjects()
	PrintFilesystems()
	PrintLUNS()
	fmt.Println("DONE!")
}
