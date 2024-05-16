package data

type PlatformType int8

const (
	UnknownPlatform PlatformType = iota + 1
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
	platformTypeToLabelMap = map[PlatformType]string{
		UnknownPlatform:                "Unkown",
		CESimulator:                    "CESimulator",
		DOS:                            "DOS",
		WindowsCE:                      "WindowsCE",
		HPCPro:                         "HPCPro",
		HPC2000:                        "HPC2000",
		PocketPC:                       "PocketPC",
		WindowsMobile:                  "WindowsMobile",
		Smartphone:                     "Smartphone",
		WindowsCEDotNet:                "WindowsCEDotNet",
		Windows:                        "Windows",
		Windows95:                      "Windows95",
		Windows98:                      "Windows98",
		WindowsME:                      "WindowsME",
		WindowsNT351:                   "WindowsNT351",
		WindowsNT40DomainController:    "WindowsNT40DomainController",
		WindowsNT40Server:              "WindowsNT40Server",
		WindowsNT40Workstation:         "WindowsNT40Workstation",
		Windows2000DomainController:    "Windows2000DomainController",
		Windows2000Server:              "Windows2000Server",
		Windows2000Professional:        "Windows2000Professional",
		WindowsXPProfessional:          "WindowsXPProfessional",
		WindowsXPHome:                  "WindowsXPHome",
		Windows2003Server:              "Windows2003Server",
		WindowsVista:                   "WindowsVista",
		WindowsVista64:                 "WindowsVista64",
		Windows2008Server:              "Windows2008Server",
		Windows2008Server64:            "Windows2008Server64",
		WindowsSeven:                   "WindowsSeven",
		WindowsSeven64:                 "WindowsSeven64",
		Windows2008ServerR2:            "Windows2008ServerR2",
		Windows2008ServerR264:          "Windows2008ServerR264",
		WindowsEight:                   "WindowsEight",
		WindowsEight64:                 "WindowsEight64",
		Windows2012Server:              "Windows2012Server",
		Windows2012Server64:            "Windows2012Server64",
		WindowsEightOne:                "WindowsEightOne",
		WindowsEightOne64:              "WindowsEightOne64",
		Windows2012ServerR2:            "Windows2012ServerR2",
		Windows2012ServerR264:          "Windows2012ServerR264",
		WindowsTen:                     "WindowsTen",
		WindowsTen64:                   "WindowsTen64",
		Windows2016Server:              "Windows2016Server",
		Windows2016Server64:            "Windows2016Server64",
		HHPImager:                      "HHPImager",
		Android:                        "Android",
		iOS:                            "iOS",
		macOS:                          "macOS",
		AndroidPlus:                    "AndroidPlus",
		ScannerPlatformGeneral:         "ScannerPlatformGeneral",
		WindowsSeven32BitDeviceManager: "WindowsSeven32BitDeviceManager",
		WindowsSeven64BitDeviceManager: "WindowsSeven64BitDeviceManager",
		WindowsXPDeviceManager:         "WindowsXPDeviceManager",
		WEPosDeviceManager:             "WEPosDeviceManager",
		PosReadyDeviceManager:          "PosReadyDeviceManager",
		RemHub:                         "RemHub",
		RemScanner:                     "RemScanner",
		RemBiopticScanner:              "RemBiopticScanner",
		RemMaximal:                     "RemMaximal",
		WindowsPhone:                   "WindowsPhone",
		WindowsPhone81:                 "WindowsPhone81",
		WindowsPhone10:                 "WindowsPhone10",
		WindowsDesktop10:               "WindowsDesktop10",
		WindowsPhone10RS1:              "WindowsPhone10RS1",
		WindowsDesktop10RS1:            "WindowsDesktop10RS1",
		WindowsHolographic10:           "WindowsHolographic10",
		WindowsHolographic10RS1:        "WindowsHolographic10RS1",
		ZebraPrinter:                   "ZebraPrinter",
		Linux:                          "Linux",
	}

	platformLabelToTypeMap = map[string]PlatformType{
		"Unkown":                         UnknownPlatform,
		"CESimulator":                    CESimulator,
		"DOS":                            DOS,
		"WindowsCE":                      WindowsCE,
		"HPCPro":                         HPCPro,
		"HPC2000":                        HPC2000,
		"PocketPC":                       PocketPC,
		"WindowsMobile":                  WindowsMobile,
		"Smartphone":                     Smartphone,
		"WindowsCEDotNet":                WindowsCEDotNet,
		"Windows":                        Windows,
		"Windows95":                      Windows95,
		"Windows98":                      Windows98,
		"WindowsME":                      WindowsME,
		"WindowsNT351":                   WindowsNT351,
		"WindowsNT40DomainController":    WindowsNT40DomainController,
		"WindowsNT40Server":              WindowsNT40Server,
		"WindowsNT40Workstation":         WindowsNT40Workstation,
		"Windows2000DomainController":    Windows2000DomainController,
		"Windows2000Server":              Windows2000Server,
		"Windows2000Professional":        Windows2000Professional,
		"WindowsXPProfessional":          WindowsXPProfessional,
		"WindowsXPHome":                  WindowsXPHome,
		"Windows2003Server":              Windows2003Server,
		"WindowsVista":                   WindowsVista,
		"WindowsVista64":                 WindowsVista64,
		"Windows2008Server":              Windows2008Server,
		"Windows2008Server64":            Windows2008Server64,
		"WindowsSeven":                   WindowsSeven,
		"WindowsSeven64":                 WindowsSeven64,
		"Windows2008ServerR2":            Windows2008ServerR2,
		"Windows2008ServerR264":          Windows2008ServerR264,
		"WindowsEight":                   WindowsEight,
		"WindowsEight64":                 WindowsEight64,
		"Windows2012Server":              Windows2012Server,
		"Windows2012Server64":            Windows2012Server64,
		"WindowsEightOne":                WindowsEightOne,
		"WindowsEightOne64":              WindowsEightOne64,
		"Windows2012ServerR2":            Windows2012ServerR2,
		"Windows2012ServerR264":          Windows2012ServerR264,
		"WindowsTen":                     WindowsTen,
		"WindowsTen64":                   WindowsTen64,
		"Windows2016Server":              Windows2016Server,
		"Windows2016Server64":            Windows2016Server64,
		"HHPImager":                      HHPImager,
		"Android":                        Android,
		"iOS":                            iOS,
		"macOS":                          macOS,
		"AndroidPlus":                    AndroidPlus,
		"ScannerPlatformGeneral":         ScannerPlatformGeneral,
		"WindowsSeven32BitDeviceManager": WindowsSeven32BitDeviceManager,
		"WindowsSeven64BitDeviceManager": WindowsSeven64BitDeviceManager,
		"WindowsXPDeviceManager":         WindowsXPDeviceManager,
		"WEPosDeviceManager":             WEPosDeviceManager,
		"PosReadyDeviceManager":          PosReadyDeviceManager,
		"RemHub":                         RemHub,
		"RemScanner":                     RemScanner,
		"RemBiopticScanner":              RemBiopticScanner,
		"RemMaximal":                     RemMaximal,
		"WindowsPhone":                   WindowsPhone,
		"WindowsPhone81":                 WindowsPhone81,
		"WindowsPhone10":                 WindowsPhone10,
		"WindowsDesktop10":               WindowsDesktop10,
		"WindowsPhone10RS1":              WindowsPhone10RS1,
		"WindowsDesktop10RS1":            WindowsDesktop10RS1,
		"WindowsHolographic10":           WindowsHolographic10,
		"WindowsHolographic10RS1":        WindowsHolographic10RS1,
		"ZebraPrinter":                   ZebraPrinter,
		"Linux":                          Linux,
	}
)

func (pt *PlatformType) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, pt, platformLabelToTypeMap, UnknownPlatform)
}

func (pt PlatformType) String() string {
	return platformTypeToLabelMap[pt]
}
