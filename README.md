## Install
* Prepare your database (create a new database and role)
* Clone or download repo
* Go into project directory and install dependencies
```shell script
$ go mod download
```
* Install [protobuf compiler](https://github.com/google/protobuf/blob/master/README.md#protocol-compiler-installation) from binary or install from apt
```shell script
sudo apt install protobuf-compiler
```
* Install the protoc Go plugin
 ```shell script
 $ go get -u github.com/golang/protobuf/protoc-gen-go
 ```
* (Optional) Rebuild the generated Go code
```shell script
cd grpc-boilerplate
protoc --go_out=plugins=grpc:. api/*.proto
```
* Install sql-migrate
```shell script
$ go get -v github.com/rubenv/sql-migrate/...
```
* Setup environment variables for your database
```shell script
export DB_HOST=<YOUR HOST>
export DB_NAME=<YOUR NAME OF DATABASE>
export DB_USER=<YOUR NAME OF ROLE>
export DB_PASSWORD=<YOUR ROLE PASSWORD>
```
* Run sql-migrate and check status of migrations
```shell script
$ sql-migrate up
$ sql-migrate status
```
* Install sqlboiler and driver for your database
```shell script
$ go get -u -t github.com/volatiletech/sqlboiler
$ go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
```
* Create sqlboiler configuration file from example (sqlboiler.toml.example)
```toml
output           = "models"
wipe             = true
no_tests         = false
no_context       = true
add_soft_deletes = true

[psql]
  dbname = "dbname"
  host   = "localhost"
  port   = 5432
  user   = "dbusername"
  pass   = "dbpassword"
  schema = "myschema"
  blacklist = ["migrations"]
```
* Run sqlboiler to generate models from your database schema
```shell script
$ sqlboiler psql --no-tests --no-context --add-soft-deletes
$ go test ./<YOUR MODELS DIR>
```
