package main

import (
	"runtime"

	"github.com/pkg/errors"

	"github.com/lucas59356/http2growl/handlers/growl"
	"github.com/lucas59356/http2growl/handlers/libnotify"
	"github.com/lucas59356/http2growl/utils"
)

// ErrSysNotSupported Op System not supported by thsi program
var ErrSysNotSupported = errors.New("System not supported")

// GetHandler Returns the notify function according the system
func GetHandler() (func(*utils.Message) error, error) {
	println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		g := growl.NewGNTP(AppName, GNTPServer)
		return g.Notify, nil
	}
	if runtime.GOOS == "linux" {
		l, err := libnotify.NewLibnotify(AppName, Delay)
		return l.Notify, err
	}

	return func(*utils.Message) error { return nil },
		ErrSysNotSupported
}
