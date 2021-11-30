module github.com/kiem-toan

go 1.16

replace root => ./

require (
	github.com/bshuster-repo/logrus-logstash-hook v1.0.2
	github.com/casbin/casbin/v2 v2.1.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/gin v1.7.3
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-resty/resty/v2 v2.4.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.7.3
	github.com/hashicorp/consul/api v1.3.0
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/lib/pq v1.10.2
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/sirupsen/logrus v1.8.1
	github.com/subosito/gotenv v1.2.0
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210812204632-0ba0e8f03122
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.11
)
