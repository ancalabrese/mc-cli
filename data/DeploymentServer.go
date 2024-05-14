package data

type DeploymentServer struct {
	Name                       string
	Status                     ServerStatus
	IsConnected                bool
	PrimaryManagementAddress   string
	SecondaryManagementAddress string
	PrimaryAgentAddress        string
	SecondaryAgentAddress      string
	PulseTimeout               int
	RuleReload                 int
	ScheduleInter              int
	MinThreads                 int
	MaxThreads                 int
	MaxBurstThreads            int
	PulseWaitInter             int
	ConnectedDeviceCount       int
	ConnectedManagerCount      int
	MsgQueueLength             int
	CurrentThreadCount         int
	DeviceManagementAddress    string
	Certificate                SystemCertificate
}
