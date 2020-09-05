package utils

// OhmData : data structrue for OpenHardwareMonitor
type OhmData struct {
	Name     string    `json:"Text"`
	Children []OhmData `json:"Children"`
	ID       int       `json:"id"`
	Min      string    `json:"Min"`
	Value    string    `json:"Value"`
	Max      string    `json:"Max"`
	ImageURL string    `json:"ImageURL"`
}
