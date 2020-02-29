# Constant variables for simplifying code below.
appID=com.github.jacalz.sparta
icon=assets/icon-512.png
name=Sparta

android:
	# Export the variable to specify where the ndk is located.
	export ANDROID_NDK_HOME=~/Android/Sdk/ndk/21.0.6113669/

	# Build the Android apk using the Android SDK.
	~/go/bin/fyne package -os android -appID ${appID} -name ${name} -icon ${icon}

check:
	# Check the whole codebase for misspellings.
	~/go/bin/misspell -w .

	# Run full formating on the code.
	gofmt -s -w .

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
	sudo rm -rf fyne-cross/dist/*

	# Start by cross compiling for all our targets.
	sudo ~/go/bin/fyne-cross -targets=windows/amd64,darwin/amd64,linux/amd64 -icon ${icon} -appID ${appID} -output=${name} .

fix-permissions:
	# Docker has to be run with sudo and thus the fyne-cross directory has the wrong file permissions.
	sudo chmod -R 777 fyne-cross/

# Run the full release to prepare for an upcoming release.
release: cross-compile fix-permissions compress