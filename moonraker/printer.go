package moonraker

import (
	"github.com/sebasptsch/discraker/moonraker/structs"
)

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#get-klippy-host-information
func (s *Session) PrinterInfo() (structs.PrinterInfo, error) {
	return rpc[structs.PrinterInfo](s, "printer.info", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#emergency-stop
func (s *Session) PrinterEmergencyStop() (structs.Okay, error) {
	return rpc[structs.Okay](s, "printer.emergency_stop", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#host-restart
func (s *Session) PrinterRestartHost() (structs.Okay, error) {
	return rpc[structs.Okay](s, "printer.restart", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#get-klippy-host-information
func (s *Session) PrinterRestartFirmware() (structs.Okay, error) {
	return rpc[structs.Okay](s, "printer.firmware_restart", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#list-loaded-printer-objects
func (s *Session) PrinterObjectsList() (structs.PrinterObjectsList, error) {
	return rpc[structs.PrinterObjectsList](s, "printer.firmware_restart", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#query-printer-object-status
func (s *Session) PrinterObjectsQuery(params structs.PrinterObjectsQueryParams) (structs.PrinterObjectsQuery, error) {
	return rpc[structs.PrinterObjectsQuery](s, "printer.objects.query", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#subscribe-to-printer-object-status-updates
func (s *Session) PrinterObjectsSubscribe(params structs.PrinterObjectsQueryParams) (structs.PrinterObjectsQuery, error) {
	return rpc[structs.PrinterObjectsQuery](s, "printer.objects.subscribe", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#query-endstops
func (s *Session) PrinterQueryEndstopsStatus() (structs.PrinterQueryEndstopsStatus, error) {
	return rpc[structs.PrinterQueryEndstopsStatus](s, "printer.query_endstops.status", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#run-a-gcode-command
func (s *Session) PrinterGcodeScript(params structs.PrinterGcodeScriptParams) (structs.Okay, error) {
	return rpc[structs.Okay](s, "printer.gcode.script", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#get-gcode-help
func (s *Session) PrinterGcodeHelp() (structs.PrinterGcodeHelp, error) {
	return rpc[structs.PrinterGcodeHelp](s, "printer.gcode.help", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#start-a-print-job
func (s *Session) PrinterPrintStart(params structs.PrinterPrintStartParams) (structs.Okay, error) {
	return rpc[structs.Okay](s, "printer.print.start", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#pause-a-print-job
func (s *Session) PrinterPrintPause() (structs.Okay, error) {
	return rpc[structs.Okay](s, "printer.print.pause", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#resume-a-print-job
func (s *Session) PrinterPrintResume() (structs.Okay, error) {
	return rpc[structs.Okay](s, "printer.print.resume", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#resume-a-print-job
func (s *Session) PrinterPrintCancel() (structs.Okay, error) {
	return rpc[structs.Okay](s, "printer.print.cancel", nil)
}
