package utils

import (
	"testing"

	"fmt"

	"github.com/aldenso/zfssareport/zfssareportfs"
	"github.com/spf13/afero"
)

var (
	FsMem          = zfssareportfs.InitMemFs()
	Fs             = zfssareportfs.InitOSFs()
	tempIP         = "192.168.56.150"
	tempUser       = "root"
	tempPassword   = "password"
	dirtest        = fmt.Sprintf("%s_2017-01-21T220321-0400", tempIP)
	testconfigfile = "../config.yml"
	filetestname   = "poolstest.csv"
	filetest       afero.File
)

func Test_CreateTemplate(t *testing.T) {
	msg, err := CreateTemplate(FsMem, testconfigfile)
	if err != nil {
		t.Errorf("Expected not nil for CreateTemplate, got '%s'\n", err)
	}
	if msg != configcreated {
		t.Errorf("Expected template '%s'\n, got '%s'\n", configcreated, msg)
	}
	//Test config
	msg, err = CreateTemplate(FsMem, testconfigfile)
	if err != nil {
		t.Errorf("Expected not nil for CreateTemplate, got '%s'\n", err)
	}
	if msg != configexists {
		t.Errorf("Expected template '%s'\n, got '%s'\n", configexists, msg)
	}
}

func Test_ReadConfigFile(t *testing.T) {
	ip, user, password, _ := ReadConfigFile(testconfigfile)
	if ip != tempIP {
		t.Errorf("Expected '%s', got '%s'\n", tempIP, ip)
	}
	if user != tempUser {
		t.Errorf("Expected '%s', got '%s'\n", tempUser, user)
	}
	if password != tempPassword {
		t.Errorf("Expected '%s', got '%s'\n", tempPassword, password)
	}
}

func Test_CreateDir(t *testing.T) {
	if err := CreateDir(FsMem, dirtest); err != nil {
		t.Errorf("Failed to create dir, Error: '%v'\n", err)
	}
	ok, _ := afero.DirExists(FsMem, dirtest)
	if !ok {
		t.Errorf("CreateDir did not create dir: '%s'\n", dirtest)
	}
}

func Test_CreateFile(t *testing.T) {
	var err error
	filetest, err = CreateFile(FsMem, dirtest, filetestname)
	if err != nil {
		t.Errorf("Failed to create test file '%s' with error: '%v'\n", dirtest+"/"+filetestname, err)
	}
	ok, _ := afero.Exists(FsMem, dirtest+"/"+filetestname)
	if !ok {
		t.Errorf("CreateFile did not create file: '%s'\n", dirtest+"/"+filetestname)
	}
}

func Test_CloseFile(t *testing.T) {
	if err := CloseFile(filetest); err != nil {
		t.Errorf("CloseFile did not close file: '%s', err: '%v'\n", dirtest+"/"+filetestname, err)
	}
}

func Test_ZipDir(t *testing.T) {
	if err := ZipDir(FsMem, dirtest); err != nil {
		t.Errorf("Failed to create zip file: '%v'", err)
	}
}
