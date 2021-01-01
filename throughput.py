#!/usr/bin/env python
import sys
import socket
import numpy as np
import math
import time
from time import sleep

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
args = sys.argv
mode = args[1] 
# server side: ./pingpong.py server
# cilent side: ./pingpong.py <server's ip address>

size = 1
maxsize = 1024*1024
limit = math.log(maxsize,2)

if mode == "server": 
	s.bind((socket.gethostname(), 5000))
	s.listen(5)
	print('listening')
	conn, addr = s.accept()
	while True:
		rbuf = conn.recv(maxsize)
	conn.close()
	s.close()
else:
	s.connect((mode, 5000))
	for _ in range(int(limit)+1):
		data = np.array([x for x in range(size)], dtype=str)
		iter = 100000 if size < 512 else(1000 if size <1024*32 else 10)
		start = time.time()
		for _ in range(iter):
			s.send(data)
		elasped_time = time.time() - start
		th = 8 * data.size * iter / elasped_time / 1024 / 1024
		print data.size,'[B],\t', elasped_time,'[s],\t', iter,'[loops],\t', th,'[Mbps],'
		size *= 2
	s.close()

