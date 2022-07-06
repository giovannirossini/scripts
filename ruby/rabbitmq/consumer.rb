require 'bunny'

connection = Bunny.new(:host => "localhost", :vhost => "/", :user => "guest", :password => "guest")

def callback(channel, queue)
  queue = channel.queue(queue)
  queue.subscribe(:auto_ack => true) do |delivery_info, properties, payload|
    puts "Received #{payload}, message properties are #{properties.inspect}"
  end
end

connection.start
channel = connection.create_channel
queueName = 'to-be-processed'
callback(channel, queueName)
connection.close