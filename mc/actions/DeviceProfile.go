package actions

// DeviceProfile rappresents a profile that is associated with a device
type DeviceProfile struct {
	ReferenceId               string
	Name                      string
	VersionNumber             int
	Status                    DeviceProfileStatus
	IsMandatory               bool
	AssignmentDate            string
	Configurations            []DeviceProfileConfiguration
	Packages                  []DevicePackage
	DeviceFamilyQualification DeviceFamilyQualification
}

type DeviceProfileConfiguration struct {
	Name                    string
	DeviceConfigurationType DevicePolicyConfigurationType
	Status                  DevicePolicyConfigurationStatus
}

type DeviceProfileStatus struct {
	//TODO: enum
	//Unknown, InstallPending, InstallFailed, Installed, InstalledPartially, RemovalPending, RemovalFailed, NotInstalled, AdministrativelyRemoved, InstallationDisabled
}

type DevicePolicyConfigurationType struct {
	ConfigurationType DeviceConfigType
	Subtype           DeviceConfigType
}

type DeviceConfigType struct {
	Name string
}

type DevicePackage struct {
	Name        string
	Version     string
	Size        int
	Status      DevicePackageStatus
	ReferenceId string
}

type DevicePackageStatus struct {
	//TODO: enum
	//    Unknown, Installed, PendingInstall, ForceInstall, Downloaded, Uninstalled, PendingUninstall, AnotherVersionAlreadyInstalled, CircularDependency, NotSupported, PrescriptAbort, MissingDependency, UserAborted, InvalidPackage, IncompatiblePlatform, InsufficientFreeSpace, DuplicatedPackage, LowerVersion, FileIO, CreateFile, PackageFileNotFound, CommunicationError, FailedToInstall
}

type DeviceFamilyQualification struct {
	//TODO:enum
	//    Regular, Knox, AndroidWork, Windows10Phone, Windows10Desktop, AppleIOS, AppleMACUser, AppleMACDevice, WindowsHoloLens, AndroidWorkProfileOwner, AndroidCope
}

type DevicePolicyConfigurationStatus struct {
	//TODO:ENUM
	// Unspecified, InstallPending, Delivered, Installed, InstallFailed, InstallFailing, Diagnosing, Disabled, Ignored, UninstallPending, Uninstalled, UninstallFailed
}
