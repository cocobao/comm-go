package com_conf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func YamlFromFile(cfg interface{}, path ...string) error {
	var b bytes.Buffer

	for _, p := range path {
		data, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}
		b.Write(data)
	}

	if yerr := yaml.Unmarshal(b.Bytes(), cfg); yerr != nil {
		return yerr
	}
	return nil
}

func JsonFromFile(cfg interface{}, rootpath string) error {
	data, err := ioutil.ReadFile(rootpath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, cfg)
}
