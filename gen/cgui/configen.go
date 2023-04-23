package cgui

import "github.com/bitwormhole/starter/application"

func ConfigForGUI(x application.ConfigBuilder) error {
	return autoGenConfig(x)
}
