package kata2

import "fmt"

func rot13(i byte) byte {
    switch {
    case 65 <= i && i <= 77, 97 <= i && i <= 109:
        return i + 13
    case 78 <= i && i <= 100, 110 <= i && i <= 122:
        return i - 13
    }
    return i
}

func MakeKata(input string) string {
    length := len(input)
    counter := 0
    out := ""
    for {
        if counter == length {
            break
        }
        out += fmt.Sprintf("%c", rot13(input[counter]))
        counter += 1
    }
    return out
}

// import (
//     "io"
//     //"os"
//     "fmt"
//     "strings"
// )

// type rot13Reader struct {
//     r io.Reader
// }

// func (self *rot13Reader) Read(input []byte) (count int, e error) {
//     var output []byte = []byte{}
//     container := make([]byte, 1, 1)
//     for {
//         c, err := self.r.Read(container)
//         count += c
//         fmt.Println("container is ", container)
//         if err != nil {
//             if err.Error() != "EOF" {
//                 e = err
//             }
//             return
//         }
//         output = append(output, container[0])
//         copy(input, output)
//         container[0] = 0
//     }

// }

// func main() {
//     s := strings.NewReader("Lbh penpxrq gur pbqr!")
//     r := rot13Reader{s}

//     b := make([]byte, 100, 100)
//     c, e := r.Read(b)
//     fmt.Println(c, e, b)
//     //io.Copy(os.Stdout, &r)
// }
