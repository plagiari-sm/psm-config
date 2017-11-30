package psmconfig

import (
	"flag"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

// Config Global Variable
// Store Configuration Options
var Config *conf

type conf struct {
	Name  string `yaml:"name,omitempty"`
	Env   string `yaml:"env,omitempty"`
	Host  string `yaml:"host,omitempty"`
	Port  int    `yaml:"port,omitempty"`
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
	Twitter struct {
		ConsumerKey       string `yaml:"consumer-key,omitempty"`
		ConsumerSecret    string `yaml:"consumer-secret,omitempty"`
		AccessToken       string `yaml:"access-token,omitempty"`
		AccessTokenSecret string `yaml:"access-token-secret,omitempty"`
	} `yaml:"twitter,omitempty"`
	Log   struct{} `yaml:"log,omitempty"`
	Proto struct{} `yaml:"proto,omitempty"`
	API   struct{} `yaml:"api,omitempty"`
}

// NewConfig Init
func NewConfig() {
	var path string
	flag.StringVar(&path, "config", "conf/default.yaml", "Parse Configuration File")
	flag.Parse()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("[PSM-CONFIG] Configuration parsing error: %v", err)
	}

	yaml.Unmarshal(data, &Config)
	log.Infof("[PSM-CONFIG] Configuration loaded: %v", Config)
}
