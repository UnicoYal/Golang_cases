# Go: Обход строки

Так как строка — это массив байт, ее можно обойти с помощью цикла for:
```go
package main

import (
    "fmt"
)

func main() {
    s := "hello"
    for i := 0; i < len(s); i++ {
        fmt.Println(string(s[i]))
    }

}
```
Вывод:
```go
h
e
l
l
o
```

Таким способом можно обойти только строки, состоящие из ASCII символов. Если строка содержит мультибайтовые символы, вывод будет некорректен:
```go
package main

import (
    "fmt"
)

func main() {
    s := "привет"
    for i := 0; i < len(s); i++ {
        fmt.Println(string(s[i]))
    }

}
```
# Task


Реализуйте функцию shiftASCII(s string, step int) string, которая принимает на вход состоящую из ASCII символов строку s и возвращает новую строку, где каждый символ из входящей строки сдвинут вперед на число step. Например:
```go
shiftASCII("abc", 0) // "abc"
shiftASCII("abc1", 1) // "bcd2"
shiftASCII("bcd2", -1) // "abc1"
shiftASCII("hi", 10) // "rs"
```
Если после сдвига код символа выходит за рамки ASCII, то число должно быть взято по модулю 256:
```go
shiftASCII("abc", 256) // "abc"
shiftASCII("abc", -511) // "bcd"
```