package cli

import (
	"fmt"
	"github.com/stelligent/mu/common"
	"github.com/urfave/cli"
	"github.com/stelligent/mu/workflows"
)

func newPipelinesCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:  "pipeline",
		Usage: "options for managing pipelines",
		Subcommands: []cli.Command{
			*newPipelinesListCommand(ctx),
			*newPipelinesShowCommand(ctx),
			*newPipelinesUpsertCommand(ctx),
			*newPipelinesTerminateCommand(ctx),
		},
	}

	return cmd
}

func newPipelinesListCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "list pipelines",
		Action: func(c *cli.Context) error {
			fmt.Println("listing pipelines")
			return nil
		},
	}

	return cmd
}

func newPipelinesShowCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:  "show",
		Usage: "show pipeline details",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "service, s",
				Usage: "service to show",
			},
		},
		Action: func(c *cli.Context) error {
			service := c.String("service")
			fmt.Printf("showing pipeline: %s\n", service)
			return nil
		},
	}

	return cmd
}

func newPipelinesTerminateCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:  "terminate",
		Aliases:   []string{"term"},
		Usage: "terminate pipeline",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "service, s",
				Usage: "service to terminate pipeline",
			},
		},
		Action: func(c *cli.Context) error {
			service := c.String("service")
			fmt.Printf("terminating pipeline: %s\n", service)
			return nil
		},
	}

	return cmd
}

func newPipelinesUpsertCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:  "upsert",
		Aliases:   []string{"up"},
		Usage: "upsert pipeline",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "token, t",
				Usage: "GitHub token ",
			},
		},
		Action: func(c *cli.Context) error {
			token := c.String("token")
			workflow := workflows.NewPipelineUpserter(ctx, func () string {
				// TODO: prompt for STDIN if not provided via a flag
				return token
			})
			return workflow()
		},
	}

	return cmd
}
