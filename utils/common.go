package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
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
	if password == "" {
		password = credentials()
	}
	url := fmt.Sprintf("https://%s:215/api", ip)
	return ip, user, password, url
}

func credentials() string {
	fmt.Printf("Enter Password:\n")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("Failed to read password: '%v'\n", err)
	}
	password := string(bytePassword)

	return strings.TrimSpace(password)
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
	file, err := fs.Create(filepath.Join(dirname, filename))
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

//ZipDir to create a zip file of the output directory.
func ZipDir(fs afero.Fs, source string) error {
	zipfile, err := fs.Create(source + ".zip")
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	info, err := fs.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})
	fs.RemoveAll(source)
	return err
}

//Cleanup to remove temporary directory for results.
func Cleanup(fs afero.Fs, source string) error {
	err := fs.RemoveAll(source)
	if err != nil {
		return err
	}
	return nil
}
