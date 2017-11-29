package main

import (
	"flag"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

// Config Global Variable
// Store Configuration Options
var Config *conf

type conf struct {
	Name string `yaml:"name,omitempty"`
	Env  string `yaml:"env,omitempty"`
	Host string `yaml:"host,omitempty"`
	Port int    `yaml:"port,omitempty"`
	DB   struct {
		Mongo struct {
			Host string `yaml:"host,omitempty"`
			Port int    `yaml:"port,omitempty"`
			User string `yaml:"user,omitempty"`
			Pass string `yaml:"pass,omitempty"`
			Path string `yaml:"path,omitempty"`
		} `yaml:"mongo,omitempty"`
		ES struct {
			Host string `yaml:"host,omitempty"`
			Port int    `yaml:"port,omitempty"`
			User string `yaml:"user,omitempty"`
			Pass string `yaml:"pass,omitempty"`
			Path string `yaml:"path,omitempty"`
		} `yaml:"es,omitempty"`
		Redis struct {
			Host string `yaml:"host,omitempty"`
			Port int    `yaml:"port,omitempty"`
			User string `yaml:"user,omitempty"`
			Pass string `yaml:"pass,omitempty"`
			Path string `yaml:"path,omitempty"`
		} `yaml:"redis,omitempty"`
	} `yaml:"db,omitempty"`
	Log   struct{} `yaml:"log,omitempty"`
	Proto struct{} `yaml:"proto,omitempty"`
	API   struct{} `yaml:"api,omitempty"`
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
func main() {
	var path string
	flag.StringVar(&path, "config", "conf/dev.yaml", "Parse Configuration File")
	flag.Parse()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("[PSM-CONFIG] Configuration parsing error: %v", err)
	}

	yaml.Unmarshal(data, &Config)
	if Config.Env == "development" {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	}
	log.Infof("[PSM-CONFIG] Configuration loaded: %v", Config)
}
