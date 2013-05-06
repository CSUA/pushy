# Getting Started:
     cd $SOMEWHERE
     mkdir -p pushy/bin
     mkdir -p pushy/pkg
     mkdir -p pushy/src/github.com/CSUA
     cd pushy/src/github.com/CSUA
     git clone git@github.com:CSUA/pushy.git

Make sure to add `$SOMEWHERE/pushy` to your `GOPATH`. If you want to use the binary produced after `go install`, then add `$SOMEWHERE/pushy/bin` to your PATH as well.

# Running Pushy (development)
     go run pushy.go  --config pushy.json --log pushy.log

# Installing Pushy
     go install github.com/CSUA/pushy

# Testing Pushy
To test the pushy server, take the file `GitHubTestResponse` and feed it straight to the host and port the server is running on, like so:

     nc localhost 8001 < GitHubTestResponse

`nc`, aka Netcat, will transfer exactly those bytes over the network. A sample pushy.json configuration file is included.