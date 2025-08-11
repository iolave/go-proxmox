#!/bin/bash

# set opts
PREVIEW=false
while getopts ":p" opt; do
  case $opt in
    p)
      PREVIEW=true
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done

# set a variable to know this script path
SCRIPT_PATH="$( cd "$(dirname "$0")" ; pwd -P )"
ROOT="$(realpath $SCRIPT_PATH/..)"

################################################################################
## Build api-wrapper docs
################################################################################
which swag &> /dev/null
if [ "$?" -ne "0" ]; then
	set -e
	echo "info: installing swag"
    	go install github.com/swaggo/swag/cmd/swag@v1.16.6
	set +e
fi
swag init \
	-d $ROOT \
	-g ./cmd/pve_api_wrapper/pve_api_wrapper.go \
	--exclude ./ \
	--parseInternal \
	-o $ROOT/docs/api-wrapper \

################################################################################
## Build go-client docs and generate site
################################################################################
# Check if $TMPDIR is set
if [ -z "$TMPDIR" ]; then
    echo "warn: TMPDIR is not set, using /tmp"
    TMPDIR="/tmp"
fi

# check if TMPDIR is exists
if [ ! -d "$TMPDIR" ]; then
    echo "error: TMPDIR '$TMPDIR' does not exist"
    exit 1
fi

VENV="${TMPDIR}/venv/go-proxmox"
mkdir -p $VENV

python3 -m venv ${VENV}
source "${VENV}/bin/activate"

PIP_DEPS=" \
	mkdocs-material \
	pymdown-extensions \
	markdown-callouts \
	mkdocs-render-swagger-plugin \
"

# iterate over the list of dependencies
for dep in $PIP_DEPS; do
	pip3 show $dep &> /dev/null
	if [ "$?" -ne "0" ]; then
		echo "info: installing $dep"
		pip3 install $dep
	fi
done

rm -rf $ROOT/docs/go-client/pkg
mkdir -p $ROOT/docs/go-client/pkg
go run $ROOT/cmd/gomarkdoc/main.go
mkdocs build -f $ROOT/mkdocs.yml

################################################################################
## Preview docs
################################################################################
if [ "$PREVIEW" == "true" ]; then
	mkdocs serve -f $ROOT/mkdocs.yml
fi
