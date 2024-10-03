//go:build !(android && cmfa)

package process

import "github.com/kitty314/1.17.0/constant"

func FindPackageName(metadata *constant.Metadata) (string, error) {
	return "", nil
}
