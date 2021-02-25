package appconfig

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func GetValue(paramName string) string {

	var config Config
	filename, _ := filepath.Abs("appconfig.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//fmt.Print(yamlFile)
	// var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	switch strings.ToLower(paramName) {
	case "database.host":
		return config.Database.Host
	case "database.port":
		return config.Database.Port
	case "database.db_name":
		return config.Database.Dbname
	case "database.login":
		return config.Database.Login
	case "database.password":
		return config.Database.Password
	}
	return "dd"

}

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Dbname   string `yaml:"dbname"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}
