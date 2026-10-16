package releaseversionutils

import (
	"os"
	"runtime/debug"
	"strings"
)

func GetReleaseVersion() string {
	if strings.EqualFold(os.Getenv("SDK_TESTING"), "true") {
		return ""
	}
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	for _, dep := range bi.Deps {
		if dep.Path == "github.com/globalpayments/go-sdk" {
			return dep.Version
		}
	}
	return ""
}
