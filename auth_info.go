package mqtt_auth_info

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

type ConnectionProtocol string

const (
	ConnectionProtocol_TCP ConnectionProtocol = "tcp" // 端口 1883
	ConnectionProtocol_SSL ConnectionProtocol = "ssl" // 端口 8883
	ConnectionProtocol_WS  ConnectionProtocol = "ws"  // 端口 80
	ConnectionProtocol_WSS ConnectionProtocol = "wss" // 端口 443
)

type MQTTAuthInfo struct {
	Protocol   ConnectionProtocol // 连接协议
	InstanceID string             // 服务实例标识
	Host       string             // 服务接入点
	Port       int                // 服务接入点端口
	AccessKey  string
	SecretKey  string
	GroupID    string
	ClientID   string
}

func (a *MQTTAuthInfo) GetSignatureInfo() *SignatureInfo {
	port := a.Port
	if port == 0 {
		switch a.Protocol {
		case ConnectionProtocol_TCP:
			port = 1883
		case ConnectionProtocol_SSL:
			port = 8883
		case ConnectionProtocol_WS:
			port = 80
		case ConnectionProtocol_WSS:
			port = 443
		}
	}
	clientId := fmt.Sprintf("%s@@@%s", a.GroupID, a.ClientID)
	broker := fmt.Sprintf("%s://%s:%d", a.Protocol, a.Host, port)
	username := fmt.Sprintf("Signature|%s|%s", a.AccessKey, a.InstanceID)
	mac := hmac.New(sha1.New, []byte(a.SecretKey))
	mac.Write([]byte(clientId))
	password := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return &SignatureInfo{
		Username: username,
		Password: password,
		Broker:   broker,
		ClientID: clientId,
	}
}
