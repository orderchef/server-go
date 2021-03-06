from escpos import *

import redis
import time

epson = printer.Network("192.168.0.66", 9100)
epson._raw('\x1d@')
epson._raw(constants.CHARCODE_PC437)

r = redis.StrictRedis(host='192.168.18.4', port=6379, db=0)
p = r.pubsub()
p.subscribe('oc_print.local')

while True:
	message = p.get_message()

	if message and message["type"] == "message":
		epson.text(message["data"])
		# epson.cut()

	time.sleep(0.001)
