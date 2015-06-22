#!/usr/bin/env python

import sys
import os

script_path = os.path.dirname(os.path.realpath(sys.argv[0]))

os.environ['GOPATH'] = script_path #+  ';' + os.environ['GOPATH']
pkg_name = 'main'
os.chdir(script_path)
if os.name  == 'nt':
    # os.environ['GOBIN']='C:\tools\cygwin\bin'
    os.system('echo $GOPATH && cd src && pwd && go get .\...')
    os.system('go build %s && bin\%s.exe && rm %s.exe' % (pkg_name,  pkg_name,  pkg_name))
    # os.system('start cmd')
else:
    os.system('echo $GOPATH && cd src/%s && pwd && go get ./...' % (pkg_name))
    os.system('cd bin && pwd && ./%s && rm *' % (pkg_name))
