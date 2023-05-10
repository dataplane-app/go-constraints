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

Each function makes use of the filename suffix e.g. _linux.go instead of using // go:build linux commentary

More examples:
https://www.jetbrains.com/help/go/configuring-build-constraints-and-vendoring.html#build-constraints