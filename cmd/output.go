package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//for mac
		//fileName := "/usr/local/var/log/nginx/error.log"
		Output := "~/Docker/development/mytools/error.txt"
		//for linux
		fileName := "/var/log/nginx/error.log"
		if len(args) == 3 {

			if args[1] != "" {
				fileName = args[1]
			}

			if args[2] != "" {
				Output = args[2]
			}
		} else {

			if args[0] != "" {
				fileName = args[0]
			}

			if args[1] != "" {
				Output = args[1]
			}
		}

		f, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		defer f.Close()
		buf := make([]byte, 1024)
		for {
			n, err := f.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			if n > 0 {

				formatFile := fmt.Sprintf("%s", Output)
				f, err := os.Create(formatFile)

				if err != nil {
					log.Fatal(err)
				}

				defer f.Close()
				_, err2 := f.WriteString(string(buf[:n]))

				if err2 != nil {
					log.Fatal(err2)
				}

			}
		}
		fmt.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(outputCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// outputCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	outputCmd.Flags().BoolP("output", "o", false, "change directory output")
}
