import pika

credentials = pika.PlainCredentials('guest', 'guest')
parameters = pika.ConnectionParameters('localhost', 5672, '/', credentials)
connection = pika.BlockingConnection(parameters)
queueName='to-be-processed'

def callback(channel, method, properties, body):
  print("Received %r, message properties are %r" % (body, properties))


channel = connection.channel()
channel.queue_declare(queue=queueName)
channel.basic_consume(
  queue=queueName, on_message_callback=callback, auto_ack=True
)
channel.start_consuming()