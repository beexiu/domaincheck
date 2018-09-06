#!/bin/bash

ROOTDIR=$(cd "$(dirname "$0")"; pwd)
cd $ROOTDIR

sh $ROOTDIR/build.go checkdict

$ROOTDIR/bin/checkdict -tld=com -dic=conf/testds.txt -max=30 -wait=200
#$ROOTDIR/bin/checkdict -tld=net -dic=conf/custom.txt -max=30 -wait=200
