package packageutils

import (
	"github.com/globalpayments/go-sdk/api/utils/extrautils"
	"runtime/debug"
)

func GetPackageVersion() string {
	_ = extrautils.IfThenElse
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
