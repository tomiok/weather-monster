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

============================

	golang course
	

	* Que problemas nos resuelve GO? 


	- Standard library muy efectiva, amplia, escalable (http, zip, file system, OS, acceso a hardware, syscalls, etc).
	- Sysntaxis reducida, lo que lo hace facil para aprender (25 palabras reservadas, 53 tiene Java 1.11).
	- No es necesario descargar nada para su uso web, o de file system por ejemplo, todo está contenido en su runtime.
	- Muchas herramientas incluidas en su runtime, fmt (formato para identacion del codigo), build, run, test entre otras.
	- Facil descarga de dependencias con go get
	- Go playground
	- Facil manejo de dependencias con go mod (go modules)
	- Aplicaciones web (APIS, webs SSR server side rendering), backend, server-side rendering con template engine integrado, machine learning, big data, juegos (se vienen videos :D)
	- Tipos primitivos de concurrencia
	- Comunidad activa
	- Releases continuos (14 en 10.3 años)


	https://golang.org/cmd/go/
	
	https://tour.golang.org/