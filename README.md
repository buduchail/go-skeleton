# Go skeleton

A sample microservice written in Go.
This microservice is an example of common infrastructure facilities.

The service can be used to manage a Day of the Dead parade with a REST style interface.
Shrines (`altares`) can be created where gifts (`ofrendas`) can be displayed on shelves (`niveles`) as a tribute to illustrious deceased mexicans (`difuntos`).
To do so, it exposes the following endpoints:

- `http://localhost:8080/api/v1/ofrendas`
- `http://localhost:8080/api/v1/altares`
- `http://localhost:8080/api/v1/altares/*/niveles`
- `http://localhost:8080/api/v1/difuntos`

An `http://localhost:8080/api/v1/status` endpoint is also exposed, which provides some basic introspection into the service.

The following commands can be used to explore what is available:

```shell
$ curl -X GET http://localhost:8080/api/v1/ofrendas/1
$ curl -X GET http://localhost:8080/api/v1/ofrendas?type=flower
$ curl -X GET http://localhost:8080/api/v1/difuntos
$ curl -X GET http://localhost:8080/api/v1/altares/1
```
These commands can also be used to add and remove shrines and display gifts on its shelves:

```shell
$ curl -X POST http://localhost:8080/api/v1/altares -H "Content-type: application/json" -d '{"MexicanID":1,"Levels":3}'
$ curl -X PUT http://localhost:8080/api/v1/altares/1/niveles/1 -H "Content-type: application/json" -d '{"ID":10,"Name":"Cempasúchil","Type":"flower"}'
$ curl -X DELETE http://localhost:8080/api/v1/altares/1
```

A sample `test-use-cases.sh` script is provided. The script tries to launch the service on port 1234, creates a
shrine and puts a gift in the shrine. The output of each call is printed on screen and finally the service is stopped.

