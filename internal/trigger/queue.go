package trigger

import (
    "log"
    "github.com/rabbitmq/amqp091-go"
)

func StartQueueTrigger(queueName, rabbitMQURL string, handler func(msg string)) error {
    conn, err := amqp091.Dial(rabbitMQURL)
    if err != nil {
        return err
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
    if err != nil {
        return err
    }

    log.Printf("Listening to queue: %s", queueName)
    for msg := range msgs {
        handler(string(msg.Body))
    }

    return nil
}
