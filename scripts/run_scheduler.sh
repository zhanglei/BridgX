#!/bin/bash
RUN_NAME="gf.bridgx.scheduler"
CURDIR=$(cd $(dirname $0); pwd)
ServiceName=gf.bridgx.scheduler $CURDIR/bin/$RUN_NAME
