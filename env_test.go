package env

import (
	vlog "log"
	"os"
	"testing"

	"github.com/coredns/caddy"
	clog "github.com/coredns/coredns/plugin/pkg/log"
)

func TestEnvPlugin(tst *testing.T) {
	OverwriteStdOut()
	clog.D.Set()

	c := caddy.NewTestController("dns", `
		env .env
	`)

	err := setup(c)
	if err != nil {
		tst.Errorf("Unable to complete setup: %v", err)
	}
}

func OverwriteStdOut() error {
	tempFile, err := os.CreateTemp("", "coredns-env-plugin-log")
	if err != nil {
		return err
	}

	defer os.Remove(tempFile.Name())

	orig := log
	log = clog.NewWithPlugin("env")
	vlog.SetOutput(os.Stdout)

	defer func() {
		log = orig
	}()

	return nil
}
