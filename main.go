package main

import (
	"log"
	"os"

	"github.com/ansidev/md2html-cli/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      cmd.CommandName,
		Usage:     "Markdown to HTML Converter",
		UsageText: "Command line application to convert Markdown file to HTML",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Value:   "",
				Usage:   "Output path. Default output is stdout",
			},
			&cli.StringFlag{
				Name:        "excerpt-separator",
				Aliases:     []string{"es"},
				Value:       "",
				DefaultText: "empty string",
				Usage:       "Excerpt separator string. By default, this value is an empty string and will be ignored",
			},
			&cli.StringFlag{
				Name:        "style",
				Aliases:     []string{"s"},
				Value:       "light",
				DefaultText: "light. Supported styles: dark_colorblind, dark_dimmed, dark_high_contrast, dark_tritanopia, dark, light_colorblind, light_high_contrast, light_tritanopia, light",
				Usage:       "CSS to be injected into the output",
			},
			&cli.StringFlag{
				Name:        "template",
				Aliases:     []string{"t"},
				Value:       "",
				DefaultText: "By default, the built-in template will be used",
				Usage:       "Pass a custom HTML template file path for the output file",
			},
			&cli.BoolFlag{
				Name:    "minify",
				Aliases: []string{"m"},
				Value:   false,
				Usage:   "Minify the output HTML",
			},
			&cli.BoolFlag{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Print out the version of this command",
			},
		},
		Action: cmd.Md2HtmlCommandHandler,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
