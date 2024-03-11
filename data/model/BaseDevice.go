package model

type BaseDevice struct {
	Kind DeviceType
	CompliancePolicyStatus CompliancePolicyStatusType
	ComplianceStatus bool
	ComplianceItems []ComplianceItem
	DeviceId string
	deviceName string
	enrollmentType DeviceEnrollmentType
	enrollmentTime string
	family DeviceFamilyType
	hostName string
	isAgentOnline bool
	macAddress string
	bluetoothMAC string
	wifiMAC string
	mode DeviceMode
	model string
	osVersion string
	path string
	serverName string
	platform PlatformType
	manufacturer string
}

type ComplianceItem struct {
     ComplianceType ComplianceItemType
     ComplianceValue bool
}


type DeviceCustomAttribute struct{
	Name string  
	Value string 
	Type DataType
}

type PlatformType struct{
	//TODO:ENUM
  // Unknown, CESimulator, DOS, WindowsCE, HPCPro, HPC2000, PocketPC, WindowsMobile, Smartphone, WindowsCEDotNet, Windows, Windows95, Windows98, WindowsME, WindowsNT351, WindowsNT40DomainController, WindowsNT40Server, WindowsNT40Workstation, Windows2000DomainController, Windows2000Server, Windows2000Professional, WindowsXPProfessional, WindowsXPHome, Windows2003Server, WindowsVista, WindowsVista64, Windows2008Server, Windows2008Server64, WindowsSeven, WindowsSeven64, Windows2008ServerR2, Windows2008ServerR264, WindowsEight, WindowsEight64, Windows2012Server, Windows2012Server64, WindowsEightOne, WindowsEightOne64, Windows2012ServerR2, Windows2012ServerR264, WindowsTen, WindowsTen64, Windows2016Server, Windows2016Server64, HHPImager, Android, iOS, macOS, AndroidPlus, ScannerPlatformGeneral, WindowsSeven32BitDeviceManager, WindowsSeven64BitDeviceManager, WindowsXPDeviceManager, WEPosDeviceManager, PosReadyDeviceManager, RemHub, RemScanner, RemBiopticScanner, RemMaximal, WindowsPhone, WindowsPhone81, WindowsPhone10, WindowsDesktop10, WindowsPhone10RS1, WindowsDesktop10RS1, WindowsHolographic10, WindowsHolographic10RS1, ZebraPrinter, Linux
}

type DeviceMode struct{
  //Unknown, Active, Disabled, UnenrollPendingUser, UnenrollPendingAdmin, UnenrolledByUser, UnenrolledByAdmin
}

type DeviceType struct{
	//TODO:ENUM
    // AndroidForWork, AndroidElm, AndroidKnox, AndroidPlus, AndroidGeneric, Ios, Mac, WindowsCE, WindowsDesktop, WindowsDesktopLegacy, WindowsPhone, WindowsRuntime, ZebraPrinter, Linux, WindowsHoloLens, WindowsXtHub, Unknown
}

type DataType struct{
	//TODO:ENUM
    // Boolean, Numeric, Text, DateTime, Enumerator, Date
}

type CompliancePolicyStatusType struct{
	//TODO:ENUM
    // Unknown, Pending, NonCompliant, Compliant
}

type ComplianceItemType struct{
	//TODO:ENUM
    // None, IsSecured, IsDeviceAdmin, NotWiped, IsEnabled, IsEnrolled
}

type DeviceEnrollmentType struct{
	//TODO:ENUM
    // NotApplicable, Device, User
}

type DeviceFamilyType struct{
	//TODO:ENUM
    // Unknown, WindowsCE, Apple, WindowsDesktop, AndroidPlus, Scanner, WindowsPhone, Blackberry, Printer, WindowsRuntime, Linux
}
