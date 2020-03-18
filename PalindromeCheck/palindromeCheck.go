package main
import "fmt"
 
func main() {
    var number,remainder,temp int
    var reverse int = 0
 
    fmt.Print("Masukkan angka : ")
    fmt.Scan(&number)
 
    temp=number
 
    // Untuk Loop yang digunakan dalam format While Loop
    for{
        remainder = number%10
        reverse = reverse*10 + remainder
        number /= 10
 
        if(number==0){
            break // Break Statement yang digunakan untuk keluar dari loop
        }
    }
 
    if(temp==reverse){
        fmt.Printf("%d ini Palindrome",temp)
    }else{
        fmt.Printf("%d bukan Palindrome",temp)
    }
}