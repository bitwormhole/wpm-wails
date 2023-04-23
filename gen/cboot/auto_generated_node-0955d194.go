// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cboot

import (
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
	boot0x2acd90 "github.com/bitwormhole/wpm-wails/app/boot"
)

type pComLoader struct {
	instance *boot0x2acd90.Loader
	 markup0x23084a.Component `class:"server-loader life"`
	Port int `inject:"${server.port}"`
	WithGui bool `inject:"${wpm.options.run-with-gui}"`
	WithServer bool `inject:"${wpm.options.run-with-server}"`
	AC application0x67f6c5.Context `inject:"context"`
}

