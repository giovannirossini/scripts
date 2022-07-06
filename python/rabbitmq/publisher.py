import pika
from datetime import datetime

# Connection init
credentials = pika.PlainCredentials('guest', 'guest')
parameters = pika.ConnectionParameters('localhost', 5672, '/', credentials)
connection = pika.BlockingConnection(parameters)

def send(channel, queue, message, headers):
  channel.queue_declare(queue=queue)
  channel.basic_publish(
    '',
    queue,
    message,
    pika.BasicProperties(content_type=contentType,
                        delivery_mode=deliveryMode,
                        headers=headers
                        )
  )

  print("Message %s published!" % message)

channel = connection.channel()
queueName='to-be-processed'
contentType='application/octet-stream'
deliveryMode=1 # 1 - Non-Persistent / 2 - Persistent
headers= {
  "timestamp": str(datetime.timestamp(datetime.now()))
}
payload = "Message published with Python!"
send(channel,queueName, payload, headers)
connection.close()