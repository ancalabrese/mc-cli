package data

type ManagementServer struct {
	Name                        string
	Status                      ServerStatus
	Fqdn                        string
	PortNumber                  int
	Description                 string
	StatusTime                  string
	MacAddress                  string
	SOTIAssistServerURLOverride string
	TotalConsoleUsers           int
}
