package moonrakerclient

import (
	"github.com/sebasptsch/discraker/moonraker-client/structs"
)

type ConnectionRegisterRemoteMethodParams struct {
	MethodName string `json:"method_name"`
}

func (s *Session) ConnectionRegisterRemoteMethod(params ConnectionRegisterRemoteMethodParams) (structs.Okay, error) {
	return rpc[structs.Okay](s, "connection.register_remote_method", params)
}
