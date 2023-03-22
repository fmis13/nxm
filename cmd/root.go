/*
Copyright © 2023 [github.com/nixmember]
*/

package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/nixmember/nixmember/programs"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(InstallCmd)
	rootCmd.AddCommand(UninstallCmd)
	rootCmd.AddCommand(programsCmd)
	rootCmd.AddCommand(UpdateCmd)

	InstallCmd.AddCommand(programs.Bitwarden)
	InstallCmd.AddCommand(programs.HomeAssistant)
	InstallCmd.AddCommand(programs.Jellyfin)
	InstallCmd.AddCommand(programs.Nextcloud)

	UninstallCmd.AddCommand(programs.BitwardenUninstallation)
	UninstallCmd.AddCommand(programs.HomeAssistantUninstallation)
	UninstallCmd.AddCommand(programs.JellyfinUninstallation)
	UninstallCmd.AddCommand(programs.NextcloudUninstallation)

	UpdateCmd.AddCommand(programs.BitwardenUpdate)
	UpdateCmd.AddCommand(programs.HomeAssistantUpdate)
	UpdateCmd.AddCommand(programs.JellyfinUpdate)
	UpdateCmd.AddCommand(programs.NextcloudUpadte)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nxm",
	Short: "The package manager for linux (home-)servers.",
	Long:  `error: no operation specified (use --h for help)`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		figure.NewFigure("nxm", "", true).Print()
		fmt.Printf("\nnxm v1.0.0 - a command line tool for (home-)servers that aims to make self-hosting easier.")
		fmt.Printf("Copyright © 2022-2023 NixMember Team \n")
		fmt.Println("This program is open source, you can freely redestribute it under the terms of the MIT license.")
		fmt.Println("You can find the source code at https://github.com/nixmember/nixmember")
	},
}

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a program",
	Long:  "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println("Something went wrong, please report this to the issues page on GitHub.")
		}
		if currentUser.Username != "root" {
			fmt.Println("Please run this command as root.")
			os.Exit(1)
		}
	},
	Args: cobra.MinimumNArgs(1),
}

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update",
	Long:  "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println("Something went wrong, please report this to the issues page on GitHub.")
		}
		if currentUser.Username != "root" {
			fmt.Println("Please run this command as root.")
			os.Exit(1)
		}
	},
	Args: cobra.MinimumNArgs(1),
}

var UninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a program",
	Long:  "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println("Something went wrong, please report this.")
		}
		if currentUser.Username != "root" {
			fmt.Println("Please run this command as root.")
			os.Exit(3)
		}
	},
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Please supply a program to uninstall \n")
	},
}

var programsCmd = &cobra.Command{
	Use:   "programs",
	Short: "List all avalible programs",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Programs currently avalible:\n")
		fmt.Print("Bitwarden - a open source password manager - ")
		color.Blue("`bitwarden`")
		fmt.Print("Jellyfin - The Free Software Media System - ")
		color.Magenta("`jellyfin`")
		fmt.Print("Home Assistant - open source home automation that puts local control and privacy first - ")
		color.Cyan("`homeassistant`")
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
