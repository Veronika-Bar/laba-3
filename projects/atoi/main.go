package main
import "fmt"

func main() {
    var x int
    fmt.Scan(&x)

    for i:=0; x > 0; i++ {
        sum := x%10
        defer fmt.Print(sum*sum) //Проходим число с конца, вывод квадрата откладываем в стек вызовов
        x = x/10
    }
    
}