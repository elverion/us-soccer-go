### Requirements

- Go 1.2.2 or later
- Development library tools `https://github.com/mattn/go-sqlite3?tab=readme-ov-file#linux`
- Docker (if using docker-compose to build)
- gcc

You can build the application one of two ways

- makefile: run `make` and it will default to building the application
- docker: run `docker composer up` to build and run the application.

For debugging locally, you can use the following command: `make go-debug`

For testing, you can run `go test ./...` from the command line or if you have dev tools installed, you can run `make go-test` 
If you would like coverage, make sure `https://github.com/jandelgado/gcov2lcov` is installed and then run `make go-test-coverage`