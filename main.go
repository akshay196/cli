package main

import (
	"runtime"

	"github.com/RafaySystems/rcloud-cli/cmd"
	"github.com/RafaySystems/rcloud-cli/pkg/versioninfo"
)

var (
	version  string
	time     string
	arch     string
	buildNum string
)

func main() {
	if arch == "" {
		arch = runtime.GOOS + "/" + runtime.GOARCH
	}
	versioninfo.Init(version, buildNum, time, arch)
	cmd.Execute()
}