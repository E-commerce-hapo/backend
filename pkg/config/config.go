package config

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"

	"github.com/kiem-toan/pkg/env"
)

const (
	Port = 8080
)

var (
	appConfig = &Config{}
)

// App Config ...
type Config struct {
	ProjectDir      string
	ApplicationName string      `json:"application_name"`
	Databases       DBConfig    `json:"databases"`
	Log             Log         `json:"log"`
	LogStash        LogStash    `json:"log_stash"`
	Zipkin          Zipkin      `json:"zipkin"`
	Env             env.EnvType `json:"env"`
	Consul          Consul      `json:"consul"`
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

type Zipkin struct {
	URL string `json:"url"`
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

func GetAppConfig() *Config {
	return appConfig
}

func SetAppConfig(cfg *Config) {
	appConfig = cfg
}

func (c *Config) Info() {
	fmt.Println("Thông số biến môi trường:")
	pp.Println(c)
}

func (c *Config) AssignEnv() {
	if os.Getenv("APPLICATION_NAME") != "" {
		c.ApplicationName = os.Getenv("APPLICATION_NAME")
	}
	if os.Getenv("CONSUL_IP") != "" {
		c.Consul.IP = os.Getenv("CONSUL_IP")
	}
	if os.Getenv("CONSUL_PORT") != "" {
		c.Consul.Port = os.Getenv("CONSUL_PORT")
	}
	if os.Getenv("CONSUL_ACL_TOKEN") != "" {
		c.Consul.ACLToken = os.Getenv("CONSUL_ACL_TOKEN")
	}

	if os.Getenv("CONSUL_ACL_TOKEN") != "" {
		c.Consul.ACLToken = os.Getenv("CONSUL_ACL_TOKEN")
	}

	if os.Getenv("CONSUL_ACL_TOKEN") != "" {
		c.Consul.ACLToken = os.Getenv("CONSUL_ACL_TOKEN")
	}

	if os.Getenv("CONSUL_ACL_TOKEN") != "" {
		c.Consul.ACLToken = os.Getenv("CONSUL_ACL_TOKEN")
	}

	if os.Getenv("CONSUL_ACL_TOKEN") != "" {
		c.Consul.ACLToken = os.Getenv("CONSUL_ACL_TOKEN")
	}
}

func DefaultConfig() *Config {
	return &Config{
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
		Zipkin: Zipkin{
			URL: "",
		},
	}
}
