package banner

import "fmt"

const Logo = `
╔══════════════════╗
║  ▄▄▄▄▄▄▄▄▄▄▄▄▄▄  ║
║  █ ░▒▓█▓▒░ █    ║
║  █ ░▒▓▀▓▒░ █    ║
║  ▀▀▀▀▀▀▀▀▀▀▀▀▀▀  ║
╚══════════════════╝
    ░░░MIRRORGATE░░░
    ░░REVERSE PROXY░░
`

const Version = "v0.0.1"

func PrintBanner() {
    fmt.Printf("\033[36m%s\033[0m\n", Logo)
    fmt.Printf("\033[33mVersion: %s\033[0m\n\n", Version)
}
