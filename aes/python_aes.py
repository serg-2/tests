#!/usr/bin/python
# -*- coding: utf-8 -*-

from Crypto.Cipher import AES

import base64

secret_pass = AES.new("1234567890123456")
m = 'latitude=16.213123 longtitude=213.412412'

#AES need 16 byte block
while len(m) % 16 != 0: m+=" "

c= base64.b64encode(secret_pass.encrypt(m))

print c
print "================"


d=base64.b64decode(c)

e=secret_pass.decrypt(d)

print e
