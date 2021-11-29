package env

type EnvType int

var env EnvType
var isDev = false
var notProd = true

// Environment constants
const (
	EnvDev EnvType = iota + 1
	EnvStaging
	EnvProd
)

var envNames = map[EnvType]string{
	EnvDev:     "dev",
	EnvStaging: "staging",
	EnvProd:    "prod",
}

func (e EnvType) String() string {
	return envNames[e]
}

var envValues = map[string]EnvType{
	"dev":     EnvDev,
	"staging": EnvStaging,
	"prod":    EnvProd,
}

func Env() EnvType {
	return env
}

func IsDev() bool {
	return env == EnvDev
}

func IsProd() bool {
	return env == EnvProd
}

func IsStaging() bool {
	return env == EnvStaging
}

func SetEnvironment(e string) EnvType {
	if env != 0 {
		panic("Already initialize environment")
	}
	env = envValues[e]
	switch env {
	case EnvDev:
		isDev = true
	case EnvStaging:
	case EnvProd:
		notProd = false
	default:
		panic("invalid environment: " + e)
	}
	return env
}
