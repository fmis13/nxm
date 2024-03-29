package programs

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

//go:embed scripts/home-assistant.sh
var homeAssistantScript string

var HomeAssistant = &cobra.Command{
	Use: "home-assistant",
	Run: func(cmd *cobra.Command, args []string) {
		homeAssistantInstallation()
	},
}

var HomeAssistantUninstallation = &cobra.Command{
	Use: "home-assistant",
	Run: func(cmd *cobra.Command, args []string) {
		homeAssistantUninstallation()
	},
}

var HomeAssistantUpdate = &cobra.Command{
	Use: "home-assistant",
	Run: func(cmd *cobra.Command, args []string) {
		homeAssistantUpdate()
	},
}

func homeAssistantInstallation() {

	cmd := exec.Command("/bin/bash", "-c", homeAssistantScript)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func homeAssistantUninstallation() {
	cmd := exec.Command("/bin/bash", "-c", "rm -rf $HOME/.config/nxm/home-assistant")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func homeAssistantUpdate() {
	cmd := exec.Command("/bin/bash", "-c", "cd $HOME/.config/nxm/home-assistant; docker compose pull")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func init() {
}
