package main
import (
    "fmt"
    "github.com/smudugal/gocrackcode/ch1/1_8/rotated"
)

func main() {
    fmt.Println("Enter the original string: ")
    var orig_str string
    fmt.Scanf("%s", &orig_str)

    var rotated_str string
    fmt.Println("Enter the rotated string: ")
    fmt.Scanf("%s", &rotated_str)

    if(rotated.IsRotated(orig_str, rotated_str)){
        fmt.Println("The string", rotated_str, "is a rotated version of", orig_str)
    }else{
        fmt.Println("The string", rotated_str, "is NOT a rotated version of", orig_str)
    }
}
