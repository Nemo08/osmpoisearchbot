package main

import (
	"log"

	"github.com/go-ini/ini"
)

type IniConf struct {
	confpath string
	cfg      ini.File
}

func (c *IniConf) CheckAndLoadConf(path string) {
	c.confpath = path
	cfg, err := ini.LooseLoad(path)
	c.cfg = *cfg

	if err != nil {
		log.Panic("Configuration file '"+path+"' not found ", err)
	}

}

func (c *IniConf) GetStringKey(section, key string) string {
	getsection, err := c.cfg.GetSection(section)
	if err != nil {
		log.Panic("Section '"+section+"' of configuration file not found ", err)
	}

	if !getsection.HasKey(key) {
		log.Panic("'"+key+"' key not found", err)
	}

	stringkey, err := c.cfg.Section(section).GetKey(key)
	if err != nil {
		log.Panic("Key '"+key+"' not found ", err)
	}

	return stringkey.String()
}
