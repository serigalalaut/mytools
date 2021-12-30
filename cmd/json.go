package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Convert file to json format",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//for mac
		//fileName := "/usr/local/var/log/nginx/error.log"

		//for linux
		fileName := "/var/log/nginx/error.log"
		if len(args) >= 1 && args[0] != "" {
			fileName = args[0]
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

				formatFile := fmt.Sprintf("error.json")
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
	rootCmd.AddCommand(jsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	jsonCmd.Flags().BoolP("tail", "t", false, "display file")
}
