#!/bin/bash

function get_latest_release() {
    curl --silent "https://api.github.com/repos/$1/releases/latest" | grep -Po '"tag_name": "\K.*?(?=")'
}

function push_release() {
    echo "Create tag $1 locally"
    git tag -a $1 -m "Version $1"
    echo "Push tag $1 to origin"
    git push origin $1        
    echo "Create release $1 in Github"
    goreleaser
}

cmd="status"
if [ "$#" -ge 1 ]; then
        cmd=$1
fi

case $cmd in
        status)
        latest_release=$( get_latest_release $2 )
        echo "latest release $latest_release"
        ;;
        push)
        echo "Pushing release $2"
        push_release $2
        ;;
esac