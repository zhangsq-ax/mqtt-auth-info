package mqtt_auth_info

import mqtt "github.com/eclipse/paho.mqtt.golang"

type ConnectOptions struct {
	Username string
	Password string
	Broker   string
	ClientID string
}

func (co *ConnectOptions) GetMQTTClientOptions() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(co.Broker)
	opts.SetClientID(co.ClientID)
	opts.SetUsername(co.Username)
	opts.SetPassword(co.Password)

	return opts
}
