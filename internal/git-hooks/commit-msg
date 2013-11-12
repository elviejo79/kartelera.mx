#!/bin/bash

regex="^(refs|fix)\s\#[0-9]+\:"
var=`head -n 1 $1`

function info {
    echo >&2 $1
}

function debug {
    debug=false
    if $debug
    then
        echo >&2 $1
    fi
}

if [[ "$var" =~ $regex ]]
then
    debug "Commit message: OK"
else
    # Define format  message forfirst line in commit message
    info "Commit need to be \"[refs|fix] #ID: message\""
    exit 1
fi
