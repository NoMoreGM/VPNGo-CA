package sign

import (
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"math/big"
	"time"

	"vpngo-ca/utils"
)

// SignServerCert 用 Root CA 签发服务器证书
func SignServerCert() error {
	// 读取 Root CA 私钥 & 证书
	caPrivKey, err := utils.LoadPrivateKey("ca.key")
	if err != nil {
		return fmt.Errorf("读取 Root CA 私钥失败: %v", err)
	}

	caCert, err := utils.LoadCertificate("ca.crt")
	if err != nil {
		return fmt.Errorf("读取 Root CA 证书失败: %v", err)
	}

	// 读取服务器 CSR
	csr, err := utils.LoadCSR("server.csr")
	if err != nil {
		return fmt.Errorf("读取 CSR 失败: %v", err)
	}

	// 服务器证书模板
	serverCertTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      csr.Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(3, 0, 0), // 3年有效期
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	// 用 Root CA 签发 server.crt
	serverCertBytes, err := x509.CreateCertificate(rand.Reader, &serverCertTemplate, caCert, csr.PublicKey, caPrivKey)
	if err != nil {
		return fmt.Errorf("生成服务器证书失败: %v", err)
	}

	// 保存 server.crt
	if err := utils.SavePEM("server.crt", "CERTIFICATE", serverCertBytes); err != nil {
		return err
	}

	fmt.Println("✅ 服务器证书签发成功: server.crt")
	return nil
}
