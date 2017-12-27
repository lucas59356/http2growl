package growlhandler

import (
	"github.com/lucas59356/http2growl/utils"
	"github.com/mattn/go-gntp"
)

// GNTP handler
type GNTP struct {
	Client        *gntp.Client
	Notifications map[string]gntp.Notification
}

// NewGNTP Creates a new GNTP connection
func NewGNTP(server string) *GNTP {
	client := gntp.NewClient()
	client.AppName = "http2growl"
	client.Server = server
	return &GNTP{
		Client:        client,
		Notifications: map[string]gntp.Notification{},
	}
}

// Notify Sends a notification to Growl
func (g *GNTP) Notify(msg *utils.Message) error {
	m := gntp.Message{
		DisplayName: msg.From,
		Event:       msg.Event,
		Icon:        msg.Icon,
		Sticky:      msg.Sticky,
		Text:        msg.Message,
		Title:       msg.Title,
	}
	n := gntp.Notification{
		DisplayName: msg.From,
		Event:       msg.Event,
		Enabled:     true,
	}
	err := g.Register(n)
	if err != nil {
		return err
	}
	err = g.Client.Notify(&m)
	if err != nil {
		return err
	}
	return nil
}

// Register the events in growl
func (g *GNTP) Register(n gntp.Notification) error {
	g.Notifications[n.Event] = n
	var ns []gntp.Notification
	for _, v := range g.Notifications {
		ns = append(ns, v)
	}
	return g.Client.Register(ns)
}
