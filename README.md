# Nym Metrics Server

A central metrics server which keeps track of the current mixing state of the network. 

## Dependencies

* Go 1.12 or later

## Building and running

`go run main.go` builds and runs the metrics server


## Usage

Nym nodes periodically send metrics information (how many Sphinx packets they've sent and received in a given interval). These metrics allow us to easily build visualizations of the network for demonstration, education, and debugging purposes during development and testnet.

To see documentation of the server's capabilities, go to http://localhost:8080/swagger/index.html in your browser. All methods are runnable through the Swagger docs interface, so you can poke at the server to see what it does. 

## Developing

`go test ./...` will run the test suite.

`swag init` rebuilds the Swagger docs if you've changed anything there. Otherwise it should not be needed.

If you update any of the HTML assets, `go-assets-builder server/html/index.html -o server/html/index.go` will put it in the correct place to be built into the binary. 

