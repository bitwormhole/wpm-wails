package cboot

import "github.com/bitwormhole/starter/application"

func ConfigForBoot(x application.ConfigBuilder) error {
	return autoGenConfig(x)
}
