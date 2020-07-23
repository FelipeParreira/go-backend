package servers

type HostName string

type ServerInfoRequest struct {
	Hostname HostName `json:"hostname" binding:"required" example:"server7"`
}
