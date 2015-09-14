from escpos import *

try:
	import Image
except ImportError:
	from PIL import Image

import redis
import time

#epson = printer.Network("192.168.0.66", 9100)
#epson._raw('\x1d@')

r = redis.StrictRedis(host='192.168.0.64', port=6379, db=0)
p = r.pubsub()
p.subscribe('oc_print.receipt')

""" Open image file """
im_open = Image.open("taberu-small.jpg")

# Remove the alpha channel on transparent images
if im_open.mode == 'RGBA':
	im_open.load()
	im = Image.new("RGB", im_open.size, (255, 255, 255))
	im.paste(im_open, mask=im_open.split()[3])
else:
	im = im_open.convert("RGB")

pixels   = []
pix_line = ""
im_left  = ""
im_right = ""
switch   = 0
img_size = [ 0, 0 ]

if im.size[0] > 512:
	print  ("WARNING: Image is wider than 512 and could be truncated at print time ")
if im.size[1] > 255:
	raise ImageSizeError()

im_border = printer.Network("192.168.0.66", 9100)._check_image_size(im.size[0])
for i in range(im_border[0]):
	im_left += "0"
for i in range(im_border[1]):
	im_right += "0"

for y in range(im.size[1]):
	img_size[1] += 1
	pix_line += im_left
	img_size[0] += im_border[0]
	for x in range(im.size[0]):
		img_size[0] += 1
		RGB = im.getpixel((x, y))
		im_color = (RGB[0] + RGB[1] + RGB[2])
		im_pattern = "1X0"
		pattern_len = len(im_pattern)
		switch = (switch - 1 ) * (-1)
		for x in range(pattern_len):
			if im_color <= (255 * 3 / pattern_len * (x+1)):
				if im_pattern[x] == "X":
					pix_line += "%d" % switch
				else:
					pix_line += im_pattern[x]
				break
			elif im_color > (255 * 3 / pattern_len * pattern_len) and im_color <= (255 * 3):
				pix_line += im_pattern[-1]
				break
	pix_line += im_right
	img_size[0] += im_border[1]

while True:
	message = p.get_message()

	if message and message["type"] == "message":
		epson = printer.Network("192.168.0.66", 9100)
		epson._raw('\x1d@')
		# print message["data"]
		epson.set('center')
		epson.text(constants.CTL_LF)
		epson.text(constants.CTL_LF)
		epson.text('\n')
		# epson.image("taberu-small.jpg")
		epson._print_image(pix_line, img_size)

		epson.text(message["data"])
		# epson.cut()

	time.sleep(0.01)