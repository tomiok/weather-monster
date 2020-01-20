# weather monster API
---------------------

### Simple REST API in Golang

#### uses
* Go 1.13
* Go mods
* go-chi
* database sql and go sql driver
* mysql
* db migrations library

#### build & run (macOS and *nix based OS)
- mysql in localhost, create a DB <weather_monster> or any other in
the env variable (DB_NAME), and set username and password
in env variables (DB_USER and DB_PASS). Otherwise username root and password root will be set.

- build with make, with the command `make build`

- run the tests, with the command `make test`

- (optional) format the code with the command `make fmt`

- (optional) a linter is available using those commands `make lint-prepare` and `make lint`

- run the project with the command `make web` and by default uses port 9000

------------------------------------------------------
- the dependencies management is from go mods, use `go mod tidy` to set up the dependencies
and `go mod vendor` to use the vendor folder.

### out of scope
* the http request are not validated at all

* not mocking tests are done since all the logic resides in db queries. Some tests in http APIs only

* not documented to run it on Windows

* error handling is just the basics

* no dockerfile since is not deployed yet

* no godocs for the PoC