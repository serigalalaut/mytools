package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mytools",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		//for mac
		//fileName := "/usr/local/var/log/nginx/error.log"
		format := "txt"

		//for linux
		fileName := "/var/log/nginx/error.log"
		if len(args) >= 1 && args[0] != "" {
			fileName = args[0]
		}

		if len(args) >= 1 && args[0] != "" {
			format = args[0]
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
				formatFile := fmt.Sprintf("test.%s", format)
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mytools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("tail", "t", false, "display file")
	rootCmd.Flags().BoolP("output", "o", false, "change the output file")

}
