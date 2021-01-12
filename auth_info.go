package mqtt_auth_info

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

type MQTTAuthInfo struct {
	InstanceID string // 服务实例标识
	Host       string // 服务接入点
	Port       int    // 服务接入点端口
	AccessKey  string
	SecretKey  string
	GroupID    string
	ClientID   string
}

func (a *MQTTAuthInfo) GetSignatureInfo() *SignatureInfo {
	clientId := fmt.Sprintf("%s@@@%s", a.GroupID, a.ClientID)
	broker := fmt.Sprintf("tcp://%s:%d", a.Host, a.Port)
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
