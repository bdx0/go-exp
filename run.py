#!/usr/bin/env python

import sys
import os

script_path = os.path.dirname(sys.argv[0])

os.environ['GOPATH']=script_path#+  ';' + os.environ['GOPATH']
os.environ['GOBIN']='C:\tools\cygwin\bin'
#os.system('env')
pkg_name = 'dosip'
os.system(' go get ./...')
#os.system('cd src/%s &&  go build ./...' % (pkg_name))
os.system('go build %s && %s.exe && rm %s.exe' % (pkg_name,  pkg_name,  pkg_name))
#os.system('start cmd')
