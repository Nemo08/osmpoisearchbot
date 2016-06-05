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
	getsection := c.CheckSection(section)
	if !getsection.HasKey(key) {
		log.Panic("'" + key + "' key not found")
	}

	stringkey, err := c.cfg.Section(section).GetKey(key)
	if err != nil {
		log.Panic("Key '" + key + "' not found ")
	}

	return stringkey.String()
}

func (c *IniConf) GetBoolKey(section, key string) bool {
	c.CheckSection(section)
	return c.cfg.Section(section).Key(key).MustBool(false)
}

func (c *IniConf) CheckSection(section string) *ini.Section {
	getsection, err := c.cfg.GetSection(section)
	if err != nil {
		log.Panic("Section '" + section + "' of configuration file not found ")
	}
	return getsection
}
