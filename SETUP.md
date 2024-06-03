### Requirements

- Go 1.2.2 or later
- Development library tools `https://github.com/mattn/go-sqlite3?tab=readme-ov-file#linux`
- Docker (if using docker-compose to build)

You can build the application one of two ways

- makefile: run `make` and it will default to building the application
- docker: run `docker composer up` to build and run the application.

For debugging locally, you can use the following command: `make go-debug`