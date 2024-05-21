// Copyright (c) 2024 John Dewey

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package cmd

import (
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var (
	logger *slog.Logger
	//go:embed resources/art.txt
	asciiArt string
)

const (
	desc    = "A golang based FINGERD(8) server.\n"
	website = "https://github.com/retr0h/fingerme"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fingerme",
	Short: desc,
	Long:  fmt.Sprintf("%sfingerme: %s\n%s", asciiArt, desc, website),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable or disable debug mode")

	_ = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	cobra.OnInitialize(initLogger)
}

func initLogger() {
	logLevel := slog.LevelInfo
	if viper.GetBool("debug") {
		logLevel = slog.LevelDebug
	}

	logger = slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      logLevel,
			TimeFormat: time.Kitchen,
			NoColor:    !term.IsTerminal(int(os.Stdout.Fd())),
		}),
	)
}
