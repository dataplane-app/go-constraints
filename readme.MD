### An example of cross compiling Golang with import functions specific to different OS

Each function is for a different operating system. 

## Windows
```
GOOS=windows go run main.go
```

## Mac (Darwin)
```
GOOS=darwin go run main.go
```

## Linux
```
GOOS=linux go run main.go
```