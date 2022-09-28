from datetime import datetime
from modules.rabbitmq.connector import RabbitMQ

amqp = RabbitMQ()
connection = amqp.connect('localhost', 5672, 'guest', 'guest')
channel =  amqp.create_channel(connection)

queue='some-queue'
headers= {
  "timestamp": str(datetime.timestamp(datetime.now()))
}

payload = "{'name': 'John Doe', 'age': '26'}"

amqp.publish(channel,queue, payload, headers)

connection.close()