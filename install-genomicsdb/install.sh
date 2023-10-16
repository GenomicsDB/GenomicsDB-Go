#!/bin/bash

#
# The MIT License
#
# Copyright (c) 2023 dātma, inc™
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.
#

GENOMICSDB_BRANCH=${GENOMICSDB_BRANCH:-develop}

CMAKE_INSTALL_PREFIX=${GENOMICSDB_INSTALL_DIR:-/usr/local}
GENOMICSDB_DIR=$(mktemp -d).GenomicsDB
if [[ $CMAKE_INSTALL_PREFIX == "/usr/local" ]]; then
  SUDO="sudo"
else
  SUDO=""
fi
GENOMICSDB_NO_CLEAN=${GENOMICSDB_NO_CLEAN:false}
GENOMICSDB_DEVELOPER=${GENOMICSDB_DEVELOPER:false}
if [[ $GENOMICSDB_DEVELOPER == 1 || $GENOMICSDB_DEVELOPER == true ]]; then
  GENOMICSDB_DEVELOPER=true
else
  GENOMICSDB_DEVELOPER=false
fi

# Get absolute path for CMAKE_INSTALL_PREFIX
CMAKE_INSTALL_PREFIX=$(python3 -c "import os,sys; print(os.path.abspath(sys.argv[1]))" $CMAKE_INSTALL_PREFIX)

cleanup() {
  if [[ $1 -eq 1 ]]; then
    if [[ $GENOMICSDB_NO_CLEAN == false ]]; then
      echo "*** Error encountered. Removing $GENOMICSDB_DIR..."
    else
      echo "*** Error encountered building workspace at $GENOMICSDB_DIR "
    fi
  fi

  if [[ $GENOMICSDB_NO_CLEAN == false ]]; then
    rm -fr $GENOMICSDB_DIR
    echo "Removing $GENOMICSDB_DIR DONE"
  fi
  exit $1
}

echo "Removing old genomicsdb artifacts from $CMAKE_INSTALL_PREFIX..."
$SUDO rm -vf $CMAKE_INSTALL_PREFIX/lib/lib*genomicsdb*
$SUDO rm -vf $CMAKE_INSTALL_PREFIX/lib/pkgconfig/genomicsdb.pc
$SUDO rm -vfr $CMAKE_INSTALL_PREFIX/include/genomicsdb*
echo "Removing old genomicsdb artifacts DONE"

echo "Installing GenomicsDB from branch $GENOMICSDB_BRANCH into $CMAKE_INSTALL_PREFIX..."
git clone https://github.com/GenomicsDB/GenomicsDB.git -b $GENOMICSDB_BRANCH $GENOMICSDB_DIR

pushd $GENOMICSDB_DIR

echo "Installing prerequisites on System=$(uname)..."
if [[ $(uname) == "Darwin" ]]; then
  PREREQS_ENV=$GENOMICSDB_DIR/prereqs.sh scripts/prereqs/install_prereqs.sh
else
  PREREQS_ENV=$GENOMICSDB_DIR/prereqs.sh $SUDO scripts/prereqs/install_prereqs.sh
fi
if [[ -f $GENOMICSDB_DIR/prereqs.sh ]]; then
  echo "Sourcing $GENOMICSDB_DIR/prereqs.sh"
  source $GENOMICSDB_DIR/prereqs.sh
fi
echo "Install prerequisites DONE"

mkdir build
pushd build

if [[ -n $OPENSSL_ROOT_DIR ]]; then
  CMAKE_PREFIX_PATH="-DCMAKE_PREFIX_PATH=$OPENSSL_ROOT_DIR"
fi
if [[ $GENOMICSDB_DEVELOPER == true ]]; then
  BUILD_FOR_GO="-DBUILD_FOR_GO=1"
fi
cmake $BUILD_FOR_GO $CMAKE_PREFIX_PATH -DCMAKE_INSTALL_PREFIX=$CMAKE_INSTALL_PREFIX -DAWSSDK_ROOT_DIR=$GENOMICSDB_DIR/awssdk -DGCSSDK_ROOT_DIR=$GENOMICSDB_DIR/gcssdk -DPROTOBUF_ROOT_DIR=$GENOMICSDB_DIR/protobuf .. || cleanup 1
make -j4 || cleanup 1
$SUDO make install
popd
popd

PARENT_DIR=$(dirname $(pwd))
TOPLEVEL_GIT_DIR=$(git rev-parse --show-toplevel)
if [[ $GENOMICSDB_DEVELOPER == true && $TOPLEVEL_GIT_DIR == $PARENT_DIR ]]; then
  # This script is from the git repository and we probably want to install protobuf go sources
  echo "Copying GenomicsDB protobuf go generated sources..."
  cp -vf $CMAKE_INSTALL_PREFIX/genomicsdb/protobuf/go/* $PARENT_DIR/bindings/protobuf
  echo "Copying GenomicsDB protobuf go generated sources DONE"
fi
$SUDO rm -fr $CMAKE_INSTALL_PREFIX/genomicsdb

if [[ $(uname) == "Darwin" ]]; then
  echo "export DYLD_LIBRARY_PATH=$CMAKE_INSTALL_PREFIX/lib:$DYLD_LIBRARY_PATH" > genomicsdb.env
else
  echo "export LD_LIBRARY_PATH=$CMAKE_INSTALL_PREFIX/lib:$LD_LIBRARY_PATH" > genomicsdb.env
fi

if [[ -f $CMAKE_INSTALL_PREFIX/lib/pkgconfig/genomicsdb.pc ]]; then
  echo "export PKG_CONFIG_PATH=$CMAKE_INSTALL_PREFIX/lib/pkgconfig" >> genomicsdb.env
elif [[ -f $CMAKE_INSTALL_PREFIX/lib64/pkgconfig/genomicsdb.pc ]]; then
  echo "export PKG_CONFIG_PATH=$CMAKE_INSTALL_PREFIX/lib64/pkgconfig" >> genomicsdb.env
fi

echo "Installing GenomicsDB DONE"
cleanup 0

