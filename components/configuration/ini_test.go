package configuration_test

import (
	"testing"

	"github.com/r-52/aurora/components/configuration"
)

func Test_ConfigurationNew(t *testing.T) {
	config := configuration.New()
	if config == nil {
		t.Error("got no config")
	}
}

func Test_IniFromStringEmptyLines(t *testing.T) {
	config := configuration.New()
	if config == nil {
		t.Error("got no config")
	}

	str := `
	test=1



`
	err := config.From(&str)
	if err != nil {
		t.Error(err)
	}
}
