package core

import (
	"code_memory/config"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func Config() *config.Server {
	yamlFile, err := ioutil.ReadFile(`config.yaml`)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	c := &config.Server{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("Unmarshal: %v\n", err)
	}
	return c
}
