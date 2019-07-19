package config

import "os"

type config struct {
	workingDir string
}

var (
	c config
)

func init() {
	c = config{
		workingDir: os.Getenv("APP_WORK_DIR"),
	}
}

func GetWD() string {
	return c.workingDir
}
