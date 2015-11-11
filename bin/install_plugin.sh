#!/usr/bin/env bash
CURRENTDIR=`pwd`
PLUGIN_PATH="$CURRENTDIR/out/manifest-generator"

$CURRENTDIR/bin/build
cf uninstall-plugin manifest-generator
cf install-plugin "$PLUGIN_PATH" -f