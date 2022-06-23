#! /bin/bash

# default env is dev
env=dev 

while test $# -gt 0; do 
    case $1 in 
        -env)
            shift 
            if test $# -gt 0; then 
                env=$1
            fi 
            # shift 
            ;;
        *)
        break 
        ;;
    esac
done 

cd ../../go-pic2text-site 
source .env 
go build -o cmd/pic2text/pic2text cmd/pic2text/main.go 
cmd/pic2text/pic2text -env $env &