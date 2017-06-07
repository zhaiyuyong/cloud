package main

import (
	"context"
	"flag"
	"os"

	"github.com/PaddlePaddle/cloud/go/paddlecloud"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&paddlecloud.SubmitCmd{}, "")
	subcommands.Register(&paddlecloud.LogsCommand{}, "")
	subcommands.Register(&paddlecloud.GetCommand{}, "")
	subcommands.Register(&paddlecloud.KillCommand{}, "")
	subcommands.Register(&paddlecloud.SimpleFileCmd{}, "")
	subcommands.Register(&paddlecloud.LsCommand{}, "")
	subcommands.Register(&paddlecloud.CpCommand{}, "")
	subcommands.Register(&paddlecloud.RmCommand{}, "")
	subcommands.Register(&paddlecloud.MkdirCommand{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}