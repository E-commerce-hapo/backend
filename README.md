# Thương mại điện tử Ecommerce
Tạo mới token HOTP, validate HOTP
## Tài liệu tham khảo để code
{ \
&nbsp; &nbsp; "databases": {\
&nbsp; &nbsp; &nbsp; &nbsp; "postgres_db": {\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "protocol":"localhost",\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "host":"localhost",\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "port":5432,\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "username":"postgres",\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "password":"postgres",\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "database":"postgres",\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "sslmode": "disable",\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "timeout": 15,\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "max_open_conns": 0,\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "max_idle_conns": 0,\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "max_conn_lifetime": 0,\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "google_auth_file": ""\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; }\
&nbsp; &nbsp; &nbsp; &nbsp;},\
&nbsp; &nbsp; "log": {\
&nbsp; &nbsp; &nbsp; &nbsp; "level": "info"\
&nbsp; &nbsp; },\
&nbsp; &nbsp; "zipkin": {\
&nbsp; &nbsp; &nbsp; &nbsp; "url": "http://10.90.68.35:30208"\
&nbsp; &nbsp; },\
}

+ LOG_LEVEL : info, fatal , panic, warn, debug . Mặc định là error
+ Khi thay đổi bộ thông số trên consul config này thì app sẽ không phải reset mà sẽ tự load lại biến môi trường
+ Điều này cũng tương tự khi sử dụng với k8s và nomad, khi thay đổi biến môi trường thì các job sẽ không cần phải restart lại
## Build with docker
- Để build với docker thì trước tiên phải cài đặt docker

- Sử dụng tiếp các câu lệnh sau:
+ docker login 10.91.120.43:8000 sau đó sử dụng user/pass để login admin/admin123
+ docker repo.mafc.vn:8000 sau đó sử dụng user/pass để login
### Build image
docker build --rm -f Dockerfile -t {name image}:{version} .
+ Example : docker build --rm -f Dockerfile -t golang-hotp:v1.2.7 .
### Tag image
docker tag {name image}:{version}  {repo address}/{name image}:{version}
+ Example : docker tag golang-hotp:v1.2.7 10.91.120.43:8000/repository/mobile-project/golang-hotp:v1.2.7
### Push image
docker push {repo address}/{name image}:{version} đường dẫn vừa tag bên trên
+ Example : docker push 10.91.120.43:8000/repository/mobile-project/golang-hotp:v1.2.7
### Pull image
docker pull  {repo address}/{name image}:{version} đường dẫn vừa push bên trên
+ Example: docker pull 10.91.120.43:8000/repository/mobile-project/golang-hotp:v1.2.7
### Run with docker
docker run -d -p 8080:8080 -e {env name}={ env value} {repo address}/{name image}:{version}

+ Example:
+ docker run -d -p 8080:8080 -e CONSUL_IP=10.91.120.55 -e CONSUL_PORT=8500 -e LOGSTASH_IP=10.90.68.35 -e LOGSTASH_PORT=30204 -e CONSUL_ACL_TOKEN=7caf93ca-2112-2f84-3bc9-39e812983ed1 -e APPLICATION_NAME=otp-api -e SERVER_PORT=8080 10.91.120.43:8000/repository/mobile-project/golang-hotp:v1.2.7

+Lưu ý: FIXED_OTP chỉ có giá trị khi ở môi trường dev/test, đây là giá trị otp luôn trả kết quả thành công
### Run with main.go
+ cd vào project
+ Gõ lệnh go run main.go

### Các lỗi gặp phải khi cài đặt
- "gcc": executable file not found in %PATH% : https://stackoverflow.com/questions/43580131/exec-gcc-executable-file-not-found-in-path-when-trying-go-build (lỗi này ở window)
-  swag: command not found :
    + https://github.com/swaggo/swag/issues/209
    + go get -v -u github.com/swaggo/swag/cmd/swag
    + export PATH=$PATH:$HOME/go/bin
    + swag init
    + truy cập swag tại máy localhost thì truy cập vào link  http://localhost:8080/swagger/index.html

## Author
## tiendc@mafc.com.vn



