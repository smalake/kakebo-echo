#!/bin/bash

SRC=$1
DIR=$(dirname "$SRC")
FILE=$(basename "$SRC")
PKG=$(basename "$DIR")

# Generate the mock using mockgen
mockgen -source "$SRC" -destination "$DIR/mock_${FILE}" -package "$PKG"

echo "Mock generated at $DIR/mock_${FILE}"