package moonraker

import (
	"github.com/sebasptsch/discraker/moonraker/structs"
)

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#get-klippy-host-information
func (s *Session) PrinterInfo() (structs.PrinterInfo, error) {
	var reply structs.PrinterInfo

	err := s.Call("printer.info", nil, &reply)
	if err != nil {
		return structs.PrinterInfo{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#emergency-stop
func (s *Session) PrinterEmergencyStop() (structs.Okay, error) {
	var reply structs.Okay

	err := s.Call("printer.emergency_stop", nil, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#host-restart
func (s *Session) PrinterRestartHost() (structs.Okay, error) {
	var reply structs.Okay

	err := s.Call("printer.restart", nil, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#get-klippy-host-information
func (s *Session) PrinterRestartFirmware() (structs.Okay, error) {
	var reply structs.Okay

	err := s.Call("printer.firmware_restart", nil, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#list-loaded-printer-objects
func (s *Session) PrinterObjectsList() (structs.PrinterObjectsList, error) {
	var reply structs.PrinterObjectsList

	err := s.Call("printer.objects.list", nil, &reply)
	if err != nil {
		return structs.PrinterObjectsList{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#query-printer-object-status
func (s *Session) PrinterObjectsQuery(params structs.PrinterObjectsQueryParams) (structs.PrinterObjectsQuery, error) {
	var reply structs.PrinterObjectsQuery

	err := s.Call("printer.objects.query", params, &reply)
	if err != nil {
		return structs.PrinterObjectsQuery{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#subscribe-to-printer-object-status-updates
func (s *Session) PrinterObjectsSubscribe(params structs.PrinterObjectsQueryParams) (structs.PrinterObjectsQuery, error) {
	var reply structs.PrinterObjectsQuery

	err := s.Call("printer.objects.subscribe", params, &reply)
	if err != nil {
		return structs.PrinterObjectsQuery{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#query-endstops
func (s *Session) PrinterQueryEndstopsStatus() (structs.PrinterQueryEndstopsStatus, error) {
	var reply structs.PrinterQueryEndstopsStatus

	err := s.Call("printer.query_endstops.status", nil, &reply)
	if err != nil {
		return structs.PrinterQueryEndstopsStatus{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#run-a-gcode-command
func (s *Session) PrinterGcodeScript(params structs.PrinterGcodeScriptParams) (structs.Okay, error) {
	var reply structs.Okay

	err := s.Call("printer.gcode.script", params, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#get-gcode-help
func (s *Session) PrinterGcodeHelp() (structs.PrinterGcodeHelp, error) {
	var reply structs.PrinterGcodeHelp

	err := s.Call("printer.gcode.help", nil, &reply)
	if err != nil {
		return structs.PrinterGcodeHelp{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#start-a-print-job
func (s *Session) PrinterPrintStart(params structs.PrinterPrintStartParams) (structs.Okay, error) {
	var reply structs.Okay

	err := s.Call("printer.print.start", params, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#pause-a-print-job
func (s *Session) PrinterPrintPause() (structs.Okay, error) {
	var reply structs.Okay

	err := s.Call("printer.print.pause", nil, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#resume-a-print-job
func (s *Session) PrinterPrintResume() (structs.Okay, error) {
	var reply structs.Okay

	err := s.Call("printer.print.resume", nil, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/printer/#resume-a-print-job
func (s *Session) PrinterPrintCancel() (structs.Okay, error) {
	var reply structs.Okay

	err := s.Call("printer.print.cancel", nil, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}
