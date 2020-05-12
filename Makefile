# Constant variables for simplifying code below.
appID=com.github.jacalz.sparta
icon=internal/assets/icon-512.png
name=Sparta

bundle:
	# Bundle the correct logo into sparta/src/bundled/bundled.go
	~/go/bin/fyne bundle -package assets assets/icon-256.png > assets/bundled.go

	# Modify the variable name to be correct.
	sed -i 's/resourceIcon256Png/AppIcon/g' assets/bundled.go

check:
	# Check the whole codebase for misspellings.
	~/go/bin/misspell -w .

	# Run full formating on the code.
	gofmt -s -w .

	# Check the whole program for security issues.
	~/go/bin/gosec ./...

compress:
	# Compress the MacOS application into a zip file.
	(cd fyne-cross/dist/ && zip -r sparta-darwin-amd64.zip darwin-amd64/Sparta.app)

	# Compress the Windows binary into a zip file.
	(cd fyne-cross/dist/ && zip sparta-windows-amd64.zip windows-amd64/Sparta.exe)

	# Move out the Linux package and rename it.
	(cd fyne-cross/dist/ && mv linux-amd64/Sparta.tar.gz sparta-linux-amd64.tar.gz)

	# Remove the old folders frm the folder.
	(cd fyne-cross/dist/ && rm -rf darwin-amd64/ windows-amd64/ linux-amd64/)

cross-compile:
	# Remove all dist files.
	rm -rf fyne-cross/dist/*

	# Start by cross compiling for all our targets.
	~/go/bin/fyne-cross -targets=windows/amd64,darwin/amd64,linux/amd64 -icon ${icon} -appID ${appID} -output=${name} .

# Run the full release to prepare for an upcoming release.
release: cross-compile compress