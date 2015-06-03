#!/bin/sh

BRANCH=`git rev-parse --abbrev-ref HEAD`
if test $BRANCH = "master" ; then
  gox -output "bin/{{.OS}}_{{.Arch}}_{{.Dir}}"
  VERSION=`grep Version version.go | cut -c 24- | sed -e 's/"//g'`
  ghr -u DeployGate -r dgate --token $GITHUB_TOKEN --replace $VERSION bin/
fi
