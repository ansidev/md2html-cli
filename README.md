# md2html

[![GoDoc](https://pkg.go.dev/badge/github.com/ansidev/md2html-cli?status.svg)](https://pkg.go.dev/github.com/ansidev/md2html-cli?tab=doc)
[![Release](https://img.shields.io/github/release/ansidev/md2html-cli.svg)](https://github.com/ansidev/md2html-cli/releases)
[![Build Status](https://github.com/ansidev/md2html-cli/workflows/test/badge.svg?branch=main)](https://github.com/ansidev/md2html-cli/actions?query=branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/ansidev/md2html-cli)](https://goreportcard.com/report/github.com/ansidev/md2html-cli)
[![Sourcegraph](https://sourcegraph.com/github.com/ansidev/md2html-cli/-/badge.svg)](https://sourcegraph.com/github.com/ansidev/md2html-cli?badge)

Command line application to convert Markdown file to HTML.

## Usage

```
NAME:
   md2html - Markdown to HTML Converter

USAGE:
   Command line application to convert Markdown file to HTML

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --output value, -o value               Output path. Default output is stdout
   --excerpt-separator value, --es value  Excerpt separator string. By default, this value is an empty string and will be ignored (default: empty string)
   --style value, -s value                CSS to be injected into the output (default: light)
                                          Supported styles:
                                          - dark.
                                          - dark_colorblind.
                                          - dark_dimmed.
                                          - dark_high_contrast.
                                          - dark_tritanopia.
                                          - light.
                                          - light_colorblind.
                                          - light_high_contrast.
                                          - light_tritanopia.
   --template value, -t value             Pass a custom HTML template file path for the output file (default: By default, the built-in template will be used)
   --minify, -m                           Minify the output HTML (default: false)
   --version, -v                          Print out the version of this command (default: false)
   --help, -h                             show help (default: false)
```

## Contact

Le Minh Tri [@ansidev](https://ansidev.xyz/about).

## License

This source code is available under the [AGPL-3.0 LICENSE](/LICENSE).
