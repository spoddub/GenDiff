package main

import (
	"code"
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func main() {
	cmd := &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows a difference.",
		UsageText: "gendiff [global options]",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "--format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},

		Action: func(ctx context.Context, command *cli.Command) error {
			if command.Args().Len() != 2 {
				return fmt.Errorf("need two files")
			}

			file1 := command.Args().Get(0)
			file2 := command.Args().Get(1)
			format := command.String("format")

			result, err := code.GenDiff(file1, file2, format)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
