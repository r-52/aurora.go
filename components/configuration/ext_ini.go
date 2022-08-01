package configuration

import (
	"gopkg.in/ini.v1"
)

var Cfg *ini.File

func BuildIni(prod []byte, dev []byte) {
	Cfg, _ := ini.Load(prod, dev)
	Cfg.HasSection("")
}
