package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Packages []struct {
		Name                string `yaml:"name"`
		Path                string `yaml:"path"`
		Queries             string `yaml:"queries"`
		Schema              string `yaml:"schema"`
		Engine              string `yaml:"engine"`
		EmitJSONTags        bool   `yaml:"emit_json_tags"`
		EmitPreparedQueries bool   `yaml:"emit_prepared_queries"`
		EmitInterface       bool   `yaml:"emit_interface"`
		EmitExactTableNames bool   `yaml:"emit_exact_table_names"`
	} `yaml:"packages"`
}

func main() {
	config := Config{}
	config.Packages[0].Name = ""

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
