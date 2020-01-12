#! /bin/bash

# Bundle the correct logo into sparta/src/bundled/bundled.go
~/go/bin/fyne bundle -package bundled icon-256.png > ~/go/src/sparta/src/bundled/bundled.go

# Modify the variable name to be correct.
sed -i 's/resourceIcon256Png/AppIcon/g' ~/go/src/sparta/src/bundled/bundled.go