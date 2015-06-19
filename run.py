#!/usr/bin/env python

import sys
import os

script_path = os.path.dirname(sys.argv[0])

os.environ['GOPATH'] = script_path #+  ';' + os.environ['GOPATH']
pkg_name = 'main'
if os.name  == 'nt':
    # os.environ['GOBIN']='C:\tools\cygwin\bin'
    os.system('echo $GOPATH && cd src && pwd && go get .\...')
    os.system('go build %s && bin\%s.exe && rm %s.exe' % (pkg_name,  pkg_name,  pkg_name))
    # os.system('start cmd')
else:
    os.system('echo $GOPATH && cd src/%s && pwd && go get ./...' % (pkg_name))
    os.system('cd bin && ./%s && rm *' % (pkg_name))
