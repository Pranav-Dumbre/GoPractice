package main

import "fmt" // imports the fmt (formatting) package from the standard library.

// go env -> check env variables
// go env -w GO111MODULE=on  -> enable module
// go env GOPATH  -> check go path is set
// go mod init myproject -> for new project creates go.mod file that handle dependancies
// go mod tidy -> clean up dependancies

func main() {
	fmt.Println("My first Go code")

	// initialize variables here
	var smsSendingLimit int = 0
	var costPerSMS float64 = 0.1
	var hasPermission bool = true
	var username string = "Pranav"
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// walrus operator (:=)
	// it has type inference -> to the automatic detection of the type of an expression in a formal language.
	// cant be used outside of function
	myMessegeCount := 35
	fmt.Printf("%v", myMessegeCount)

	// typecasting int to float64
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)

	// This is a single line comment

	/*
	   This is a multi-line comment

	   working of go code :
	   code -> compile (go build) -> executable file .exe
	*/

	/*
		Type sizes :
		integer                 -> int  int8  int16  int32  int64
		unsigned integer        -> uint uint8 uint16 uint32 uint64 uintptr
		signed decimal numbers  -> float32 float64
		complex numbers         -> complex64 complex128

		standard types are      -> int uint float64 complex128
	*/

	accountAgeFloat := 2.6
	accountAgeInt := int(accountAgeFloat)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
