package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/urfave/cli/v2"
	"heckel.io/ntfy/v2/cmd"
)

// These variables are set during build time using -ldflags
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	cli.AppHelpTemplate += fmt.Sprintf(`
Try 'ntfy COMMAND --help' or https://ntfy.sh/docs/ for more information.

To report a bug, open an issue on GitHub: https://github.com/binwiederhier/ntfy/issues.
If you want to chat, simply join the Discord server (https://discord.gg/cT7ECsZj9w), or
the Matrix room (https://matrix.to/#/#ntfy:matrix.org).

ntfy %s (%s), runtime %s, built at %s
Copyright (C) Philipp C. Heckel, licensed under Apache License 2.0 & GPLv2
`, version, maybeShortCommit(commit), runtime.Version(), date)

	app := cmd.New()
	app.Version = version
	app.Metadata = map[string]any{
		cmd.MetadataKeyDate:   date,
		cmd.MetadataKeyCommit: commit,
	}

	args := os.Args
	if len(args) == 1 {
		args = append(args, "serve", "--listen-http", ":8080")
		go func() {
			time.Sleep(500 * time.Millisecond)
			exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:8080/app.html").Start()
		}()
	}

	if err := app.Run(args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func maybeShortCommit(commit string) string {
	if len(commit) > 7 {
		return commit[:7]
	}
	return commit
}
