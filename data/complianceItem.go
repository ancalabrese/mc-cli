package data

type ComplianceItemType int8

type ComplianceItem struct {
	ComplianceType  ComplianceItemType `json:"ComplianceType"`
	ComplianceValue bool               `json:"ComplianceValue"`
}

const (
	ComplianceItemTypeNone ComplianceItemType = iota
	ComplianceItemTypeIsSecured
	ComplianceItemTypeIsDeviceAdmin
	ComplianceItemTypeNotWiped
	ComplianceItemTypeIsEnabled
	ComplianceItemTypeIsEnrolled
	UnknownComplianceItemType
)

var (
	labelToComplianceItemMap = map[string]ComplianceItemType{
		"None":          ComplianceItemTypeNone,
		"IsSecured":     ComplianceItemTypeIsSecured,
		"IsDeviceAdmin": ComplianceItemTypeIsDeviceAdmin,
		"NotWiped":      ComplianceItemTypeNotWiped,
		"IsEnabled":     ComplianceItemTypeIsEnabled,
		"IsEnrolled":    ComplianceItemTypeIsEnrolled,
		"Unknown":       UnknownComplianceItemType,
	}
	complianceItemMapToLabel = map[ComplianceItemType]string{
		ComplianceItemTypeNone:          "None",
		ComplianceItemTypeIsSecured:     "IsSecured",
		ComplianceItemTypeIsDeviceAdmin: "IsDeviceAdmin",
		ComplianceItemTypeNotWiped:      "NotWiped",
		ComplianceItemTypeIsEnabled:     "IsEnabled",
		ComplianceItemTypeIsEnrolled:    "IsEnrolled",
		UnknownComplianceItemType:       "Unknown",
	}
)

func (cit *ComplianceItemType) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, cit, labelToComplianceItemMap, UnknownComplianceItemType)
}

func (cit ComplianceItemType) String() string {
	return complianceItemMapToLabel[cit]
}
