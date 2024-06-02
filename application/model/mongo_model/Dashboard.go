package mongo_model

type DashboardData struct {
	Alerts     []Alert   `json:"alerts"`
	Whitelists []List    `json:"whitelists"`
	Blacklists []List    `json:"blacklists"`
	Devices    []Device  `json:"devices"`
	Requests   []Request `json:"requests"`
}
