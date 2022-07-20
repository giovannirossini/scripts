import pika


class RabbitMQ:
  def connect(self, host, port, username, password):
    credentials = pika.PlainCredentials(username, password)
    parameters = pika.ConnectionParameters(host, port, '/', credentials)
    try:
        conn = pika.BlockingConnection(parameters)
    except Exception as error:
        print(error)
    return conn

  def create_channel(self, connection):
    self.connection = connection
    try:
        channel = self.connection.channel()
    except Exception as error:
        print(error)
    return channel

  def consume(self, channel, queueName, callback):
    self.channel = channel
    self.queueName = queueName
    self.callback = callback
    self.channel.basic_consume(
        queue=self.queueName, on_message_callback=self.callback, auto_ack=True
    )
    try:
        self.channel.start_consuming()
    except Exception as error:
        print(error)

  def publish(self, channel, queue, message, headers = None, contentType = None, deliveryMode = None):
    self.channel = channel
    self.queue = queue
    self.message = message
    self.headers = headers
    if contentType == None:
        self.contentType = 'application/octet-stream'
    else:
        self.contentType = contentType
    if deliveryMode == None:
        self.deliveryMode = 1 # 1 - Non-Persistent / 2 - Persistent
    else:
        self.deliveryMode = deliveryMode
    try:
        self.channel.basic_publish(
            '',
            self.queue,
            self.message,
            pika.BasicProperties(content_type=self.contentType,
                delivery_mode=self.deliveryMode,
                headers=self.headers
            )
        )
    except Exception as error:
        print(error)

    return print(f"Message {self.message} published!")
