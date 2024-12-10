package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/segwin/adventofcode-2024/internal/solutions"
	"github.com/urfave/cli/v3"
)

var (
	errInvalidArg = errors.New("invalid argument")
)

func main() {
	cmd := cli.Command{
		Name:  "aoc",
		Usage: "Run solutions for the Advent of Code 2024 event. If -day is not set, all solutions are run.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "day",
				Usage: "If set, only run this day's solution. Must be an integer between 1 and 25.",
				Validator: func(v int64) error {
					if v < 1 || v > 25 {
						return fmt.Errorf("%w: -day must be between 1 and 25, got %d", errInvalidArg, v)
					}
					return nil
				},
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			day := cmd.Int("day")
			if day == 0 {
				slog.Info("Running all solutions...")
				if err := solutions.RunAll(); err != nil {
					return fmt.Errorf("one or more solutions failed: %w", err)
				}
			} else {
				slog.Info("Running single day's solution...", "day", day)
				if err := solutions.RunOne(int(day)); err != nil {
					return fmt.Errorf("one or more solutions failed: %w", err)
				}
			}

			slog.Info("Done!")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
