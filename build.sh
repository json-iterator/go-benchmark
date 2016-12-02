#!/bin/bash

# Halt on errors
set -e

export PROJECT_DIR=$(pwd)

function install() {
    reset_env__
    go install github.com/json-iterator/go-benchmark
    # update-naming
}
DEFAULT_TARGET="install"

function dep() {
    # download & compile govendor command
    # remove govendor directory to update govendor to a newer version
    export GOPATH=${PROJECT_DIR}/govendor
    export GOBIN=${PROJECT_DIR}/govendor/bin
    cd ${PROJECT_DIR}
    go install github.com/kardianos/govendor
    local GOVENDOR=${PROJECT_DIR}/govendor/bin/govendor
    reset_env__ # reset back project ${GOPATH}

    # force vendor directory to be created at top level directory
    mkdir vendor >/dev/null 2>&1 || true
    mkdir -p src/github.com >/dev/null 2>&1 || true
    ln -n -s ../../vendor src/github.com/vendor >/dev/null 2>&1 || true
    cd ${PROJECT_DIR}/src/github.com
    $GOVENDOR init

    # now, govendor is initialized properly
    $GOVENDOR "$@"
}

function reset_env__() {
    export GOPATH=${PROJECT_DIR}
    export GOBIN=${PROJECT_DIR}/output
}

if [ -z "$PROJECT_DIR" ]; then
    echo '$PROJECT_DIR is not defined'
    exit 1
fi


## main
TARGET="$1"
if [[ ! -z "$TARGET" ]] ; then
    shift
fi

if [ -z ${DEFAULT_TARGET+x} ]; then
    DEFAULT_TARGET="install"
fi

case $TARGET in
    "help" )
        echo "available targets"
        declare -F | grep -v "__"
        ;;
    * )
        # Halt on errors
        set -e
        # Be verbose
        set -x
        if [[ -z "$TARGET" ]] ; then
            $DEFAULT_TARGET "$@"
        else
            $TARGET "$@"
        fi
        ;;
esac
