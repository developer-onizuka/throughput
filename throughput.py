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
# server side: ./throughput.py server
# cilent side: ./throughput.py <server's ip address>

size = 1
maxsize = 1024*1024
limit = math.log(maxsize,2)

if mode == "server": 
	#s.bind(('127.0.0.1', 5000))
	s.bind((socket.gethostname(), 5000))
	s.listen(5)
	s.setblocking(1)
	print('listening')
	conn, addr = s.accept()
	for _ in range(int(limit)+1):
		iter = 100 if size < 1024*8 else(10 if size <1024*64 else 1)
		#while True:
		for _ in range(iter):
			#rbuf = conn.recv(maxsize)
			rbuf = conn.recv(size)
			#print rbuf
			#if not rbuf: break	
			#conn.sendall(b'ack')
			conn.sendall(rbuf)
			#print size
		print(size,'[B],\t')
		size *= 2
	conn.close()
	s.close()
else:
	s.connect((mode, 5000))
	for _ in range(int(limit)+1):
		data = np.array([1 for x in range(size)], dtype=str)
		iter = 100 if size < 1024*8 else(10 if size <1024*64 else 1)
		start = time.time()
		for _ in range(iter):
			s.sendall(data)
			#tmp = 'nak'
			#print tmp
			rbuf = s.recv(size)
			#print tmp
		elasped_time = (time.time() - start)/2
		th = 8 * data.size * iter / elasped_time / 1024 / 1024
		print(data.size,'[B],\t', elasped_time,'[s],\t', iter,'[loops],\t', th,'[Mbps],')
		size *= 2
	s.close()

