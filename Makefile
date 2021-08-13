check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swaggercmd/swagger

swagger: check_install
	GO111MODULE=off swagger generate spec -o ./swagger.json --scan-models