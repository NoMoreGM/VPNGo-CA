# VPNGo-CA: 用 Golang 生成 TLS 1.3 证书体系 / Generate TLS 1.3 Certificate System with Golang

## 🔥 项目介绍 / Project Introduction

`VPNGo-CA` 是一个用 **Golang** 编写的 **TLS 1.3 证书管理系统**，可以一键生成 **Root CA 证书、服务器证书**，并支持 **ECDSA P-384（TLS 1.3 原生支持）**。

本项目适用于：
- 自建 VPN 服务器（如 WireGuard、OpenVPN）
- 内网 HTTPS 服务器（企业内部服务）
- 需要 **自签 TLS 证书** 的应用

---

`VPNGo-CA` is a **TLS 1.3 certificate management system** written in **Golang**, capable of generating **Root CA and server certificates** with **ECDSA P-384** (natively supported by TLS 1.3).  

This project is useful for:
- **Self-hosted VPN servers** (e.g., WireGuard, OpenVPN)
- **Internal HTTPS servers** (enterprise services)
- **Any application requiring self-signed TLS certificates**

---

## 📜 主要功能 / Features

✅ **生成 Root CA** (`ca.crt`, `ca.key`)  
✅ **生成服务器私钥 & CSR** (`server.key`, `server.csr`)  
✅ **用 Root CA 签发服务器证书** (`server.crt`)  
✅ **完全兼容 TLS 1.3**（使用 **ECDSA P-384**）  
✅ **模块化设计，代码清晰，可扩展**  

---

✅ **Generate Root CA** (`ca.crt`, `ca.key`)  
✅ **Generate server private key & CSR** (`server.key`, `server.csr`)  
✅ **Sign server certificate using Root CA** (`server.crt`)  
✅ **Fully compatible with TLS 1.3** (uses **ECDSA P-384**)  
✅ **Modular design, clean and extendable code**  

---
## 🚀 特别鸣谢 / Special Thanks

🎩 **特别鸣谢 [ChatGPT](https://openai.com/chatgpt)，它在整个 TLS 证书系统的设计、代码优化、报错修复、README 编写等方面提供了完全无可匹敌的帮助！**  
🚀 **如果你觉得这个项目有用，也要感谢 ChatGPT！**

## ⚙️ 安装 & 运行 / Installation & Usage

### **1️⃣ 克隆代码 / Clone the repository**
```sh
git clone https://github.com/YOUR_GITHUB_USERNAME/VPNGo-CA.git
