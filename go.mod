module github.com/kiem-toan

go 1.16

replace root => ./

require (
	github.com/go-openapi/runtime v0.19.29
	github.com/go-resty/resty/v2 v2.4.0
	github.com/google/wire v0.5.0
	github.com/gorilla/mux v1.8.0
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/lib/pq v1.10.2
	github.com/mattn/go-colorable v0.1.8 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.11
)
