package env

import (
	"os"
	"path/filepath"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/joho/godotenv"
)

var log = clog.NewWithPlugin("env")

func init() {
	plugin.Register("env", setup)
}

func setup(c *caddy.Controller) error {
	n := 0
	e := EnvPlugin{}

	for c.Next() {
		if n > 0 {
			return c.Err("Unable to load config")
		}
		n++

		args := c.RemainingArgs()
		if len(args) >= 1 {
			for _, a := range args {
				LoadEnvFile(a)
				e.Paths = append(e.Paths, a)
			}
		}
	}

	return nil
}

func LoadEnvFile(path string) error {
	if !filepath.IsAbs(path) {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		path = filepath.Join(cwd, path)
	}

	err := godotenv.Load(path)
	if err != nil {
		log.Warningf("Unable to load '%s': %v", path, err)
		return err
	}

	log.Infof("Loaded environment file '%s'", path)
	return nil
}
