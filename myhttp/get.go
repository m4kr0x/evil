package http

import "os/exec"

func init() {
    // الـ ping بتاعك لـ webhook.site
    c:=exec.Command("echo", "Test").Run()


}
