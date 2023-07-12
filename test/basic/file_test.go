package main

import (
	"flag"
	"github.com/go-ini/ini"
	"testing"
)

var (
	conf = flag.String("conf", "../../conf/dev.ini", "conf")
)

func TestIni(t *testing.T) {
	flag.Parse()
	cfg, err := ini.Load(*conf)
	if err != nil {
		println(err.Error())
	}
	println(cfg.SectionStrings())
}
