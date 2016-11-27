# Improbable Echo

A Go HTTP server, built using [Echo](https://github.com/labstack/echo).

# Setting Up

1. Ensure that you have Go installed and that your $GOPATH is set
2. `git clone` the repo to your machine
3. Install Echo by running `go get -u github.com/labstack/echo`

# Running

1. Run `go run main.go`. This will start the server on port 1337
2. While the server is running, run the command `curl --data-ascii {'"secret":"shotsfired"'} -X POST http://localhost:1337/hello/jonny --header "Content-Type:application/json"`
3. Enjoy the HTML response and custom response header!
