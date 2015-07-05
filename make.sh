#!/bin/bash

BINDATA_ARGS="-o util/bindata.go -pkg util config.json"

if [ "$1" == "help" ]; then
	echo "Use <dev> to generate dev bindata. Defaults to production"
	exit
fi

if [ "$1" == "dev" ]; then
	BINDATA_ARGS="-debug ${BINDATA_ARGS}"
	echo "Created util/bindata.go with file proxy"
else
	echo "Created util/bindata.go with all files cached"
fi

go-bindata $BINDATA_ARGS

exit 0