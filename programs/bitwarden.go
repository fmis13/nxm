package programs

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
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
	cmd := exec.Command("/bin/bash", "-c", bitwardenScript)
	stdin, err := cmd.StdinPipe()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("An error ocurred")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	defer stdin.Close()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("START")

	if err = cmd.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	}

	//io.WriteString(stdin, "4\n")
	cmd.Wait()
	fmt.Println("END") //for debug

	//fmt.Println(string(stdout))
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
