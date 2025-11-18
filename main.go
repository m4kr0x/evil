package poc

/*
#cgo CFLAGS: -Werror
void x() {}
*/
import "C"

import "fmt"

func init() {
    fmt.Println(">>> SAFE POC RAN DURING BUILD <<<")
}
