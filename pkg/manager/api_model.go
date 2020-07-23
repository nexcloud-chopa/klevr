package manager

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Klevry/klevr/pkg/common"
)

type CustomHeader struct {
	APIKey         string `header:"X-API-KEY"`
	AgentKey       string `header:"X-AGENT-KEY"`
	HashCode       string `header:"X-HASH-CODE"`
	ZoneID         uint64 `header:"X-ZONE-ID"`
	SupportVersion string `header:"X-SUPPORT-AGENT-VERSION"`
}

type Primary struct {
	IP      string `json:"ip"`
	Running bool   `json:"running"`
}

type Agent struct {
	IP       string `json:"ip"`
	Running  bool   `json:"running"`
	AgentKey string `json:"agentKey"`
}

type Resource struct {
	Core   int `json:"core"`
	Memory int `json:"memory"`
	Disk   int `json:"disk"`
}

type Task struct {
}

func getCustomHeader(r *http.Request) *CustomHeader {
	zoneID, _ := strconv.ParseUint(strings.Join(r.Header.Values("X-ZONE-ID"), ""), 10, 64)

	return &CustomHeader{
		APIKey:         strings.Join(r.Header.Values("X-API-KEY"), ""),
		AgentKey:       strings.Join(r.Header.Values("X-AGENT-KEY"), ""),
		HashCode:       strings.Join(r.Header.Values("X-HASH-CODE"), ""),
		ZoneID:         zoneID,
		SupportVersion: strings.Join(r.Header.Values("X-SUPPORT-AGENT-VERSI"), ""),
	}
}

func addCustomHeader(w *common.ResponseWrapper, h CustomHeader) {

}
