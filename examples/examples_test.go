package examples_test

import(
  "fmt"
  "math/rand"
)

func Perm(x int) []int{
  return rand.Perm(x)
}


func ExampleHello() {
    fmt.Println("hello")
    // Output: hello
}

func ExampleSalutations() {
    fmt.Println("hello, and")
    fmt.Println("goodbye")
    // Output:
    // hello, and
    // goodbye
}

func ExamplePerm() {
    for _, value := range Perm(4) {
        fmt.Println(value)
    }
    // Unordered output:
    // 2
    // 1
    // 3
    // 0
}
