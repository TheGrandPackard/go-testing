package nats

import (
	"log"
	"time"

	"github.com/nats-io/nats"
	"github.com/thegrandpackard/go-testing/cases"
)

type NATSConnection struct {
	nc *nats.Conn
	c  *cases.Cases
}

type NATSEndpoint interface {
	Publish(subject string, message []byte) error
	Request(subject string, message []byte) ([]byte, error)
	LogRequest(f func(msg *nats.Msg) (response []byte)) func(msg *nats.Msg)
	Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error)
	QueueSubscribe(subject, queueName string, cb nats.MsgHandler) (*nats.Subscription, error)
	Close() (err error)
}

func Init(NATSUrl string, cases *cases.Cases) (natsConnection *NATSConnection, err error) {
	natsConnection = &NATSConnection{
		c: cases,
	}

	natsConnection.nc, err = nats.Connect(NATSUrl, nats.MaxReconnects(-1),
		nats.DisconnectHandler(func(nc *nats.Conn) {
			log.Printf("NATS connection disconnected")
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("NATS reconnected to %v", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Printf("NATS Connection closed. Reason: %s", nc.LastError())
		}))
	if err != nil {
		return
	}

	return
}

func (n *NATSConnection) Publish(subject string, message []byte) error {
	return n.nc.Publish(subject, message)
}

func (n *NATSConnection) Request(subject string, message []byte) (response []byte, err error) {
	var msg *nats.Msg
	msg, err = n.nc.Request(subject, message, 10*time.Second)
	if err != nil {
		return
	}

	response = msg.Data
	return
}

func (n *NATSConnection) LogRequest(f func(msg *nats.Msg, c *cases.Cases) (response []byte)) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		start := time.Now().UnixNano() / int64(time.Millisecond)
		response := f(msg, n.c)
		if len(msg.Reply) > 0 {
			err := n.nc.Publish(msg.Reply, response)
			if err != nil {
				log.Printf("Error publishing message: %s", err.Error())
			}
		}
		log.Printf("%s %d %dms\n", msg.Subject, len(response), ((time.Now().UnixNano() / int64(time.Millisecond)) - start))
	}
}

func (n *NATSConnection) Subscribe(subject string, cb func(msg *nats.Msg, c *cases.Cases) (response []byte)) (sub *nats.Subscription, err error) {
	sub, err = n.nc.Subscribe(subject, n.LogRequest(cb))
	if err != nil {
		return
	}

	log.Printf("Subscribed to subject: %s", subject)
	return
}

func (n *NATSConnection) QueueSubscribe(subject, queue string, cb func(msg *nats.Msg, c *cases.Cases) (response []byte)) (sub *nats.Subscription, err error) {
	sub, err = n.nc.QueueSubscribe(subject, queue, n.LogRequest(cb))
	if err != nil {
		return
	}

	log.Printf("Subscribed to subject: %s on queue: %s", subject, queue)
	return
}

func (n *NATSConnection) Close() (err error) {
	n.nc.Close()
	return
}
