from time import sleep, perf_counter
from threading import Thread

def task(id):
    print(f'Running task {id}...')
    sleep(0.5)
    print('Done')


print('Starting new tasks...')
start_time = perf_counter()

for i in range(1,11):
    task(i)

end_time = perf_counter()
whithout_thread = f'It took {end_time- start_time: 0.2f} second(s) to complete all tasks without using threads.'

print('Starting new tasks...')
start_time = perf_counter()

threads = []
for i in range(1, 11):
    t = Thread(target=task, args=(i,))
    threads.append(t)
    t.start()

# wait for the threads to complete
for t in threads:
    t.join()

end_time = perf_counter()
print(whithout_thread)
print(f'It took {end_time- start_time: 0.2f} second(s) to complete all tasks using threads.')
