# HOW 

```
go get github.com/user/hello
hello
```

Responses with

```
Bonjour
```

# Run as single file
```
go run main.go
```

# Other
```
go build
go install
go test -v github.com/syzer/go-hello/util
godoc ./util/
```

# chat
```
go build -o webServer/chat
cd webServer
./chat -addr="192.168.0.1:3000"

# or
./chat -addr=":3001"

# or
fresh go run main.go    
```
