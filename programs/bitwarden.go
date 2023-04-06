package programs

import (
	_ "embed"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

//go:embed scripts/bitwarden.sh
var bitwardenScript string

var Bitwarden = &cobra.Command{
	Use: "bitwarden",
	Run: func(cmd *cobra.Command, args []string) {
		bitwardenInstallation()
	},
}

var BitwardenUninstallation = &cobra.Command{
	Use: "bitwarden",
	Run: func(cmd *cobra.Command, args []string) {
		bitwardenUninstallation()
	},
}

var BitwardenUpdate = &cobra.Command{
	Use: "bitwarden",
	Run: func(cmd *cobra.Command, args []string) {
		bitwardenUpdate()
	},
}

func bitwardenInstallation() {
	cmd := exec.Command(".", bitwardenScript)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func bitwardenUninstallation() {
	cmd := exec.Command("/bin/bash", "-c", "userdel bitwarden; rm -rf /opt/bitwarden")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func bitwardenUpdate() {
	cmd := exec.Command("/bin/bash", "-c", "runuser -l bitwarden -c './bitwarden.sh update'")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func init() {
}
