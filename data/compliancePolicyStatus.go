package data

type CompliancePolicyStatus int8

const (
	CompliancePolicyStatusTypeUnknown CompliancePolicyStatus = iota
	CompliancePolicyStatusTypePending
	CompliancePolicyStatusTypeNonCompliant
	CompliancePolicyStatusTypeCompliant
)

var (
	labelToCompliancePolicyStatusTypeMap = map[string]CompliancePolicyStatus{
		"Unknown":      CompliancePolicyStatusTypeUnknown,
		"Pending":      CompliancePolicyStatusTypePending,
		"NonCompliant": CompliancePolicyStatusTypeNonCompliant,
		"Compliant":    CompliancePolicyStatusTypeCompliant,
	}

	compliancePolicyStatusToLabelMap = map[CompliancePolicyStatus]string{
		CompliancePolicyStatusTypeUnknown:      "Unknown",
		CompliancePolicyStatusTypePending:      "Pending",
		CompliancePolicyStatusTypeNonCompliant: "NonCompliant",
		CompliancePolicyStatusTypeCompliant:    "Compliant",
	}
)

func (cpst *CompliancePolicyStatus) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, cpst, labelToCompliancePolicyStatusTypeMap, CompliancePolicyStatusTypeUnknown)
}

func (cpst CompliancePolicyStatus) String() string {
	return compliancePolicyStatusToLabelMap[cpst]
}
