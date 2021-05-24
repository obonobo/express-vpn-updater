#!/bin/bash

main() (
    case $1 in
        latest) curl_latest                         ;;
        *) echo 'Command' "$1" 'not recognized...'  ;;
    esac
)

curl_latest() {
    curl -i $(sls info | grep -P 'GET.*latest' | cut -d " " -f 5)
}

[[ ${BASH_SOURCE[0]} == $0 ]] && main $@