For code layout, a standard Go [workspace](https://golang.org/doc/code.html#Workspaces) is used with a Clean Architecture approach to organize packages.


![La Catrina](la-catrina.jpg)

## Build and run

Sample `Dockerfile` and `build.xml` files are provided for development and CI environments.
If using `docker`, Go will be provided by the downloaded container. Otherwise [Go 1.8](https://golang.org/dl) should be installed first.

With `ant`, you can run:

```shell
$ ant build
$ ant test
$ ant run
```

With `docker`, you can run:

```shell
$ docker build -t skel .
$ docker run --detach --publish 8080:8080 --name ms-skel --rm skel
```

To use the standard Go toolchain, run:

```shell
$ source setup.sh
$ go build skel
$ ./skel
```

In all cases, the microservice will be available at `http://localhost:8080`.

On launch, the service accepts the following command line flags:

- `-prefix` < URL prefix for REST resource routes > - default is `/api/v1`
- `-router` <`nethttp`|`iris`|`httprouter`|`echo`|`fasthttp`|`gin`> - default is `nethttp`
- `-port` <`port`> - default is `8080`
- `-corrid` < HTTP header to inspect for Correlation ID value > - default is `X-Correlation-ID`
- `-profile` < path to save profiling info or `web` > - default is empty
- `-loglevel` <`debug`|`info`|`warn`|`error`|`panic`|`fatal`> - default is `info`
- `-logfile` < path to file to write output, stdout if empty > - default is empty
- `-logformat` <`text`|`json`> default is `text`


## Common services

This project implements several common infrastructure services. For each service type, one or more popular libraries are evaluated.

* Configuration
    * [multiconfig](https://github.com/koding/multiconfig) is a lightweight library that allows to load configuration from files, environment variables, command line flags and application defaults. It's used in this project to load default configuration values that can be overwritten by command line flags or environment variables.
    * [viper](https://github.com/spf13/viper) is a powerful and extensible library, but is not evaluated in this project.

* Routing
    * [net/http](https://golang.org/pkg/net/http) is the request handler provided by Go's standard library. This project implements a simple REST wrapper on top of it, as well as on top of several other popular HTTP routers.
    * [HttpRouter](https://github.com/julienschmidt/httprouter)
    * [echo](https://github.com/labstack/echo)
    * [fasthttp](https://github.com/valyala/fasthttp)
    * [iris](https://github.com/kataras/iris)
    * [gin](https://github.com/gin-gonic/gin)

* Logging
    * [logrus](https://github.com/sirupsen/logrus) is a stdlib compatible library that provides context (structured logging), formatters (text, JSON) and output handlers (syslog, logstash). A custom interface and wrapper implementation is provided, which simplifies sending custom context in each log invocation.
     
* Middleware
    * There are several middleware stacks for golang, such as [alice](https://github.com/justinas/alice), [gorilla toolkit](http://www.gorillatoolkit.org/pkg/handlers)'s handlers or the middleware package from the [echo](https://echo.labstack.com/middleware) framework. In this project a simple middleware stack is provided by using functions that implement the http.HandlerFunc interface. The middleware stack is only implemented for the net/http adapter.

* Correlation ID
    * A correlation ID middleware function is implemented and used in the net/http adapter. It checks for a correlation ID header and inserts a new ID if the header does not exist. The ID is then displayed in a request log, handled by another middleware function.

* Input validation
    * This project implements input validation at controller and use case level. Custom libraries can be used to extract validation logic to a validation package.
    * [validator](https://github.com/go-validator/validator)
    * [govalidator](https://github.com/asaskevich/govalidator)

* Output formatting
    * Native JSON encoding is convenient but slow.
    * [easyjson](https://github.com/mailru/easyjson) provides easy and fast marshalling/unmarshaling between Go structs and JSON by generating native code for each custom data type.
    * [Protocol Buffers/gRPC](http://google.golang.org/grpc) - [gRPC](http://www.grpc.io) is a framework that can be used to build efficient communication stacks based on binary serialization.


## Development tasks

### Dependency management

Package dependency management (vendoring) is notably [basic](https://blog.gopheracademy.com/advent-2016/saga-go-dependency-management/) in the Go ecosystem.
There are several popular [tools](https://github.com/golang/go/wiki/PackageManagementTools), such as
[glide](https://github.com/Masterminds/glide) and [gb](https://getgb.io/), that help with organizing project dependencies.
Some use the official [GO15VENDOREXPERIMENT](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit) experiment while others use a custom build process.
This project uses [gpm](https://github.com/pote/gpm), a simple tool inspired in tools such as npm and composer. (TBD)

### Tests

[Testing Microservices in Go](https://blog.gopheracademy.com/advent-2014/testing-microservices-in-go) is a good introduction
to testing in Go and using component tests when testing microservices.
In this project, domain classes are covered with unit tests using the standard golang testing framework, while functional/component
tests are provided to test routes at component level with [httpexpect](https://github.com/gavv/httpexpect).
[Table driven tests](https://github.com/golang/go/wiki/TableDrivenTests) are used in unit testing.

### Profiling

This project uses [pprof](https://golang.org/pkg/net/http/pprof), the standard golang profiling tool, to enable profiling of the service.

To enable profiling, start the service with the `-profile` flag:

```shell
./skel -profile web
```

This will expose a web interface to the integrated profiling tools on a port one number higher than that of the service itself.
If the service is started on port 8080, the profiling interface will be on port 8081.

To profile the service, run a benchmark tool such as Apache Bench against the endpoint you want to profile.
While the benchmark tool is running, run the following commands to get a CPU profile (replace `profile` with `heap` to get a memory profile):

```shell
$ go tool pprof http://localhost:8081/debug/pprof/profile
Entering interactive mode (type "help" for commands)
(pprof) web
```

This will use the default browser to open an image with a visual call stack and time spent in each function.
To get a list of functions taking the most CPU, run `top` instead of `web`.

A sample `benchmark.sh` script is provided to launch Apache Bench benchmarks. It assumes `ab` is installed system wide, the service is
built as `skel` in the current directory and port `8080` is free.

## Best practices

As it has been [said](https://talks.golang.org/2013/bestpractices.slide#2),
Go code should be simple, readable and maintainable. "[Go best practices, six years in](https://peter.bourgon.org/go-best-practices-2016)"
gives a more detailed view on common topics such as code formatting, design, logging, testing and dependency management.
[Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) is a comprehensive "laundry list" of bits of good advice and [Error handling and Go](https://blog.golang.org/error-handling-and-go) is a great article reasoning about the idiomatic way of handling errors in Go.

Regarding code style, the [gofmt](https://golang.org/cmd/gofmt/) command, part of the standard Go toolchain, can be used to apply standard formatting to Go code. Many editors and IDEs are set up to use it on save. 

[Go vet](https://golang.org/cmd/vet) is another tool from the standard toolchain and can be used to do static analysis and detection of suspicious and potentially buggy constructs.


## Other resources

Language:

* [Official documentation](https://golang.org/doc)
* [Official wiki](https://github.com/golang/go/wiki) with lots of useful pointers
* [Writing Web Applications](https://golang.org/doc/articles/wiki) - A simple tutorial covering the very basics
* [Learn X (Go) in Y minutes](https://learnxinyminutes.com/docs/go) - A no frills, learn by example tutorial that also works as a quick Go language reference
* [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/en/) - A nice and extensive open source book about writing web applications in Go

Libraries:

* https://golang.org/pkg/#other
* https://github.com/golang/go/wiki/Projects
* http://go-search.org/tops
* https://github.com/avelino/awesome-go