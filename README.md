# directory-server

A temporarily centralised PKI, presence and metrics server allowing us to get the other 
Nym node types running.

## Dependencies

* Go 1.12 or later

## Building

Nothing too special here, `go run main.go` should see you through. 

`go test ./...` will run the test suite.

## Usage

The server exposes an HTTP interface which can be queried. To see documentation 
of the server's capabilities, go to http://localhost:8080/swagger/index.html in
your browser once you've run the server. You'll be presented with an overview
of functionality. All methods are runnable through the Swagger docs interface, 
so you can poke at the server to see what it does. 

