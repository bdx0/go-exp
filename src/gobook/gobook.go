package dobook
// https://www.golang-book.com/books/intro/13#section9
import (
	"fmt"
	"github.com/ZhangBanger/golang-book/chapter11/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
