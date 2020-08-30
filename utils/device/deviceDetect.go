package device

import (
	"fmt"
	ua "github.com/mileusna/useragent"
)

func Detect(detect ua.UserAgent) string {
	device := GetType(detect)
	result := fmt.Sprintf("%v %v, OS : %v %v, Device : %v", detect.Name, detect.Version, detect.OS, detect.OSVersion, device)
	return result
}

func GetType(detect ua.UserAgent) string {
	var device string
	if detect.Mobile {
		device = "Mobile"
	}
	if detect.Tablet {
		device = "Tablet"
	}
	if detect.Desktop {
		device = "Desktop"
	}
	if detect.Bot {
		device = "Bot"
	}

	if detect.OS == "" {
		detect.OS = "null"
	}
	if device == "" {
		device = "null"
	}
	return device
}
