// this file was generated by gomacro command: import _b "crypto/x509"
// DO NOT EDIT! Any change will be lost when the file is re-generated

// +build go1.10

package go1_10

import (
	. "reflect"
	x509 "crypto/x509"
)

// reflection: allow interpreted code to import "crypto/x509"
func init() {
	Packages["crypto/x509"] = Package{
	Binds: map[string]Value{
		"CANotAuthorizedForExtKeyUsage":	ValueOf(x509.CANotAuthorizedForExtKeyUsage),
		"CANotAuthorizedForThisName":	ValueOf(x509.CANotAuthorizedForThisName),
		"CreateCertificate":	ValueOf(x509.CreateCertificate),
		"CreateCertificateRequest":	ValueOf(x509.CreateCertificateRequest),
		"DSA":	ValueOf(x509.DSA),
		"DSAWithSHA1":	ValueOf(x509.DSAWithSHA1),
		"DSAWithSHA256":	ValueOf(x509.DSAWithSHA256),
		"DecryptPEMBlock":	ValueOf(x509.DecryptPEMBlock),
		"ECDSA":	ValueOf(x509.ECDSA),
		"ECDSAWithSHA1":	ValueOf(x509.ECDSAWithSHA1),
		"ECDSAWithSHA256":	ValueOf(x509.ECDSAWithSHA256),
		"ECDSAWithSHA384":	ValueOf(x509.ECDSAWithSHA384),
		"ECDSAWithSHA512":	ValueOf(x509.ECDSAWithSHA512),
		"EncryptPEMBlock":	ValueOf(x509.EncryptPEMBlock),
		"ErrUnsupportedAlgorithm":	ValueOf(&x509.ErrUnsupportedAlgorithm).Elem(),
		"Expired":	ValueOf(x509.Expired),
		"ExtKeyUsageAny":	ValueOf(x509.ExtKeyUsageAny),
		"ExtKeyUsageClientAuth":	ValueOf(x509.ExtKeyUsageClientAuth),
		"ExtKeyUsageCodeSigning":	ValueOf(x509.ExtKeyUsageCodeSigning),
		"ExtKeyUsageEmailProtection":	ValueOf(x509.ExtKeyUsageEmailProtection),
		"ExtKeyUsageIPSECEndSystem":	ValueOf(x509.ExtKeyUsageIPSECEndSystem),
		"ExtKeyUsageIPSECTunnel":	ValueOf(x509.ExtKeyUsageIPSECTunnel),
		"ExtKeyUsageIPSECUser":	ValueOf(x509.ExtKeyUsageIPSECUser),
		"ExtKeyUsageMicrosoftCommercialCodeSigning":	ValueOf(x509.ExtKeyUsageMicrosoftCommercialCodeSigning),
		"ExtKeyUsageMicrosoftKernelCodeSigning":	ValueOf(x509.ExtKeyUsageMicrosoftKernelCodeSigning),
		"ExtKeyUsageMicrosoftServerGatedCrypto":	ValueOf(x509.ExtKeyUsageMicrosoftServerGatedCrypto),
		"ExtKeyUsageNetscapeServerGatedCrypto":	ValueOf(x509.ExtKeyUsageNetscapeServerGatedCrypto),
		"ExtKeyUsageOCSPSigning":	ValueOf(x509.ExtKeyUsageOCSPSigning),
		"ExtKeyUsageServerAuth":	ValueOf(x509.ExtKeyUsageServerAuth),
		"ExtKeyUsageTimeStamping":	ValueOf(x509.ExtKeyUsageTimeStamping),
		"IncompatibleUsage":	ValueOf(x509.IncompatibleUsage),
		"IncorrectPasswordError":	ValueOf(&x509.IncorrectPasswordError).Elem(),
		"IsEncryptedPEMBlock":	ValueOf(x509.IsEncryptedPEMBlock),
		"KeyUsageCRLSign":	ValueOf(x509.KeyUsageCRLSign),
		"KeyUsageCertSign":	ValueOf(x509.KeyUsageCertSign),
		"KeyUsageContentCommitment":	ValueOf(x509.KeyUsageContentCommitment),
		"KeyUsageDataEncipherment":	ValueOf(x509.KeyUsageDataEncipherment),
		"KeyUsageDecipherOnly":	ValueOf(x509.KeyUsageDecipherOnly),
		"KeyUsageDigitalSignature":	ValueOf(x509.KeyUsageDigitalSignature),
		"KeyUsageEncipherOnly":	ValueOf(x509.KeyUsageEncipherOnly),
		"KeyUsageKeyAgreement":	ValueOf(x509.KeyUsageKeyAgreement),
		"KeyUsageKeyEncipherment":	ValueOf(x509.KeyUsageKeyEncipherment),
		"MD2WithRSA":	ValueOf(x509.MD2WithRSA),
		"MD5WithRSA":	ValueOf(x509.MD5WithRSA),
		"MarshalECPrivateKey":	ValueOf(x509.MarshalECPrivateKey),
		"MarshalPKCS1PrivateKey":	ValueOf(x509.MarshalPKCS1PrivateKey),
		"MarshalPKCS1PublicKey":	ValueOf(x509.MarshalPKCS1PublicKey),
		"MarshalPKCS8PrivateKey":	ValueOf(x509.MarshalPKCS8PrivateKey),
		"MarshalPKIXPublicKey":	ValueOf(x509.MarshalPKIXPublicKey),
		"NameConstraintsWithoutSANs":	ValueOf(x509.NameConstraintsWithoutSANs),
		"NameMismatch":	ValueOf(x509.NameMismatch),
		"NewCertPool":	ValueOf(x509.NewCertPool),
		"NotAuthorizedToSign":	ValueOf(x509.NotAuthorizedToSign),
		"PEMCipher3DES":	ValueOf(x509.PEMCipher3DES),
		"PEMCipherAES128":	ValueOf(x509.PEMCipherAES128),
		"PEMCipherAES192":	ValueOf(x509.PEMCipherAES192),
		"PEMCipherAES256":	ValueOf(x509.PEMCipherAES256),
		"PEMCipherDES":	ValueOf(x509.PEMCipherDES),
		"ParseCRL":	ValueOf(x509.ParseCRL),
		"ParseCertificate":	ValueOf(x509.ParseCertificate),
		"ParseCertificateRequest":	ValueOf(x509.ParseCertificateRequest),
		"ParseCertificates":	ValueOf(x509.ParseCertificates),
		"ParseDERCRL":	ValueOf(x509.ParseDERCRL),
		"ParseECPrivateKey":	ValueOf(x509.ParseECPrivateKey),
		"ParsePKCS1PrivateKey":	ValueOf(x509.ParsePKCS1PrivateKey),
		"ParsePKCS1PublicKey":	ValueOf(x509.ParsePKCS1PublicKey),
		"ParsePKCS8PrivateKey":	ValueOf(x509.ParsePKCS8PrivateKey),
		"ParsePKIXPublicKey":	ValueOf(x509.ParsePKIXPublicKey),
		"RSA":	ValueOf(x509.RSA),
		"SHA1WithRSA":	ValueOf(x509.SHA1WithRSA),
		"SHA256WithRSA":	ValueOf(x509.SHA256WithRSA),
		"SHA256WithRSAPSS":	ValueOf(x509.SHA256WithRSAPSS),
		"SHA384WithRSA":	ValueOf(x509.SHA384WithRSA),
		"SHA384WithRSAPSS":	ValueOf(x509.SHA384WithRSAPSS),
		"SHA512WithRSA":	ValueOf(x509.SHA512WithRSA),
		"SHA512WithRSAPSS":	ValueOf(x509.SHA512WithRSAPSS),
		"SystemCertPool":	ValueOf(x509.SystemCertPool),
		"TooManyConstraints":	ValueOf(x509.TooManyConstraints),
		"TooManyIntermediates":	ValueOf(x509.TooManyIntermediates),
		"UnconstrainedName":	ValueOf(x509.UnconstrainedName),
		"UnknownPublicKeyAlgorithm":	ValueOf(x509.UnknownPublicKeyAlgorithm),
		"UnknownSignatureAlgorithm":	ValueOf(x509.UnknownSignatureAlgorithm),
	}, Types: map[string]Type{
		"CertPool":	TypeOf((*x509.CertPool)(nil)).Elem(),
		"Certificate":	TypeOf((*x509.Certificate)(nil)).Elem(),
		"CertificateInvalidError":	TypeOf((*x509.CertificateInvalidError)(nil)).Elem(),
		"CertificateRequest":	TypeOf((*x509.CertificateRequest)(nil)).Elem(),
		"ConstraintViolationError":	TypeOf((*x509.ConstraintViolationError)(nil)).Elem(),
		"ExtKeyUsage":	TypeOf((*x509.ExtKeyUsage)(nil)).Elem(),
		"HostnameError":	TypeOf((*x509.HostnameError)(nil)).Elem(),
		"InsecureAlgorithmError":	TypeOf((*x509.InsecureAlgorithmError)(nil)).Elem(),
		"InvalidReason":	TypeOf((*x509.InvalidReason)(nil)).Elem(),
		"KeyUsage":	TypeOf((*x509.KeyUsage)(nil)).Elem(),
		"PEMCipher":	TypeOf((*x509.PEMCipher)(nil)).Elem(),
		"PublicKeyAlgorithm":	TypeOf((*x509.PublicKeyAlgorithm)(nil)).Elem(),
		"SignatureAlgorithm":	TypeOf((*x509.SignatureAlgorithm)(nil)).Elem(),
		"SystemRootsError":	TypeOf((*x509.SystemRootsError)(nil)).Elem(),
		"UnhandledCriticalExtension":	TypeOf((*x509.UnhandledCriticalExtension)(nil)).Elem(),
		"UnknownAuthorityError":	TypeOf((*x509.UnknownAuthorityError)(nil)).Elem(),
		"VerifyOptions":	TypeOf((*x509.VerifyOptions)(nil)).Elem(),
	}, 
	}
}
