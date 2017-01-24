package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

const (
	configexists    = "config.yml already exist."
	configcreated   = "config.yml created."
	templatecontent = `# ZFSSA REPORT CONFIG"
ip: 192.168.56.150
user: root
password: password`
)

//Header print header for output.
func Header(msg string) {
	fmt.Println("#####################################################################################################################")
	fmt.Printf("## %-111s ##\n", msg)
	fmt.Println("#####################################################################################################################")
}

// CreateTemplate fucntion to create a base remgo.toml file
func CreateTemplate(fs afero.Fs, configfile string) (string, error) {
	if _, err := fs.Stat(configfile); err != nil {
		if os.IsNotExist(err) {
			file, err := fs.Create(configfile)
			if err != nil {
				return "", err
			}
			defer file.Close()
			if _, err := file.Write([]byte(templatecontent)); err != nil {
				return "", err
			}
		}
	} else {
		return configexists, nil
	}
	return configcreated, nil
}

//ReadConfigFile read yaml config file for zfssa.
func ReadConfigFile(configfile string) (string, string, string, string) {
	viper.SetConfigFile(configfile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	ip := viper.GetString("ip")
	user := viper.GetString("user")
	password := viper.GetString("password")
	url := fmt.Sprintf("https://%s:215/api", ip)
	return ip, user, password, url
}

//CreateDir create directory for collected info.
func CreateDir(fs afero.Fs, name string) error {
	err := fs.Mkdir(name, 0750)
	if err != nil {
		return err
	}
	return nil
}

//CreateFile create file for info dump.
func CreateFile(fs afero.Fs, dirname string, filename string) (afero.File, error) {
	file, err := fs.Create(dirname + "/" + filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

//CloseFile close file when done.
func CloseFile(file afero.File) error {
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}
