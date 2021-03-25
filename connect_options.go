package mqtt_auth_info

import mqtt "github.com/eclipse/paho.mqtt.golang"

type ConnectOptions struct {
	Username string
	Password string
	Broker   string
	ClientID string
}

func (ci *ConnectOptions) GetMQTTClientOptions() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(ci.Broker)
	opts.SetClientID(ci.ClientID)
	opts.SetUsername(ci.Username)
	opts.SetPassword(ci.Password)

	return opts
}
