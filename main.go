package poc

/*
#cgo CFLAGS: -Werror
#cgo LDFLAGS: 
void test() {}
*/
import "C"

import "fmt"

func init() {
    fmt.Println(">>> POC BUILD EXECUTED <<<")
}
