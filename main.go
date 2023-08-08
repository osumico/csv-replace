package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var path string
var replaced string
var field string
var out string
var delimiter string

var rcmd = &cobra.Command{
	Use:     "csvrpl",
	Version: "'beta-dev'",
	Long:    `This application take path for one/many csv and replace content on all rows`,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	rcmd.PersistentFlags().StringVarP(&path, "path", "p", "", "take path for csv")
	rcmd.MarkPersistentFlagRequired("path")
	rcmd.PersistentFlags().StringVarP(&replaced, "replaced", "r", "name", "what field of CSV need to be replaced?")
	rcmd.PersistentFlags().StringVarP(&field, "field", "f", "foo", "value to be replaced, all values be this value")
	rcmd.PersistentFlags().StringVarP(&out, "out", "o", "out.csv", "name for output file")
	rcmd.PersistentFlags().StringVarP(&delimiter, "delimetr", "d", ",", "used delimetr's in csv files")
	rcmd.PersistentFlags().BoolP("help", "h", false, "this help message")
	rcmd.PersistentFlags().BoolP("version", "v", false, "show version of application")
	rcmd.PersistentFlags().BoolP("silent", "s", false, "supress stdout")
}

func main() {
	if err := rcmd.Execute(); err != nil {
		os.Exit(1)
	}

	help, _ := rcmd.Flags().GetBool("help")
	version, _ := rcmd.Flags().GetBool("version")

	if help || version {
		os.Exit(0)
	}

	content, err := readCSV(path)

	if err != nil {
		fmt.Printf("Unnable to read file content: '%v'\n", err)
	}

	newContent, err := replace(content, field, replaced)
	if err != nil {
		fmt.Printf("Unnable to replace file content: '%v'\n", err)
	}

	isWrite, err := writeCSV(newContent, out)

	if err != nil {
		fmt.Printf("Unnable to write file: '%v'\n", err)
	}

	silent, _ := rcmd.Flags().GetBool("silent")

	if !silent && isWrite {
		fmt.Printf("File '%v' is writen!\n", out)
	}

}
