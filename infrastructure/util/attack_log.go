package util

// AttackLog yapısı
type AttackLog struct {
	Date       string `json:"date"`
	Time       string `json:"time"`
	Type       string `json:"type"`
	Protocol   string `json:"protocol"`
	SourceIP   string `json:"source_ip"`
	SourcePort string `json:"source_port"`
	DestIP     string `json:"dest_ip"`
	DestPort   string `json:"dest_port"`
	AllowDeny  string `json:"allow/deny"`
	Priority   string `json:"priority"`
	OS         string `json:"os"`
}

// CreateAndSaveAttackLogs fonksiyonu, attack loglarını JSON formatında oluşturur ve dosyaya yazar.
func CreateAndSaveAttackLogs() *[]AttackLog {
	attackLogs := []AttackLog{
		{
			Date:       "05/27/2024",
			Time:       "22:49",
			Type:       "ICMP Ping",
			Protocol:   "ICMP",
			SourceIP:   "192.168.6.5",
			SourcePort: "8",
			DestIP:     "192.168.6.4",
			DestPort:   "0",
			AllowDeny:  "Allowed",
			Priority:   "Medium",
			OS:         "Kali Linux",
		},
		{
			Date:       "05/27/2024",
			Time:       "22:49",
			Type:       "ICMP Ping",
			Protocol:   "ICMP",
			SourceIP:   "192.168.6.4",
			SourcePort: "0",
			DestIP:     "192.168.6.5",
			DestPort:   "0",
			AllowDeny:  "Allowed",
			Priority:   "Medium",
			OS:         "Ubuntu",
		},
		{
			Date:       "05/27/2024",
			Time:       "22:50",
			Type:       "TCP SYN Scan",
			Protocol:   "TCP",
			SourceIP:   "192.168.6.5",
			SourcePort: "61756",
			DestIP:     "192.168.6.4",
			DestPort:   "53",
			AllowDeny:  "Denied",
			Priority:   "High",
			OS:         "Kali Linux",
		},
		{
			Date:       "05/31/2024",
			Time:       "21:21",
			Type:       "TCP SYN Scan",
			Protocol:   "TCP",
			SourceIP:   "192.168.6.5",
			SourcePort: "56806",
			DestIP:     "192.168.6.4",
			DestPort:   "1233",
			AllowDeny:  "Denied",
			Priority:   "High",
			OS:         "Kali Linux",
		},
		{
			Date:       "06/01/2024",
			Time:       "16:26",
			Type:       "ICMP Ping",
			Protocol:   "ICMP",
			SourceIP:   "192.168.6.5",
			SourcePort: "8",
			DestIP:     "192.168.6.4",
			DestPort:   "0",
			AllowDeny:  "Allowed",
			Priority:   "Medium",
			OS:         "Kali Linux",
		},
		{
			Date:       "06/01/2024",
			Time:       "16:26",
			Type:       "ICMP Ping",
			Protocol:   "ICMP",
			SourceIP:   "192.168.6.4",
			SourcePort: "0",
			DestIP:     "192.168.6.5",
			DestPort:   "0",
			AllowDeny:  "Allowed",
			Priority:   "Medium",
			OS:         "Ubuntu",
		},
		{
			Date:       "06/01/2024",
			Time:       "16:27",
			Type:       "SSH Brute Force Attack",
			Protocol:   "TCP",
			SourceIP:   "192.168.6.5",
			SourcePort: "38100",
			DestIP:     "192.168.6.4",
			DestPort:   "22",
			AllowDeny:  "Denied",
			Priority:   "High",
			OS:         "Kali Linux",
		},
		{
			Date:       "06/01/2024",
			Time:       "16:29",
			Type:       "SSH Brute Force Attack",
			Protocol:   "TCP",
			SourceIP:   "192.168.6.5",
			SourcePort: "54872",
			DestIP:     "192.168.6.4",
			DestPort:   "22",
			AllowDeny:  "Denied",
			Priority:   "High",
			OS:         "Kali Linux",
		},
	}
	return &attackLogs
}
