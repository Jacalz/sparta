#! /bin/bash

# Bundle the correct logo into sparta/src/bundled/bundled.go
~/go/bin/fyne bundle -package assets icon-256.png > bundled.go

# Modify the variable name to be correct.
sed -i 's/resourceIcon256Png/AppIcon/g' bundled.go
