package cserver

import "github.com/bitwormhole/starter/application"

func ConfigForServer(x application.ConfigBuilder) error {
	return autoGenConfig(x)
}
