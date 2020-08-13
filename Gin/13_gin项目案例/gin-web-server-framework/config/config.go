package config

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	cfg  *TomlConfig
	once sync.Once
)

/*TomlConfig ... */
type TomlConfig struct {
	Debug    bool
	DB       Database `toml:"database"`
	Server   server
	Clients  clients
	Wechat   wechat
	QyWechat qyWechat
}

//Mysql config
type Mysql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DbName   string `toml:"dbName"`
}

//Redis config
type Redis struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Password string `toml:"password"`
	DbName   int    `toml:"dbName"`
}

//Database info
type Database struct {
	Mysql Mysql `toml:"mysql"`
	Redis Redis `toml:"redis"`
}

type server struct {
	Port string
}

type wechat struct {
	WechatURL string
	Appid     string
	Secret    string
	GrantType string
}

// ThirdParty qywxthird party info
type ThirdParty struct {
	SuiteTokenURL string `toml:"suiteTokenURL"`
	SuiteID       string `toml:"suiteID"`
	SuitSecret    string `toml:"suitSecret"`
	TpURL         string `toml:"tpURL"`
}
type qyWechat struct {
	WechatURL   string     `toml:"wechatURL"`
	GetTokenURL string     `toml:"getTokenURL"`
	Corpid      string     `toml:"corpid"`
	Secret      string     `toml:"secret"`
	GrantType   string     `toml:"grantType"`
	ThirdParty  ThirdParty `toml:"thirdParty"`
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

/*Config ... */
func Config() *TomlConfig {
	once.Do(func() {
		filePath, err := filepath.Abs("./config/config.toml")
		if err != nil {
			panic(err)
		}
		log.Printf("parse toml file singleton. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}
