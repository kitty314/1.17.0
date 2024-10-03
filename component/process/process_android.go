//go:build android && cmfa

package process

import "github.com/kitty314/1.17.0/constant"

type PackageNameResolver func(metadata *constant.Metadata) (string, error)

var DefaultPackageNameResolver PackageNameResolver

func FindPackageName(metadata *constant.Metadata) (string, error) {
	if resolver := DefaultPackageNameResolver; resolver != nil {
		return resolver(metadata)
	}
	return "", ErrPlatformNotSupport
}
