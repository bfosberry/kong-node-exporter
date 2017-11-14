package kong

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

var errEmptyStatus = errors.New("status_empty")

// Status holds the structued info from the Kong status API
// This indicates the health and throughput of that kong proxy node
type Status struct {
	Database struct {
		Reachable bool `json:"reachable"`
	} `json:"database"`
	Server struct {
		ConnectionsWriting  int `json:"connections_writing"`
		TotalRequests       int `json:"total_requests"`
		ConnectionsHandled  int `json:"connections_handled"`
		ConnectionsAccepted int `json:"connections_accepted"`
		ConnectionsReading  int `json:"connections_reading"`
		ConnectionsActive   int `json:"connections_active"`
		ConnectionsWaiting  int `json:"connections_waiting"`
	} `json:"server"`
}

// MetricsReader returns an io.Reader with the current metrics info available
// on it, which can be consumed by any source
func (s *Status) MetricsReader() (io.Reader, error) {
	if s == nil {
		return nil, errEmptyStatus
	}
	buf := bytes.NewBufferString("")
	if err := writeMetric(buf, "kong_database_reachable", boolAsInt(s.Database.Reachable)); err != nil {
		return nil, err
	}
	if err := writeMetric(buf, "kong_total_requests", s.Server.TotalRequests); err != nil {
		return nil, err
	}
	if err := writeMetric(buf, "kong_connections_writing", s.Server.ConnectionsWriting); err != nil {
		return nil, err
	}
	if err := writeMetric(buf, "kong_connections_handled", s.Server.ConnectionsHandled); err != nil {
		return nil, err
	}
	if err := writeMetric(buf, "kong_connections_accepted", s.Server.ConnectionsAccepted); err != nil {
		return nil, err
	}
	if err := writeMetric(buf, "kong_connections_reading", s.Server.ConnectionsReading); err != nil {
		return nil, err
	}
	if err := writeMetric(buf, "kong_connections_active", s.Server.ConnectionsActive); err != nil {
		return nil, err
	}
	if err := writeMetric(buf, "kong_connections_waiting", s.Server.ConnectionsWaiting); err != nil {
		return nil, err
	}
	return buf, nil
}

func writeMetric(buf *bytes.Buffer, key string, value int) error {
	metricString := fmt.Sprintf("%s %d\n", key, value)
	_, err := buf.WriteString(metricString)
	return err
}

func boolAsInt(val bool) int {
	ret := 0
	if val {
		ret = 1
	}
	return ret
}
