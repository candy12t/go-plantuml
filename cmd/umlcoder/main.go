package main

import (
	"fmt"
	"net/url"
	"os"
	"runtime/debug"

	"github.com/candy12t/go-plantuml"
	"github.com/urfave/cli/v2"
)

const (
	name = "umlcoder"
)

var version string

func Version() string {
	if version != "" {
		return version
	}

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		return buildInfo.Main.Version
	}
	return "dev"
}

func main() {
	app := &cli.App{
		Name:      name,
		Usage:     "Command Line Tool to encode/decode PlantUML",
		Version:   Version(),
		Reader:    os.Stdin,
		Writer:    os.Stdout,
		ErrWriter: os.Stderr,
		CommandNotFound: func(ctx *cli.Context, command string) {
			fmt.Fprintf(ctx.App.Writer, "unknown command %q. see `%s --help` for more details\n", command, name)
		},
		Commands: []*cli.Command{
			{
				Name:      "encode",
				Aliases:   []string{"enc"},
				Usage:     "Encode PlantUML source code into strings",
				Args:      true,
				ArgsUsage: "[PlantUML FILES...]",
				Action:    encode,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "format",
						Usage:   "specify which one: png, svg, txt",
						Value:   "png",
						Aliases: []string{"f"},
					},
					&cli.StringFlag{
						Name:    "url",
						Usage:   "specify plantuml server url",
						Value:   "https://www.plantuml.com/plantuml",
						Aliases: []string{"u"},
					},
				},
			},
			{
				Name:      "decode",
				Aliases:   []string{"dec"},
				Usage:     "Decode encoded string into PlantUML source codes",
				Args:      true,
				ArgsUsage: "[Encoded STRINGS...]",
				Action:    decode,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(app.ErrWriter, err)
	}
}

type format string

const (
	png format = "png"
	svg format = "svg"
	txt format = "txt"
)

func (f format) String() string {
	return string(f)
}

func parseFormat(s string) (format, error) {
	switch s {
	case png.String():
		return png, nil
	case svg.String():
		return svg, nil
	case txt.String():
		return txt, nil
	default:
		return "", fmt.Errorf("unknown format: %s", s)
	}
}

func encode(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return cli.ShowSubcommandHelp(ctx)
	}

	baseURL := ctx.String("url")
	format, err := parseFormat(ctx.String("format"))
	if err != nil {
		return cli.ShowSubcommandHelp(ctx)
	}

	for _, file := range ctx.Args().Slice() {
		text, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		result, err := plantuml.Encode(text)
		if err != nil {
			return err
		}
		u, err := url.JoinPath(baseURL, format.String(), result)
		if err != nil {
			return err
		}
		fmt.Fprintln(ctx.App.Writer, u)
	}
	return nil
}

func decode(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return cli.ShowSubcommandHelp(ctx)
	}

	for _, str := range ctx.Args().Slice() {
		result, err := plantuml.Decode(str)
		if err != nil {
			return err
		}
		fmt.Fprintln(ctx.App.Writer, string(result))
	}
	return nil
}
