package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

const (
	filename = "settings.json"
)

type logConf struct {
	Path string
}

type apiConf struct {
	Host string
	Port int
}

type commonConf struct {
	Log logConf
	API apiConf
}

type mongoConf struct {
	Host   string
	Port   int
	Dbname string
}

type redisConf struct {
	Host string
	Port int
}

// Conf is configuration struct
type Conf struct {
	Common commonConf
	Mongo  mongoConf
	Redis  redisConf
}

// Configuration map
var Configuration Conf

func init() {
	log.Print("configuration init")

	_, selfname, _, _ := runtime.Caller(0)
	if settings, err := ioutil.ReadFile(filepath.Join(filepath.Dir(selfname), filename)); err != nil {
		log.Fatal(err.Error())
	} else {
		if err = json.Unmarshal(settings, &Configuration); err != nil {
			log.Fatal(err.Error())
		}
	}

}
