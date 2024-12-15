func main() {

        x := 10

        fmt.Println("x: ", x)

        defer fmt.Println("defer x: ", x)

        x = 5

        fmt.Println("x after change: ", x)



}func main() {
        x := 10
        fmt.Println("x: ", x)
        defer func(x int) {
                fmt.Println("defer x: ", x)
        }(x)
        x = 5
        fmt.Println("x after change: ", x)
}