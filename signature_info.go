package mqtt_auth_info

import mqtt "github.com/eclipse/paho.mqtt.golang"

type SignatureInfo struct {
	Username string
	Password string
	Broker   string
	ClientID string
}

func (si *SignatureInfo) GetMQTTClientOptions() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(si.Broker)
	opts.SetClientID(si.ClientID)
	opts.SetUsername(si.Username)
	opts.SetPassword(si.Password)

	return opts
}

func (si *SignatureInfo) GenerateMQTTClient(onConnect mqtt.OnConnectHandler, onConnectionLost mqtt.ConnectionLostHandler) mqtt.Client {
	opts := si.GetMQTTClientOptions()
	opts.OnConnect = onConnect
	opts.OnConnectionLost = onConnectionLost
	client := mqtt.NewClient(opts)
	return client
}
