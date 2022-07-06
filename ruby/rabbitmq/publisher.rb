require 'bunny'

connection = Bunny.new(:host => "localhost", :vhost => "/", :user => "guest", :password => "guest")

def send(channel, queue, message, headers)
  queue = channel.queue(queue)
  exchange = channel.fanout('logging.events')
  queue.bind(exchange)
  queue.publish(
    message,
    :headers  => headers,
    :delivery_mode =>  1
  )
  puts " [x] Sent #{message}"
end

connection.start

queueName = 'to-be-processed'
headers = {
  :timestamp  => Time.now.to_i
}
payload = "Message published with Ruby!"
channel = connection.create_channel
send(channel, queueName, payload, headers)

connection.close