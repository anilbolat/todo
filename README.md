# Go Course


## Table of content

- [x] Dependency management
- [x] HTTP server
    - [x] http test utils
    - [x] json en/decoding
    - [x] graceful shutdown
- [x] Makefile
    - [x] Run tests
    - [x] Run static checks
    - [x] Build binary
    - [x] Build Docker image
    - [x] go mod in Makefile
- [x] Error handling
    - [x] fmt.Errorf()
    - [x] errors.New()
- [x] Coding patterns/conventions
    - [x] Struct methods
    - [x] Embedded structs
    - [x] Early exit strategy
    - [x] Interfaces
- [x] Folder structure
    - [x] Single command (main in root)
    - [x] Multi command (cmd/{name}/main.go})
- [x] Testing
    - [x] Basic test function
    - [x] Assert and require
    - [x] Table driven testing
    - [x] HTTP testing
    - [x] Benchmarking
- [x] Profiling
    - [x] CPU
    - [x] Memory
    - [x] Tracing
    - [x] HTTP load testing
- [x] Docker
    - [x] security
    - [x] multistage build
    - [x] image size
    - [x] boot up time

## Extra material:
- [ ] Goroutines, channels, WaitGroups, etc
- [ ] Anonymous functions and closures
- [ ] Mocks in testing
- [ ] Debugger
- [ ] Integration tests

## Flow

* Initial commit
    * create repo
    * clone repo
* Add README and Makefile
    * add test target
    * go mod init
* Initial TODO list implementation
    * test case for todo list
    * add lint target
    * add vendor target
    * make vendor
    * List struct
    * Task struct
    * Create pointer receivers for List
    * Split Task to Task and TaskData
    * Static error
* Add benchmark
    * add benchmark target
* Add TaskData validation and table-driven tests
    * table-driven testing
    * convert to subtests
    * add utf-8 string in test
    * fix string validation to support utf-8
* Use parallel subtests
    * wrap with t.Run()
    * copy test variable
* Validate string length in runes
    * fix validate function implementation
* Take TaskData validation in use and add benchcmp
    * use validate func in Add method
    * add benchcmp target
    * run benchcmp
* Add test case for rest API
    * create test case with Recorder
    * make vendor
    * create test case with TestServer
* Implement REST API for todo list
    * create Handler struct
    * wrapper List implementation
    * create endpoint handlers
* Make api test parallel
    * race detection
* Make todo list thread safe
    * atomic uint
    * sync map
    * benchcomp
* Add cmd to run server
    * add build target
    * init todo list and router
    * start server
    * listen SIGINT and SIGTERM
    * graceful shutdown
    * demonstrate with sleep in handler func
* Create api and mem subpackages
    * move List to mem
    * move Handler to api
    * create interface for List
* Create postgres implementation of List
    * create postgres implementation of List
    * test by taking into use in main.go
* Add Dockerfile
    * add docker-image target
    * add docker-run target
    * create multistage dockerfile
* Add targets for docker image size and boot up time
    * compare with reference java implementation
* Add http load testing targets
    * add docker-stats target
    * add wrk target
    * add vegeta target
    * java vs go
        * memory usage
        * avg req/s
        * java warmup time
        * perf while warming up
* Add profiling targets and pprof server
    * add profiling server
    * add cpu-profile target
    * add heap-profile target
    * add test-profile target
    * expose 6060 port for docker
    * show cpu profile
    * show heap profile
    * change CreatedAt to UTC()
    * compare mem usage
* Add execution tracer target
    * add trace-profile target
    * profile while wrk-post is running
* Add clean target
    * clean prof, html, bin, build, etc files

### Example queries

Create task:
```sh
printf '{"Name": "my task", "CreatedAt": "%s"}' $(date -Is) | curl -H 'Content-Type: application/json' -d @- 127.0.0.1:8000/api/tasks
```

Get task:
```sh
curl localhost:8000/api/tasks/0
```
