package rabbit

import (
	"github.com/streadway/amqp"
)

type RabbitBuilder interface {
	Connect() *Rabbit
	CreateChannel() *Rabbit
	SetPath(qn string) *Rabbit
	Handler() *Rabbit
	StartListen() error
}

// main struct
// TODO bindings, QoS
type Rabbit struct {
	// connection string
	cs string
	// queue name
	qN string
	// consumer
	c string
	// auto ask
	aAsk bool
	// exclusive
	exclusive bool
	// no local
	nLocal bool
	// no wait
	nWait bool
	// options
	args amqp.Table

	/////
	messages <-chan amqp.Delivery

	doneCh chan struct{}

	conn    *amqp.Connection
	channel *amqp.Channel
	q       amqp.Queue
	errs    []error
}

func NewRabbitServer(cs string /*qN string, consumer string, autoAsk bool, exclusive bool, noLocal bool, noWait bool, options amqp.Table*/) *Rabbit {
	if cs == "" {
		cs = "amqp://guest:guest@localhost:5672/"
	}
	return &Rabbit{
		cs: cs,
		/*		qN:        qN,
				c:         consumer,
				aAsk:      autoAsk,
				exclusive: exclusive,
				nLocal:    noLocal,
				nWait:     noWait,
				args:      options,*/
		errs: make([]error, 2),
	}
}

func (r *Rabbit) Connect() *Rabbit {
	var err error
	r.conn, err = amqp.Dial(r.cs)
	if err != nil {
		panic(err)
	}
	return r
}

func (r *Rabbit) CreateChannel() *Rabbit {
	var err error
	r.channel, err = r.conn.Channel()
	if err != nil {
		panic(err)
	}
	return r
}

// router
func (r *Rabbit) Router(inQ string, outQ string) *Rabbit {
	q, err := r.channel.QueueDeclare(
		"hash",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	r.messages, err = r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	return r
}

func (r *Rabbit) StartListen(listen Serve) {
	for {
		select {
		case m := <-r.messages:
			go func(msg amqp.Delivery) {
				listen.ServeRabbit(msg)
			}(m)
		case <-r.doneCh:
			return
		}
	}
}
