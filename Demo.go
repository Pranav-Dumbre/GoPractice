package main

import "fmt"

// go env -> check env variables
// go env -w GO111MODULE=on  -> enable module
// go env GOPATH  -> check go path is set
// go mod init myproject -> for new project creates go.mod file that handle dependancies
// go mod tidy -> clean up dependancies 

func main(){
	fmt.Println("My first Go code")
}