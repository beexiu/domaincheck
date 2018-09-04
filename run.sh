#!/bin/bash

ROOTDIR=$(cd "$(dirname "$0")"; pwd)
cd $ROOTDIR

sh $ROOTDIR/build.go checkdict

$ROOTDIR/bin/checkdict
