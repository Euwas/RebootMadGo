package rebootmadgo

import "github.com/rs/zerolog/log"

type Device struct {
	Name        string
	active      bool
	powerOffURL string
	owerOnURL   string
}

func (d *Device) SetActive(active bool) {
	d.active = active
	if !active {
		d.reboot()
	}
}

func (d *Device) reboot() {
	// Call powerOff url
	log.Info().Str("device_name", d.Name).Msg("Calling poweroff url")

	// Call poweron url
	log.Info().Str("device_name", d.Name).Msg("Calling poweron url")
}
