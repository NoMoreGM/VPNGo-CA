package ca

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"time"

	"vpngo-ca/utils"
)

// GenerateCA 生成 Root CA 证书
func GenerateCA() error {
	// 生成 Root CA 私钥
	caPrivKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return fmt.Errorf("生成 Root CA 私钥失败: %v", err)
	}

	// 编码私钥
	privKeyBytes, err := x509.MarshalECPrivateKey(caPrivKey)
	if err != nil {
		return fmt.Errorf("无法编码 Root CA 私钥: %v", err)
	}

	// 保存 ca.key
	if err := utils.SavePEM("ca.key", "EC PRIVATE KEY", privKeyBytes); err != nil {
		return err
	}

	// Root CA 证书模板
	caTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Organization:       []string{"MyVPN Root CA"},
			OrganizationalUnit: []string{"CA Authority"},
			CommonName:         "MyVPN Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // 10年有效期
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
	}

	// 自签 Root CA 证书
	caCertBytes, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return fmt.Errorf("生成 Root CA 证书失败: %v", err)
	}

	// 保存 ca.crt
	if err := utils.SavePEM("ca.crt", "CERTIFICATE", caCertBytes); err != nil {
		return err
	}

	fmt.Println("✅ Root CA 证书生成成功: ca.crt")
	return nil
}
