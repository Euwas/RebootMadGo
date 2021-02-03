package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/euwas/rebootmadgo/internal/config"
)

const (
	madStatusURL = "get_status"
)

// DeviceStatus is the MAD representation of one connect MAD device
type DeviceStatus struct {
	AreaID                int    `json:"area_id"`
	CurrentPos            string `json:"currentPos"`
	CurrentSleepTime      int    `json:"currentSleepTime"`
	DeviceID              int    `json:"device_id"`
	GlobalRebootCount     int    `json:"globalrebootcount"`
	GlobalRestartCount    int    `json:"globalrestartcount"`
	Init                  bool   `json:"init"`
	LastPogoRebootUnix    int64  `json:"lastPogoReboot"`
	LastPogoRestartUnix   int64  `json:"lastPogoRestart"`
	LastPos               string `json:"lastPos"`
	LastProtoDateTime     time.Time
	LastProtoDateTimeUnix int64  `json:"lastProtoDateTime"`
	Mode                  string `json:"mode"`
	Name                  string `json:"name"`
	RebootCounter         int    `json:"rebootCounter"`
	RebootingOption       int    `json:"rebootingOption"`
	RestartCounter        int    `json:"restartCounter"`
	RMName                string `json:"rmname"`
	RouteMax              int    `json:"routeMax"`
	RoutePos              int    `json:"routePos"`
}

// DeviceStatusList is a list of multiple DeviceStatus entities
type DeviceStatusList = []DeviceStatus

func readMadStatus(body []byte) DeviceStatusList {
	var statusList []DeviceStatus
	json.Unmarshal(body, &statusList)

	for i := range statusList {
		statusList[i].LastProtoDateTime = time.Unix(statusList[i].LastProtoDateTimeUnix, 0)
	}

	return statusList
}

func getMADStatus(config *config.Config) (DeviceStatusList, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.Mad.Address+"/"+madStatusURL, nil)
	req.SetBasicAuth(config.Mad.Username, config.Mad.Password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	statusList := readMadStatus(body)
	return statusList, nil
}

func findDeviceStatusByName(name string, devices DeviceStatusList) *DeviceStatus {
	var ds *DeviceStatus = nil

	for i := range devices {
		if devices[i].Name == name {
			ds = &devices[i]
		}
	}

	return ds
}
