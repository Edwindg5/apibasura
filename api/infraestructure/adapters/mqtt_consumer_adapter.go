// apibasura/api/infraestructure/adapters/mqtt_consumer_adapter.go
package adapters

import (
	"encoding/json"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTConsumerAdapter struct {
	client   mqtt.Client
	topic    string
	callback func(message map[string]interface{})
}

func NewMQTTConsumerAdapter(broker, clientID, topic string) *MQTTConsumerAdapter {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(5 * time.Second)

	return &MQTTConsumerAdapter{
		topic: topic,
		client: mqtt.NewClient(opts),
	}
}

func (m *MQTTConsumerAdapter) Connect(callback func(message map[string]interface{})) error {
	m.callback = callback

	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	if token := m.client.Subscribe(m.topic, 0, m.messageHandler); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	log.Printf("Conectado a MQTT y suscrito al topic: %s", m.topic)
	return nil
}

func (m *MQTTConsumerAdapter) messageHandler(client mqtt.Client, msg mqtt.Message) {
	var message map[string]interface{}
	if err := json.Unmarshal(msg.Payload(), &message); err != nil {
		log.Printf("Error decodificando mensaje JSON: %v", err)
		return
	}

	// Mostrar mensaje en consola
	log.Printf("Mensaje recibido [%s]: %+v", msg.Topic(), message)

	// Ejecutar callback si está definido
	if m.callback != nil {
		m.callback(message)
	}
}

func (m *MQTTConsumerAdapter) Disconnect() {
	m.client.Unsubscribe(m.topic)
	m.client.Disconnect(250)
	log.Println("Desconectado de MQTT")
}