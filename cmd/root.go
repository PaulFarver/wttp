/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/paulfarver/wttp/pkg/codes"
	"github.com/spf13/cobra"
	"gopkg.in/gookit/color.v1"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wttp",
	Short: "A small tool to find appropriate http codes",
	Long: `A tool you can use to figure out:
- What http status code you want to use in an endpoint
- What an http code means, when you're not sure`,
	Run: func(cmd *cobra.Command, args []string) {
		listCommand()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "list",
		Short:   "Prints all codes",
		Long:    "Lists all the available http codes, and their message",
		Aliases: []string{"l"},
		Run: func(cmd *cobra.Command, args []string) {
			keys := make([]string, 0, len(codes.Codes))
			for k := range codes.Codes {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				renderShort(k, codes.Codes[k])
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:     "search",
		Short:   "Searches among the codes",
		Long:    "Finds codes with a specified string in the code, short or long description",
		Args:    cobra.MinimumNArgs(1),
		Aliases: []string{"s"},
		Run: func(cmd *cobra.Command, args []string) {
			for key, code := range codes.Codes {
				if strings.Contains(strings.ToLower(fmt.Sprintf("%s - %s\n%s\n", key, code.Message, code.Description)), strings.ToLower(args[0])) {
					renderLong(key, code)
				}
			}
		},
	})

	for c, co := range codes.Codes {
		m, mo := c, co
		rootCmd.AddCommand(&cobra.Command{
			Use:    c,
			Short:  co.Message,
			Long:   co.Description,
			Hidden: true,
			Run: func(cmd *cobra.Command, args []string) {
				renderLong(m, mo)
			},
		})
	}
}

func listCommand() {
	keys := make([]string, 0, len(codes.Codes))
	for k := range codes.Codes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		renderShort(k, codes.Codes[k])
	}
}

func renderLong(code string, details codes.Code) {
	col := getColor(code)
	col.Printf("%s - %s\n", code, details.Message)
	fmt.Printf("%s\n", details.Description)
}

func renderShort(code string, details codes.Code) {
	col := getColor(code)
	col.Printf("%s - %s\n", code, details.Message)
}

func getColor(code string) color.Color {
	switch string(code[0]) {
	case "1":
		return color.Blue
	case "2":
		return color.Green
	case "3":
		return color.Cyan
	case "4":
		return color.Yellow
	case "5":
		return color.Red
	default:
		return color.Normal
	}
}
