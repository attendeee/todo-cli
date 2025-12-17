package config

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

var defaultSqlDatabaseName string = `.task.db`

var defaultSqlDatabasePath string = xdg.Home
var ConfigPath string = fmt.Sprintf("%s/%s", xdg.ConfigHome, `task/config.yaml`)
var HomeDir string = xdg.DataHome

type config struct {
	SqlDatabasePath string `yaml:"database-path"` // there is no defaults because we define
	SqlDatabaseName string `yaml:"database-name"` // it by external library
	TemplateName    string `yaml:"template-name" default:"tmp.yaml"`
	TableStyle      string `yaml:"table-style" default:"default"`
	SeparateColumns bool   `yaml:"separate-columns" default:"true"`
	SeparateRows    bool   `yaml:"separate-rows" default:"true"`
}

var c *config

func init() {
	c = &config{
		SqlDatabasePath: defaultSqlDatabasePath,
		SqlDatabaseName: defaultSqlDatabaseName,
	}
	defaults.Set(c)

}

func MustInitConfigFile() {
	b, err := openFile()
	if err != nil {
		panic(fmt.Sprintf("Unable to open config file: %s\n", err.Error()))
	}

	err = parseFile(b)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse config file: %s\n", err.Error()))
	}
}

func GetSqlPath() string {
	return fmt.Sprintf("%s/%s", c.SqlDatabasePath, c.SqlDatabaseName)

}

func openFile() ([]byte, error) {
	b, err := os.ReadFile(ConfigPath)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func parseFile(b []byte) error {
	return yaml.Unmarshal(b, &c)

}

func GetConfig() *config {
	return c
}
