package actions

type PlatformType int8

const (
	Unknown PlatformType = iota + 1
	CESimulator
	DOS
	WindowsCE
	HPCPro
	HPC2000
	PocketPC
	WindowsMobile
	Smartphone
	WindowsCEDotNet
	Windows
	Windows95
	Windows98
	WindowsME
	WindowsNT351
	WindowsNT40DomainController
	WindowsNT40Server
	WindowsNT40Workstation
	Windows2000DomainController
	Windows2000Server
	Windows2000Professional
	WindowsXPProfessional
	WindowsXPHome
	Windows2003Server
	WindowsVista
	WindowsVista64
	Windows2008Server
	Windows2008Server64
	WindowsSeven
	WindowsSeven64
	Windows2008ServerR2
	Windows2008ServerR264
	WindowsEight
	WindowsEight64
	Windows2012Server
	Windows2012Server64
	WindowsEightOne
	WindowsEightOne64
	Windows2012ServerR2
	Windows2012ServerR264
	WindowsTen
	WindowsTen64
	Windows2016Server
	Windows2016Server64
	HHPImager
	Android
	iOS
	macOS
	AndroidPlus
	ScannerPlatformGeneral
	WindowsSeven32BitDeviceManager
	WindowsSeven64BitDeviceManager
	WindowsXPDeviceManager
	WEPosDeviceManager
	PosReadyDeviceManager
	RemHub
	RemScanner
	RemBiopticScanner
	RemMaximal
	WindowsPhone
	WindowsPhone81
	WindowsPhone10
	WindowsDesktop10
	WindowsPhone10RS1
	WindowsDesktop10RS1
	WindowsHolographic10
	WindowsHolographic10RS1
	ZebraPrinter
	Linux
)

var (
	PlatformTypeKey = map[uint8]string{
		1:  "Unkown",
		2:  "CESimulator",
		3:  "DOS",
		4:  "WindowsCE",
		5:  "HPCPro",
		6:  "HPC2000",
		7:  "PocketPC",
		8:  "WindowsMobile",
		9:  "Smartphone",
		10: "WindowsCEDotNet",
		11: "Windows",
		12: "Windows95",
		13: "Windows98",
		14: "WindowsME",
		15: "WindowsNT351",
		16: "WindowsNT40DomainController",
		17: "WindowsNT40Server",
		18: "WindowsNT40Workstation",
		19: "Windows2000DomainController",
		20: "Windows2000Server",
		21: "Windows2000Professional",
		22: "WindowsXPProfessional",
		23: "WindowsXPHome",
		24: "Windows2003Server",
		25: "WindowsVista",
		26: "WindowsVista64",
		27: "Windows2008Server",
		28: "Windows2008Server64",
		29: "WindowsSeven",
		30: "WindowsSeven64",
		31: "Windows2008ServerR2",
		32: "Windows2008ServerR264",
		33: "WindowsEight",
		34: "WindowsEight64",
		35: "Windows2012Server",
		36: "Windows2012Server64",
		37: "WindowsEightOne",
		38: "WindowsEightOne64",
		39: "Windows2012ServerR2",
		40: "Windows2012ServerR264",
		41: "WindowsTen",
		42: "WindowsTen64",
		43: "Windows2016Server",
		44: "Windows2016Server64",
		45: "HHPImager",
		46: "Android",
		47: "iOS",
		48: "macOS",
		49: "AndroidPlus",
		50: "ScannerPlatformGeneral",
		51: "WindowsSeven32BitDeviceManager",
		52: "WindowsSeven64BitDeviceManager",
		53: "WindowsXPDeviceManager",
		54: "WEPosDeviceManager",
		55: "PosReadyDeviceManager",
		56: "RemHub",
		57: "RemScanner",
		58: "RemBiopticScanner",
		59: "RemMaximal",
		60: "WindowsPhone",
		61: "WindowsPhone81",
		62: "WindowsPhone10",
		63: "WindowsDesktop10",
		64: "WindowsPhone10RS1",
		65: "WindowsDesktop10RS1",
		66: "WindowsHolographic10",
		67: "WindowsHolographic10RS1",
		68: "ZebraPrinter",
		69: "Linux",
	}
)
