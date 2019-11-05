package main 
 

/**
Working with command-line arguments
**/

import ( 
    "fmt" 
    "os" 
    "strconv" 
)

func main() { 
    if len(os.Args) == 1 { 
        fmt.Println("Please give one or more floats.") 
        os.Exit(1) 
    } 
 
    arguments := os.Args 
    min, _ := strconv.ParseFloat(arguments[1], 64) 
	max, _ := strconv.ParseFloat(arguments[1], 64) 
	for i := 2; i < len(arguments); i++ { 
		n, _ := strconv.ParseFloat(arguments[i], 64) 
		// The underscore character, which is called blank identifier, is the Go way of discarding a value. If a Go function returns multiple values, you can use the blank identifier multiple times.
 
        if n < min { 
            min = n 
        } 
        if n > max { 
            max = n 
        } 
    } 
 
    fmt.Println("Min:", min) 
    fmt.Println("Max:", max) 
} 