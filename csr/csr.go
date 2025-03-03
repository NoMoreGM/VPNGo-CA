package csr

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"

	"vpngo-ca/utils"
)

// GenerateCSR 生成服务器 CSR
func GenerateCSR() error {
	// 生成 ECDSA 私钥
	privKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return fmt.Errorf("生成服务器私钥失败: %v", err)
	}

	// 编码私钥
	privKeyBytes, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		return fmt.Errorf("无法编码服务器私钥: %v", err)
	}

	// 保存 server.key
	if err := utils.SavePEM("server.key", "EC PRIVATE KEY", privKeyBytes); err != nil {
		return err
	}

	// CSR 模板
	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			Country:      []string{"CN"},
			Organization: []string{"MyVPN Corp"},
			CommonName:   "vpn.example.com",
		},
		DNSNames:           []string{"vpn.example.com"},
		SignatureAlgorithm: x509.ECDSAWithSHA384,
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privKey)
	if err != nil {
		return fmt.Errorf("生成 CSR 失败: %v", err)
	}

	// 保存 server.csr
	if err := utils.SavePEM("server.csr", "CERTIFICATE REQUEST", csrBytes); err != nil {
		return err
	}

	fmt.Println("✅ 服务器 CSR 生成成功: server.csr")
	return nil
}
