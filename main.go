package main

import (
	"fmt"
	"vpngo-ca/ca"
	"vpngo-ca/csr"
	"vpngo-ca/sign"
)

func main() {
	fmt.Println("🚀 开始 TLS 1.3 证书生成流程...")

	ca.GenerateCA()
	csr.GenerateCSR()
	sign.SignServerCert()

	fmt.Println("🎉 TLS 1.3 证书生成完成！")
}
