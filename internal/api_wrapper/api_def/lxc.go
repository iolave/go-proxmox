package apidef

type GetLXCIPResponse struct {
	IP string `json:"ip" example:"10.10.10.10"`
}

type PostLXCExecRequest struct {
	CMD   string `json:"cmd" example:"ls -l /unknowndir"`
	Shell string `json:"shell" example:"bash"`
}

type PostLXCExecResponse struct {
	Output   string `json:"output" example:"ls: cannot access '/unknowndir': No such file or directory"`
	ExitCode int    `json:"exitCode" example:"2"`
}

type PostLXCExecAsyncResponse struct {
	ID string `json:"id"`
}
