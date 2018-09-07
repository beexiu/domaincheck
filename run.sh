#!/bin/bash

ROOTDIR=$(cd "$(dirname "$0")"; pwd)
cd $ROOTDIR

#sh $ROOTDIR/build.go checkdict

$ROOTDIR/bin/checkdict -tld=com -dic=conf/testds.txt -max=300 -wait=100
