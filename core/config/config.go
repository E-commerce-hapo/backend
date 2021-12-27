package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/k0kubun/pp"

	"github.com/kiem-toan/pkg/env"
)

const (
	ServerPort = "8080"
)

var (
	appConfig = Config{}
)

// App Config ...
type Config struct {
	ProjectDir      string
	ApplicationName string
	Databases       DBConfig `json:"databases"`
	Log             Log      `json:"log"`
	LogStash        LogStash
	Env             env.EnvType `json:"env"`
	ServerPort      string
}

type Consul struct {
	IP       string `json:"ip"`
	Port     string `json:"port"`
	ACLToken string `json:"acl_token"`
}

type Log struct {
	Level string `json:"level"`
}

type LogStash struct {
	Port string `json:"port"`
	IP   string `json:"ip"`
}

type DBConfig struct {
	PostgresDB PostgresConfig `json:"postgres_db"`
}

type PostgresConfig struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode"`
	Timeout  int    `json:"timeout"`

	MaxOpenConns    int `json:"max_open_conns"`
	MaxIdleConns    int `json:"max_idle_conns"`
	MaxConnLifetime int `json:"max_conn_lifetime"`

	GoogleAuthFile string `json:"google_auth_file"`
}

func GetAppConfig() Config {
	return appConfig
}

func SetAppConfig(cfg Config) {
	appConfig = cfg
	appConfig.assignEnv()
	appConfig.Info()
}

func (c *Config) Info() {
	fmt.Println("Thông số biến môi trường:")
	pp.Println(c)
}

func (c *Config) assignEnv() {
	if os.Getenv("APPLICATION_NAME") != "" {
		c.ApplicationName = os.Getenv("APPLICATION_NAME")
	}
	if os.Getenv("LOGSTASH_IP") != "" {
		c.LogStash.IP = os.Getenv("LOGSTASH_IP")
	}
	if os.Getenv("LOGSTASH_PORT") != "" {
		c.LogStash.Port = os.Getenv("LOGSTASH_PORT")
	}

	c.ServerPort = ServerPort

}

func Asset(name string) ([]byte, error) {
	base := filepath.Join(appConfig.ProjectDir, "core/config")
	if strings.Contains(name, "..") {
		panic(fmt.Sprintf("invalid name (%v)", name))
	}
	return ioutil.ReadFile(filepath.Join(base, name))
}

func DefaultConfig() *Config {
	return &Config{
		ProjectDir:      "",
		ApplicationName: "",
		Databases: DBConfig{
			PostgresDB: PostgresConfig{
				Protocol:        "",
				Host:            "",
				Port:            0,
				Username:        "",
				Password:        "",
				Database:        "",
				SSLMode:         "",
				Timeout:         0,
				MaxOpenConns:    0,
				MaxIdleConns:    0,
				MaxConnLifetime: 0,
				GoogleAuthFile:  "",
			},
		},
		Log: Log{
			Level: "",
		},
		LogStash: LogStash{
			Port: "",
			IP:   "",
		},
		Env:        0,
		ServerPort: "",
	}
}
