package actions

type SystemCertificate struct {
	UsageType           SystemCertificateType
	SubjectName         string
	CommonName          string
	IssuerName          string
	NotBeforeDate       string
	NotAfterDatesString string
	Thumbprint          string
}

type SystemCertificateType struct {
	//TODO:ENUM
	// None, MobiControlRoot, MobiControlServer, MobiControlManager, MobiControlIntermediate, MobiControlProfileSigning, MobiControlCloudLink, MobiControlPrinterAdministrationClient, MobiControlOAuthAuthorizationServer, MobiControlOAuthResourceServer, MobiControlIdp, MobiControlIdpClient, MobiControlSearchService, MobiControlXtHub, MobiControlAPNS, ManagementServerDSE, EnterpriseRoot, EnterpriseServer, EnterpriseManager, EnterpriseProfileSigning, EnterprisePrinterAdministrationClient, TrustedRoot, TrustedServer, TrustedManager
}
