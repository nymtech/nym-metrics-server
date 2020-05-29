# Nym Directory Server

A temporarily centralised PKI, presence and metrics server allowing us to get the other 
Nym node types running. Nym nodes and clients use it to find each other, and
bootstrap the network into existence. Metrics allow us to easily build visualizations
of the network for demonstration, education, and debugging purposes during development.

Eventually some aspects of it (presence, PKI) will be 
decentralized. Other aspects of it (e.g. metrics) will likely stay centralized. 

## Dependencies

* Go 1.12 or later

## Building and running

`go run main.go` builds and runs the directory server


## Usage

Nym clients connect to the directory server and use the public key information and IP addresses it provides to make network requests to the Nym network. This happens automatically, so you shouldn't need to dig into it too much. 

But if you're interested in what it does:

The server exposes an HTTP interface which can be queried. To see documentation 
of the server's capabilities, go to http://localhost:8080/swagger/index.html in
your browser once you've run the server. You'll be presented with an overview
of functionality. All methods are runnable through the Swagger docs interface, 
so you can poke at the server to see what it does. 

## Developing

`go test ./...` will run the test suite.

`swag init` rebuilds the Swagger docs if you've changed anything there. Otherwise
it should not be needed.

If you update any of the HTML assets,
`go-assets-builder server/html/index.html -o server/html/index.go` will
put it in the correct place to be built into the binary. 

