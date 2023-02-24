#!/usr/bin/env bash

# This is needed to look for libraries in ./lib on compile and
#  at  <executable location>/lib on runtime.
export CGO_LDFLAGS="-L./lib -Wl,-rpath,\$ORIGIN/lib"