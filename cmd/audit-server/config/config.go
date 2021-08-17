package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/kiem-toan/infrastructure/redis"

	"github.com/kiem-toan/infrastructure/integration/email"

	"gopkg.in/yaml.v2"
)

var (
	flConfigFile = ""
	flConfigYaml = ""
	flExample    = false
	flNoEnv      = false
)

// Default ...
func Default() Config {
	cfg := Config{
		Databases: DBConfig{
			Postgres: DefaultPostgres(),
		},
		Env:  "dev",
		Port: "8080",
		Email: email.SMTPConfig{
			Host:        "smtp.gmail.com",
			Port:        587,
			Username:    "shinichi24567@gmail.com",
			Password:    "shinichi24567123",
			Encrypt:     "tls",
			FromAddress: "",
		},
	}
	return cfg
}

// Load loads config from file
func Load() (Config, error) {
	var cfg, defCfg Config
	defCfg = Default()
	err := LoadWithDefault(&cfg, defCfg)
	if err != nil {
		return cfg, err
	}
	return cfg, err
}

// Config ...
type Config struct {
	Databases DBConfig         `yaml:",inline"`
	Env       string           `yaml:"env"`
	Port      string           `yaml:"port"`
	Email     email.SMTPConfig `yaml:"email"`
	redis     Redis            `yaml:"redis"`
}

type ConfigPostgres struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`
	Timeout  int    `yaml:"timeout"`

	MaxOpenConns    int `yaml:"max_open_conns"`
	MaxIdleConns    int `yaml:"max_idle_conns"`
	MaxConnLifetime int `yaml:"max_conn_lifetime"`

	GoogleAuthFile string `yaml:"google_auth_file"`
}

type DBConfig struct {
	Postgres ConfigPostgres `yaml:"postgres"`
}

// DefaultPostgres ...
func DefaultPostgres() ConfigPostgres {
	return ConfigPostgres{
		Protocol:       "",
		Host:           "localhost",
		Port:           5432,
		Username:       "postgres",
		Password:       "postgres",
		Database:       "postgres",
		SSLMode:        "disable",
		Timeout:        15,
		GoogleAuthFile: "",
	}
}

type Redis = redis.Redis

// DefaultRedis ...
func DefaultRedis() Redis {
	return Redis{
		Host:     "redis",
		Port:     "6379",
		Username: "",
		Password: "",
	}
}
func LoadWithDefault(v, def interface{}) (err error) {
	defer func() {
		if flExample {
			if err != nil {
				//ll.Fatal("Error while loading config", l.Error(err))
			}
			//PrintExample(v)
			os.Exit(2)
		}
	}()
	if (flConfigFile != "") && (flConfigYaml != "") {
		//return errors.New("must provide only -config-file or -config-yaml")
	}
	if flConfigFile != "" {
		err = LoadFromFile(flConfigFile, v)
		if err != nil {
			log.Fatal("can not load config from file: %v (%v)", flConfigFile, err)
		}
		return err
	}
	if flConfigYaml != "" {
		return LoadFromYaml([]byte(flConfigYaml), v)
	}
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(def))
	return nil
}

// LoadFromFile loads config from file
func LoadFromFile(configPath string, v interface{}) (err error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	return LoadFromYaml(data, v)
}

func LoadFromYaml(input []byte, v interface{}) (err error) {
	return yaml.Unmarshal(input, v)
}
func InitFlags() {
	flag.StringVar(&flConfigFile, "config-file", "server_config.yaml", "Path to config file")
	flag.StringVar(&flConfigYaml, "config-yaml", "", "Config as yaml string")
	flag.BoolVar(&flNoEnv, "no-env", false, "Don't read config from environment")
	flag.BoolVar(&flExample, "example", false, "Print example config then exit")
}

func ParseFlags() {
	flag.Parse()
}
