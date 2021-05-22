package godogs

import (
	"flag"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress", // can define default values
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	godogs := func(s *godog.Suite) {
		FeatureContext(s)
	}

	status := godog.RunWithOptions("godogs", godogs, opt)
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
