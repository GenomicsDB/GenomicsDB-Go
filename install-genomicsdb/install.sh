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

GENOMICSDB_BRANCH=${GENOMICSDB_BRANCH:-ng_go}
CMAKE_INSTALL_PREFIX=${GENOMICSDB_INSTALL_DIR:-/usr/local}
GENOMICSDB_DIR=$(mktemp -d).GenomicsDB
if [[ $CMAKE_INSTALL_PREFIX == "/usr/local" ]]; then
  SUDO="sudo"
else
  SUDO=""
fi
GENOMICSDB_NO_CLEAN=${GENOMICSDB_NO_CLEAN:false}

# Get absolute path for CMAKE_INSTALL_PREFIX
CMAKE_INSTALL_PREFIX=$(python3 -c "import os,sys; print(os.path.abspath(sys.argv[1]))" $CMAKE_INSTALL_PREFIX)

cleanup() {
  if [[ $1 -eq 1 ]]; then
    if [[ $GENOMICSDB_NO_CLEAN == true ]]; then
      echo "*** Error encountered building workspace at $GENOMICSDB_DIR "
    else
      echo "*** Error encountered. Removing $GENOMICSDB_DIR..."
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

echo "Installing GenomicsDB into $CMAKE_INSTALL_PREFIX..."
git clone https://github.com/GenomicsDB/GenomicsDB.git -b $GENOMICSDB_BRANCH $GENOMICSDB_DIR

pushd $GENOMICSDB_DIR

#TEMP FIX
sed -i.bak -e '160d' CMakeLists.txt

mkdir build
pushd build
cmake -DBUILD_FOR_GO=1 -DCMAKE_INSTALL_PREFIX=$CMAKE_INSTALL_PREFIX .. || cleanup 1
make -j4 || cleanup 1
$SUDO make install
popd
popd

PARENT_DIR=$(dirname $(pwd))
echo "PARENT_DIR=$PARENT_DIR"
TOPLEVEL_GIT_DIR=$(git rev-parse --show-toplevel)
echo "TOPLEVEL_GIT_DIR=$TOPLEVEL_GIT_DIR"
if [[ $TOPLEVEL_GIT_DIR == $PARENT_DIR ]]; then
  # This script is from the git repository and we probably want to install protobuf go sources
  echo "Copying GenomicsDB protobuf go generated sources..."
  cp -vf $CMAKE_INSTALL_PREFIX/genomicsdb/protobuf/go/* $PARENT_DIR/bindings/protobuf
  echo "Copying GenomicsDB protobuf go generated sources DONE"
fi
$SUDO rm -fr $CMAKE_INSTALL_PREFIX/genomicsdb

if [[ $(uname) == "darwin" ]]; then
  echo "export DYLD_LIBRARY_PATH=$CMAKE_INSTALL_PREFIX/lib" > genomicsdb.env
else
  echo "export LD_LIBRARY_PATH=$CMAKE_INSTALL_PREFIX/lib" > genomicsdb.env
fi
echo "export PKG_CONFIG_PATH=$CMAKE_INSTALL_PREFIX/lib/pkgconfig" >> genomicsdb.env

echo "Installing GenomicsDB DONE"
cleanup 0

