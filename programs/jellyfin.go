package programs

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

//go:embed scripts/jellyfin.sh
var jellyfinScript string

var Jellyfin = &cobra.Command{
	Use: "jellyfin",
	Run: func(cmd *cobra.Command, args []string) {
		jellyfinInstallation()
	},
}

var JellyfinUninstallation = &cobra.Command{
	Use: "jellyfin",
	Run: func(cmd *cobra.Command, args []string) {
		jellyfinUninstallation()
	},
}

var JellyfinUpdate = &cobra.Command{
	Use: "jellyfin",
	Run: func(cmd *cobra.Command, args []string) {
		jellyfinUpdate()
	},
}

func jellyfinInstallation() {

	cmd := exec.Command("/bin/bash", "-c", jellyfinScript)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func jellyfinUninstallation() {
	cmd := exec.Command("/bin/bash", "-c", "rm -rf $HOME/.config/nxm/jellyfin")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func jellyfinUpdate() {
	cmd := exec.Command("/bin/bash", "-c", "cd $HOME/.config/nxm/jellyfin; docker compose pull")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func init() {
}
