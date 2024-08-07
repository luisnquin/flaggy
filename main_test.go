package flaggy_test

import (
	"os"
	"testing"

	"github.com/luisnquin/flaggy"
)

func TestMain(m *testing.M) {
	flaggy.PanicInsteadOfExit = true
	// flaggy.DebugMode = true
	os.Exit(m.Run())
}

func TestSetDescription(t *testing.T) {
	desc := "Test Description"
	flaggy.SetDescription(desc)
	if flaggy.DefaultParser.Description != desc {
		t.Fatal("set description does not match")
	}
}

func TestSetVersion(t *testing.T) {
	ver := "Test Version"
	flaggy.SetVersion(ver)
	if flaggy.DefaultParser.Version != ver {
		t.Fatal("set version does not match")
	}
}

func TestParserWithNoArgs(t *testing.T) {
	os.Args = []string{}
	flaggy.ResetParser()
}

func TestSetName(t *testing.T) {
	name := "Test Name"
	flaggy.SetName(name)
	if flaggy.DefaultParser.Name != name {
		t.Fatal("set name does not match")
	}
}
