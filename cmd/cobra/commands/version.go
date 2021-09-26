package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of Hugo",
	Long: "All software has versions. This is Hugo's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		fmt.Println(cmd.Flags().Lookup("main").Value)
	},
}

func init()  {
	versionCmd.Flags().StringP("main", "m", "v1", "大版本")
	rootCmd.AddCommand(versionCmd)
}
