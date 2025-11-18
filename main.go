// Source - https://stackoverflow.com/a
// Posted by Lourenco, modified by community. See post 'Timeline' for change history
// Retrieved 2025-11-18, License - CC BY-SA 4.0

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("curl", "https://webhook.site/eb4f653b-4fb0-4754-bbde-7086d88f9b8a")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
