package moonrakerclient

import "github.com/sebasptsch/discraker/moonraker-client/structs"

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#get-system-info
func (s *Session) MachineSystemInfo() (structs.MachineSystemInfo, error) {
	return rpc[structs.MachineSystemInfo](s, "machine.system_info", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#shutdown-the-operating-system
func (s *Session) MachineShutdown() (structs.Okay, error) {
	return rpc[structs.Okay](s, "machine.shutdown", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#reboot-the-operating-system
func (s *Session) MachineReboot() (structs.Okay, error) {
	return rpc[structs.Okay](s, "machine.reboot", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#restart-a-system-service
func (s *Session) MachineServicesRestart(params structs.MachineServicesRestartParams) (structs.Okay, error) {
	return rpc[structs.Okay](s, "machine.services.restart", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#stop-a-system-service
func (s *Session) MachineServicesStop(params structs.MachineServicesRestartParams) (structs.Okay, error) {
	return rpc[structs.Okay](s, "machine.services.stop", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#start-a-system-service
func (s *Session) MachineServicesStart(params structs.MachineServicesRestartParams) (structs.Okay, error) {
	return rpc[structs.Okay](s, "machine.services.start", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#get-process-statistics
func (s *Session) MachineProcStats() (structs.MachineProcStats, error) {
	return rpc[structs.MachineProcStats](s, "machine.proc_stats", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#get-sudo-info
func (s *Session) MachineSudoInfo(params structs.MachineSudoInfoParams) (structs.MachineSudoInfo, error) {
	return rpc[structs.MachineSudoInfo](s, "machine.sudo.info", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#set-sudo-password
func (s *Session) MachineSudoPassword(params structs.MachineSudoPasswordParams) (structs.MachineSudoPassword, error) {
	return rpc[structs.MachineSudoPassword](s, "machine.sudo.password", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#list-usb-devices
func (s *Session) MachinePeriphalsUSB() (structs.MachinePeripheralsUSB, error) {
	return rpc[structs.MachinePeripheralsUSB](s, "machine.peripherals.usb", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#list-serial-devices
func (s *Session) MachinePeripheralsSerial() (structs.MachinePeripheralsSerial, error) {
	return rpc[structs.MachinePeripheralsSerial](s, "machine.peripherals.serial", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#list-video-capture-devices
func (s *Session) MachinePeripheralsVideo() (structs.MachinePeripheralsVideo, error) {
	return rpc[structs.MachinePeripheralsVideo](s, "machine.peripherals.video", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/machine/#query-unassigned-canbus-uuids
func (s *Session) MachinePeripheralsCanbus(params structs.MachinePeripheralsCanbusParams) (structs.MachinePeripheralsCanbus, error) {
	return rpc[structs.MachinePeripheralsCanbus](s, "machine.peripherals.canbus", params)
}
