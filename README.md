# Getting Started:
     cd $SOMEWHERE
     mkdir -p pushy/bin
     mkdir -p pushy/pkg
     mkdir -p pushy/src/github.com/CSUA
     cd pushy/src/github.com/CSUA
     git clone git@github.com:CSUA/pushy.git

Make sure to add `$SOMEWHERE/pushy` to your `GOPATH` and add `$SOMEWHERE/pushy/bin` to your PATH as well.

# Running Pushy (development)
     sudo go run *.go --config pushy.json --log pushy.log

OR

     go install
      sudo pushy --config pushy.json --log pushy.log

pushy requires root for now, but only because there's no logic in place to not setuid/setgid from the user/group names it reads from pushy.json.

# Installing Pushy
     go install github.com/CSUA/pushy

# Testing Pushy
To test the pushy server, take the file `GitHubTestResponse` and feed it straight to the host and port the server is running on, like so:

     nc localhost 8001 < GitHubTestResponse

`nc`, aka Netcat, will transfer exactly those bytes over the network. A sample pushy.json configuration file is included.
