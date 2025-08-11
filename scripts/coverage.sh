#!/bin/bash

# checks/install uncover
which uncover &> /dev/null
if [ "$?" -ne "0" ]; then
	set -e
    	go install github.com/gregoryv/uncover/cmd/uncover@v0.9.0
	set +e
fi

# Test and create a .coverage file
PKG_LIST=$(go list ./... | grep -v /vendor/ | xargs) 
go test -covermode=count -coverprofile .coverage $PKG_LIST
if [ "$?" -ne "0" ]; then
	rm -rf .coverage
	exit 1
fi

# Do the coverage
uncover -min 85 .coverage
if [ "$?" -ne "0" ]; then
	rm -rf .coverage
	exit 1
fi

