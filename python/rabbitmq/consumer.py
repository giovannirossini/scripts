from modules.rabbitmq.connector import RabbitMQ

def callback(self, channel, method, properties, body):
    print(f"Received {body}, message properties are {properties}")

amqp = RabbitMQ()
connection = amqp.connect('localhost', 5672, 'guest', 'guest')
channel =  amqp.create_channel(connection)

queue='some-queue'

amqp.consume(channel, queue, callback)