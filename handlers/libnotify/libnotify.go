package libnotify

import (
	"github.com/lucas59356/http2growl/utils"
	"github.com/mqu/go-notify"
	"github.com/pkg/errors"
)

// ErrInitNotOK Initialization of libnotify was not sucessful
var ErrInitNotOK = errors.New("Init not ok")

// Libnotify Notification wrapper around libnotify
type Libnotify struct {
	Delay int
}

// NewLibnotify Creates a new instance
func NewLibnotify(appname string, delay int) (*Libnotify, error) {
	l := Libnotify{
		Delay: delay,
	}
	ok := notify.Init(appname)
	if !ok {
		return nil, ErrInitNotOK
	}
	return &l, nil
}

// Notify Do the comunication with
func (l *Libnotify) Notify(msg *utils.Message) error {
	m := notify.NotificationNew(msg.Title, msg.Message, msg.Icon)
	m.SetTimeout(int32(l.Delay))
	err := m.Show()
	if err != nil {
		return err
	}
	return nil
}
