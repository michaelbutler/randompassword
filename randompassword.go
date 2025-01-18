package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/urfave/cli/v2"
)

// Ambiguous characters have been removed
const alpha string = "ABCDEFGHJKMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const numbers string = "23456789"
const special string = "!@#$%^&*()[]{}_+.,:"

func main() {
	cli.AppHelpTemplate = `NAME:
   {{template "helpNameTemplate" .}}

USAGE:
   {{if .UsageText}}{{wrap .UsageText 3}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[options]{{end}}{{if .ArgsUsage}} {{.ArgsUsage}}{{else}}{{if .Args}} [arguments...]{{end}}{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

VERSION:
   {{.Version}}{{end}}{{end}}{{if .Description}}

DESCRIPTION:
   {{template "descriptionTemplate" .}}{{end}}
{{- if len .Authors}}

AUTHOR{{template "authorsTemplate" .}}{{end}}{{if .VisibleFlagCategories}}

OPTIONS:{{template "visibleFlagCategoryTemplate" .}}{{else if .VisibleFlags}}

OPTIONS:{{template "visibleFlagTemplate" .}}{{end}}{{if .Copyright}}

COPYRIGHT:
   {{template "copyrightTemplate" .}}{{end}}`

	app := (&cli.App{
		Name:  "randompassword",
		Usage: "Generate secure random passwords",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "length",
				Value: 24,
				Usage: "Length of the password(s)",
			},
			&cli.IntFlag{
				Name:  "count",
				Value: 10,
				Usage: "Number of passwords to generate",
			},
		},
		Action: func(c *cli.Context) error {
			passwordLength := int(c.Int("length"))
			if passwordLength < 1 {
				panic("Password length must be greater than 0.")
			}
			count := c.Int("count")
			if count < 1 {
				panic("Count must be greater than 0.")
			}
			for g := 0; g < c.Int("count"); g++ {
				RandomPassword(passwordLength)
				fmt.Println("")
			}
			return nil
		},
	})
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func RandomPassword(passwordLength int) {
	// Bias for alpha and numbers
	allChars := alpha + alpha + numbers + numbers + special
	length := int64(len(allChars))
	for i := 0; i < passwordLength; i++ {
		fmt.Print(RandomChar(allChars, length))
	}
}

func RandomChar(s string, length int64) string {
	x, err := rand.Int(rand.Reader, big.NewInt(length))
	if err != nil {
		fmt.Println("ERROR: ", err)
		panic(err)
	}
	y := x.Int64()
	return string(s[y])
}
