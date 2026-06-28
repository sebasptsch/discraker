package moonraker

import (
	"github.com/sebasptsch/discraker/moonraker/structs"
)

func (s *Session) ServerWebcamsList() (structs.ServerWebcamList, error) {
	return rpc[structs.ServerWebcamList](s, "server.webcams.list", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#query-server-info
func (s *Session) ServerInfo() (structs.ServerInfo, error) {
	return rpc[structs.ServerInfo](s, "server.info", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#get-server-configuration
func (s *Session) ServerConfig() (structs.ServerConfig, error) {
	return rpc[structs.ServerConfig](s, "server.config", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#request-cached-temperature-data
func (s *Session) ServerTemperatureStore(params structs.ServerTemperatureStoreParams) (structs.ServerTemperatureStore, error) {
	return rpc[structs.ServerTemperatureStore](s, "server.temperature_store", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#request-cached-gcode-responses
func (s *Session) ServerGcodeStore(params structs.ServerGcodeStoreParams) (structs.ServerGcodeStore, error) {
	return rpc[structs.ServerGcodeStore](s, "server.gcode_store", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#rollover-logs
func (s *Session) ServerRolloverLogs(params structs.ServerRolloverLogsParams) (structs.ServerRolloverLogs, error) {
	return rpc[structs.ServerRolloverLogs](s, "server.logs.rollover", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#restart-server
func (s *Session) ServerRestart() (structs.Okay, error) {
	return rpc[structs.Okay](s, "server.restart", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#identify-connection
func (s *Session) ServerConnectionIdentify(params structs.ServerConnectionIdentifyParams) (structs.ServerConnectionIdentify, error) {
	return rpc[structs.ServerConnectionIdentify](s, "server.connection.identify", params)
}
