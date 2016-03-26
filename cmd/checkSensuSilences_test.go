package cmd

import (
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestchronyStats(t *testing.T) {

	Convey("When executing 'chronyc tracking'", func() {
		chronyStats := exec.Command("chronyc", "tracking")
		out, _ := chronyStats.Output()
		chronyStats.Start()

		Convey("The command should return some output", func() {
			So(out, ShouldNotBeBlank)
		})
	})
}
