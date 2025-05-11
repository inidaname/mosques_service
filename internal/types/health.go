package types

type DBStats struct {
	MaxConns     int32 `json:"max_connections"`
	UsedConns    int32 `json:"used_connections"`
	IdleConns    int32 `json:"idle_connections"`
	WaitCount    int64 `json:"wait_count"`
	WaitDuration int64 `json:"wait_duration_ms"`
	MaxIdleTime  int64 `json:"max_idle_time_ms"`
}
