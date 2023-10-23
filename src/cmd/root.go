// metabypass
// src/cmd/root.go

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"metabypass/generator"
	"os"
)

// Yeah, I know.. using COBRA is a bit of overkill for this software, right now (as of v1.00.00), but I leave room to expand it a bit...

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Redirects a selected URL on a web page hosted on a web server you control",
	Long: `Suppose you want to publish https://my_just_found_website/its_page.html to a website that blocks such contents.
All you need is to copy that URL in the clipboard, and run this software with a new web page that will be created on your own web server, then paste that URL.
Now, to access it from the blocking website, you give the full webserver URL of the page you've just generated, and you're done.

This follows my template and allows you with minimal effort to package your software once built`,
	Version: "1.02.00-0 (2023.10.22)",
}

var generateCmd = &cobra.Command{
	Use:     "generate full_path_to_generated_htmlfile url_to_redirect",
	Aliases: []string{"create"},
	Short:   "Generate the html page",
	Long: `Suppose you want to publish https://my_just_found_website/its_page.html to a website that blocks such contents.
All you need is to copy that URL in the clipboard, and run this software with a new web page that will be created on your own web server, then paste that URL.
Now, to access it from the blocking website, you give the full webserver URL of the page you've just generated, and you're done.

This follows my template and allows you with minimal effort to package your software once built`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			if err := generator.Create(args[0], args[1]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("USAGE: %s htmlfile URL\n", os.Args[0])
			os.Exit(0)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(generateCmd)
}
