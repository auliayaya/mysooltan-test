package cli

import (
	"fmt"
	"os"

	"github.com/auliayaya/mysooltan/internal/service"
	"github.com/spf13/cobra"
)

var (
	path   string
	saveTo string
	types  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mysooltan",
	Short: "My Sooltan CLI",
	Long:  `this tools is for converting log to json or txt format file`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		service.NewService().ConvertLogFile(path, types, saveTo)
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mysooltan.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(convCmd)
	rootCmd.CommandPath()
	rootCmd.Flags().StringVarP(&path, "path", "p", "", "input ini wajib diisi untuk memilih file yang akan dikonversikan")
	if err := rootCmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}
	rootCmd.Flags().StringVarP(&saveTo, "out", "o", "", "inputan directory path untuk save output file")
	rootCmd.Flags().StringVarP(&types, "types", "t", "", "inputan json,txt memilih tipe konversi json atau txt")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
