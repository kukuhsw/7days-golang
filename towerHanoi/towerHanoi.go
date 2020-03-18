package main
  
import "fmt"
 
type solver interface {
    play(int)
}
   
// Tower adalah contoh antarmuka pemecah tipe yang sangat baik
type towers struct {
    // struck kosong
}
  
// bermain adalah satu-satunya metode yang diperlukan untuk mengimplementasikan tipe solver
func (t *towers) play(n int) {    
    t.moveN(n, 1, 2, 3)
}
  
// algoritma rekursif
func (t *towers) moveN(n, from, to, via int) {
    if n > 0 {
        t.moveN(n-1, from, via, to)
        t.moveM(from, to)
        t.moveN(n-1, via, to, from)
    }
}
 
func (t *towers) moveM(from, to int) {
    fmt.Println("Pindahkan lingkaran dari tiang", from, "ke tiang", to)
}
 
func main() {
    var t solver    
    t = new(towers) // jenis menara harus memenuhi antarmuka pemecah
    t.play(4)
}