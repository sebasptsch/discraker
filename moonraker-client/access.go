package moonrakerclient

import "github.com/sebasptsch/discraker/moonraker-client/structs"

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#login-user
func (s *Session) AccessLogin(params structs.AccessLoginParams) (structs.AccessLogin, error) {
	return rpc[structs.AccessLogin](s, "access.login", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#login-user
func (s *Session) AccessLogout() (structs.AccessLogout, error) {
	return rpc[structs.AccessLogout](s, "access.logout", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#get-current-user
func (s *Session) AccessGetUser() (structs.AccessGetUser, error) {
	return rpc[structs.AccessGetUser](s, "access.get_user", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#create-user
func (s *Session) AccessCreateUser(params structs.AccessCreateUserParams) (structs.AccessLogin, error) {
	return rpc[structs.AccessLogin](s, "access.post_user", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#delete-user
func (s *Session) AccessDeleteUser(params structs.AccessDeleteUserParams) (structs.AccessDeleteUser, error) {
	return rpc[structs.AccessDeleteUser](s, "access.delete_user", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#list-available-users
func (s *Session) AccessUsersList() (structs.AccessUsersList, error) {
	return rpc[structs.AccessUsersList](s, "access.users.list", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#reset-user-password
func (s *Session) AccessUserPassword(params structs.AccessUserPasswordParams) (structs.AccessUserPassword, error) {
	return rpc[structs.AccessUserPassword](s, "access.user.password", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#refresh-json-web-token
func (s *Session) AccessRefreshJWT(params structs.AccessRefreshJWTParams) (structs.AccessRefreshJWT, error) {
	return rpc[structs.AccessRefreshJWT](s, "access.refresh_jwt", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#generate-a-oneshot-token
func (s *Session) AccessOneshotToken() (string, error) {
	return rpc[string](s, "access.oneshot_token", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#get-authorization-module-info
func (s *Session) AccessInfo() (structs.AccessInfo, error) {
	return rpc[structs.AccessInfo](s, "access.info", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#get-the-current-api-key
func (s *Session) AccessGetAPIKey() (string, error) {
	return rpc[string](s, "access.get_api_key", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/authorization/#get-the-current-api-key
func (s *Session) AccessPostAPIKey() (string, error) {
	return rpc[string](s, "access.post_api_key", nil)
}
