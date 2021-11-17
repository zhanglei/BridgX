#!/bin/bash
RUN_NAME="gf.bridgx.api"
CURDIR=$(cd $(dirname $0); pwd)
ServiceName=gf.bridgx.api $CURDIR/bin/$RUN_NAME
