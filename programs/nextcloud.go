package programs

import (
	_ "embed"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

//go:embed scripts/nextcloud.sh
var nextcloudScript string

var Nextcloud = &cobra.Command{
	Use: "nextcloud",
	Run: func(cmd *cobra.Command, args []string) {
		nextcloudInstallation()
	},
}

var NextcloudUninstallation = &cobra.Command{
	Use: "home-assistant",
	Run: func(cmd *cobra.Command, args []string) {
		nextcloudUninstallation()
	},
}

var NextcloudUpdate = &cobra.Command{
	Use: "home-assistant",
	Run: func(cmd *cobra.Command, args []string) {
		nextcloudUpdate()
	},
}

func nextcloudInstallation() {

	cmd := exec.Command("/bin/bash", "-c", nextcloudScript)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func nextcloudUninstallation() {
	cmd := exec.Command("/bin/bash", "-c", "rm -rf $HOME/nextcloud")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func nextcloudUpdate() {
	cmd := exec.Command("/bin/bash", "-c", "echo 'For updating nextcloud, please see the documentation'")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func init() {
}
