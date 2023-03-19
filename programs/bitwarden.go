package programs

import (
	_ "embed"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
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

func bitwardenInstallation() {

	p := mpb.New(mpb.WithWidth(64))

	total := 100
	name := "Bar: "
	bar := p.New(int64(total),
		// BarFillerBuilder with custom style
		mpb.BarStyle().Lbound("[").Filler("=").Tip("=>").Padding("-").Rbound("]"),
		mpb.PrependDecorators(
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
			),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)
	// simulating some work
	cmd := exec.Command("/bin/bash", "-c", bitwardenScript)
	bar.Increment()

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	// wait for our bar to complete and flush
	p.Wait()

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

func init() {
}
