package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ansidev/md2html"
	"github.com/ansidev/md2html-cli/utils"
	"github.com/urfave/cli/v2"
)

const (
	// CommandName The command name
	CommandName = "md2html"
	// version of this command
	version = "1.0.0"
	// defaultOutputDir default output dir for generated HTML files
	defaultOutputDir = "./"
)

// Md2HtmlCommandHandler Handler for the command md2html
func Md2HtmlCommandHandler(ctx *cli.Context) error {
	cmdVersion := ctx.Bool("version")

	if cmdVersion {
		fmt.Printf("%s v%s\n", CommandName, version)
		return nil
	}

	args := ctx.Args().Slice()

	if len(args) == 0 {
		fmt.Print("Invalid input arguments. Showing app usage.\n\n")
		return cli.ShowAppHelp(ctx)
	}

	input, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}

	options := md2html.Options{
		// HTMLTemplate:     "",
		ExcerptSeparator: ctx.String("excerpt-separator"),
		Style:            ctx.String("style"),
		Minify:           ctx.Bool("minify"),
	}

	customHTMLTemplate := ctx.String("template")
	if len(customHTMLTemplate) > 0 {
		htmlTemplate, err := os.ReadFile(customHTMLTemplate)
		if err != nil {
			return err
		}
		options.HTMLTemplate = string(htmlTemplate)
	}

	html, err := md2html.Convert(input, options)
	if err != nil {
		return err
	}

	outputPath := ctx.String("output")
	if len(outputPath) == 0 {
		fmt.Print(html)
		return nil
	}

	outputDir, outputFile := filepath.Split(outputPath)

	if len(outputDir) == 0 {
		outputDir = defaultOutputDir
	}

	if len(outputFile) == 0 {
		_, inputFileName := filepath.Split(outputPath)
		fileNameWithoutExtension := utils.FileNameWithoutExtension(inputFileName)
		if len(fileNameWithoutExtension) == 0 {
			outputFile = "out.html"
		} else {
			outputFile = fmt.Sprintf("%s.html", fileNameWithoutExtension)
		}
	}

	utils.CreateDirIfNotExists(outputDir)

	outputPath = filepath.Join(outputDir, outputFile)
	f, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(html)

	return err
}
