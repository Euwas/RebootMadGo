package controller

import (
	"time"

	"github.com/euwas/rebootmadgo/internal/config"
	"github.com/euwas/rebootmadgo/pkg/rebootmadgo"
	"github.com/rs/zerolog/log"
)

type DeviceWatcher struct {
	Devices []rebootmadgo.Device
	Config  *config.Config
}

func (dw *DeviceWatcher) Length() int {
	return len(dw.Devices)
}

func (dw *DeviceWatcher) Update() {
	statusList, err := getMADStatus(dw.Config)

	if err != nil {
		log.Error().
			Err(err).
			Msg("Error loading MAD status list")
	}

	for i, s := range statusList {
		log.Debug().Msgf("[%d] `%+v`\n", i, s)
	}

	dw.updateActiveList(statusList)
	log.Debug().Msgf("%+v", dw.Devices)

}

func (dw *DeviceWatcher) updateActiveList(deviceStatuses DeviceStatusList) {

	// Oldest time still acceptable to consider online and active
	minimumTime := time.Now().Add(-5 * time.Minute)
	log.Debug().Time("minimumTime", minimumTime).Msg("Minimum time for active")

	for i, d := range dw.Devices {
		deviceStatus := findDeviceStatusByName(d.Name, deviceStatuses)
		if deviceStatus == nil {
			continue
		}

		dw.Devices[i].SetActive(deviceStatus.LastProtoDateTime.After(minimumTime))
	}
}

func DeviceWatcherFromConfig(config *config.Config) DeviceWatcher {
	var dw DeviceWatcher
	dw.Config = config
	dw.Devices = make([]rebootmadgo.Device, len(config.Devices))

	for i := range config.Devices {
		dw.Devices[i].Name = config.Devices[i].Name
	}

	return dw
}
