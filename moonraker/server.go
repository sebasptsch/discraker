package moonraker

import (
	"github.com/sebasptsch/discraker/moonraker/structs"
)

func (s *Session) ServerWebcamsList() (structs.ServerWebcamList, error) {
	var reply structs.ServerWebcamList

	err := s.Call("server.webcams.list", nil, &reply)
	if err != nil {
		return structs.ServerWebcamList{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#query-server-info
func (s *Session) ServerInfo() (structs.ServerInfo, error) {
	var reply structs.ServerInfo

	err := s.Call("server.info", nil, &reply)
	if err != nil {
		return structs.ServerInfo{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#get-server-configuration
func (s *Session) ServerConfig() (structs.ServerConfig, error) {
	var reply structs.ServerConfig

	err := s.Call("server.config", nil, &reply)
	if err != nil {
		return structs.ServerConfig{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#request-cached-temperature-data
func (s *Session) ServerTemperatureStore(params structs.ServerTemperatureStoreParams) (structs.ServerTemperatureStore, error) {
	var reply structs.ServerTemperatureStore

	err := s.Call("server.temperature_store", params, &reply)

	if err != nil {
		return structs.ServerTemperatureStore{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#request-cached-gcode-responses
func (s *Session) ServerGcodeStore(params structs.ServerGcodeStoreParams) (structs.ServerGcodeStore, error) {
	var reply structs.ServerGcodeStore

	err := s.Call("server.gcode_store", params, &reply)

	if err != nil {
		return structs.ServerGcodeStore{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#rollover-logs
func (s *Session) ServerRolloverLogs(params structs.ServerRolloverLogsParams) (structs.ServerRolloverLogs, error) {
	var reply structs.ServerRolloverLogs

	err := s.Call("server.logs.rollover", params, &reply)

	if err != nil {
		return structs.ServerRolloverLogs{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#restart-server
func (s *Session) ServerRestart() (structs.Okay, error) {
	var reply structs.Okay

	err := s.Call("server.restart", nil, &reply)

	if err != nil {
		return "", err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/server/#identify-connection
func (s *Session) ServerConnectionIdentify(params structs.ServerConnectionIdentifyParams) (structs.ServerConnectionIdentify, error) {
	var reply structs.ServerConnectionIdentify

	err := s.Call("server.connection.identify", params, &reply)

	if err != nil {
		return structs.ServerConnectionIdentify{}, err
	}

	return reply, nil
}
