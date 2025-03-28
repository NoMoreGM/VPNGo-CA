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

        // 服务器证书模板 / Server Certificate Template
        //
        // ⚠️ 注意 / Note:
        // 必须在 Subject Alternative Name (SAN) 中显式声明服务端标识，否则 TLS 验证将失败：
        // You MUST explicitly specify server identity in the SAN (Subject Alternative Name), or TLS verification will fail:
        //
        // ✅ 若通过 IP 访问，请设置 IPAddresses 字段；
        //    If you connect via IP, set the `IPAddresses` field.
        // ✅ 若通过域名访问，请设置 DNSNames 字段；
        //    If you connect via domain name, set the `DNSNames` field.
        //
        // ❌ 否则将报错，例如：
        //    Otherwise, you’ll see errors like:
        //    `x509: cannot validate certificate for <IP> because it doesn't contain any IP SANs`
        //
        // ✅ 示例 / Examples:
        //     - 通过 IP 连接 / Access by IP:
        //           IPAddresses: []net.IP{net.ParseIP("123.123.123.123")}
        //     - 通过域名连接 / Access by domain:
        //           DNSNames: []string{"vpn.example.com"}
	
        serverCertTemplate := x509.Certificate{
	        SerialNumber: big.NewInt(2),
	        Subject:      csr.Subject,
	        NotBefore:    time.Now(),
	        NotAfter:     time.Now().AddDate(3, 0, 0), // 有效期 3 年 / 3 years valid
	        KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
                ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	        IPAddresses:  []net.IP{
	        	    net.ParseIP("123.123.123.123"), // 🛠️ 替换为你的公网 IP / Replace with your public IP
	        },
	        // DNSNames: []string{"xxx.example.com"}, // 🛠️ 若用域名，请取消注释 / Uncomment if using domain
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
