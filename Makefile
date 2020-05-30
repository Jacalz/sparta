# Constant variables for simplifying code below.
appID=com.github.jacalz.sparta
icon=internal/assets/icon-512.png
name=Sparta

bundle:
	# Bundle the correct logo into sparta/src/bundled/bundled.go
	~/go/bin/fyne bundle -package assets internal/assets/icon-256.png > internal/assets/bundled.go

	# Modify the variable name to be correct.
	sed -i 's/resourceIcon256Png/AppIcon/g' internal/assets/bundled.go

check:
	# Check the whole codebase for misspellings.
	~/go/bin/misspell -w .

	# Run full formating on the code.
	gofmt -s -w .

	# Check the whole program for security issues.
	~/go/bin/gosec ./...

darwin:
	~/go/bin/fyne-cross darwin -arch amd64 -app-id ${appID} -icon ${icon} -output ${name}

linux:
	~/go/bin/fyne-cross linux -arch amd64 -app-id ${appID} -icon ${icon}

windows:
	~/go/bin/fyne-cross windows -arch amd64 -app-id ${appID} -icon ${icon}

cross-compile: darwin linux windows