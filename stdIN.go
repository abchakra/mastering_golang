package main 
 /**
 Reading from standard input
 **/

 
import ( 
    "bufio" 
    "fmt" 
    "os" 
) 

func main() {
 
    var f *os.File 
    f = os.Stdin 
    defer f.Close() 
 
    scanner := bufio.NewScanner(f) 
    for scanner.Scan() { 
        fmt.Println(">", scanner.Text()) 
    } 
} 

// According to the UNIX way, you can tell a program to stop reading data from standard input by pressing Ctrl + D.