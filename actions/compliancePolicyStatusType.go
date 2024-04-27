package actions

type CompliancePolicyStatusType int8

const (
	CompliancePolicyStatusTypeUnknown CompliancePolicyStatusType = iota
	CompliancePolicyStatusTypePending
	CompliancePolicyStatusTypeNonCompliant
	CompliancePolicyStatusTypeCompliant
)

var (
	labelToCompliancePolicyStatusTypeMap = map[string]CompliancePolicyStatusType{
		"Unknown":      CompliancePolicyStatusTypeUnknown,
		"Pending":      CompliancePolicyStatusTypePending,
		"NonCompliant": CompliancePolicyStatusTypeNonCompliant,
		"Compliant":    CompliancePolicyStatusTypeCompliant,
	}
)

func (cpst *CompliancePolicyStatusType) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, cpst, labelToCompliancePolicyStatusTypeMap)
}
