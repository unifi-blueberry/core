#!/usr/bin/env bash

set -e

export PACKAGE_NAME="unifi-blueberry-core"
export PACKAGE_VERSION="${PACKAGE_VERSION}"
export PACKAGE_REVISION="${PACKAGE_REVISION:-1}"
export PACKAGE_ARCH="arm64"

if [ -z "$PACKAGE_VERSION" ]; then
  echo "PACKAGE_VERSION not set"
  exit 1
fi

DIR="${PACKAGE_NAME}_${PACKAGE_VERSION}-${PACKAGE_REVISION}_${PACKAGE_ARCH}"

# create staging fir
mkdir $DIR

# copy package files
cp -r packageroot/* $DIR/

# render control
envsubst < packageroot/DEBIAN/control > $DIR/DEBIAN/control

# copy binaries
mkdir -p $DIR/usr/bin
cp bin/* $DIR/usr/bin/

# build
dpkg-deb --build --root-owner-group $DIR

mkdir out
mv ./*.deb out/
