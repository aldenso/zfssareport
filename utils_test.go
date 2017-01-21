package main

import (
	"testing"

	"github.com/aldenso/zfssareport/zfssareportfs"
	"github.com/spf13/afero"
)

var FsMem = zfssareportfs.InitMemFs()

func Test_CreateTemplate(t *testing.T) {
	msg, err := CreateTemplate(FsMem)
	if err != nil {
		t.Errorf("Expected not nil for CreateTemplate, got '%s'\n", err)
	}
	if msg != configcreated {
		t.Errorf("Expected template '%s', got '%s'\n", configcreated, msg)
	}
	content, err := afero.ReadFile(FsMem, "config.yml")
	if err != nil {
		t.Error("Failed to read config file")
	}
	if string(content) != templatecontent {
		t.Errorf("Expected '%s', got '%s'\n", templatecontent, string(content))
	}
	//Test config
	msg, err = CreateTemplate(FsMem)
	if err != nil {
		t.Errorf("Expected not nil for CreateTemplate, got '%s'\n", err)
	}
	if msg != configexists {
		t.Errorf("Expected template '%s', got '%s'\n", configexists, msg)
	}
}
