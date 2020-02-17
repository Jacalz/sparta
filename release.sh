#! /bin/bash

# First we make sure to remove and recreate the build directory.
rm -rf build
mkdir build

# Do the initial build for the release.
sudo ~/go/bin/fyne-cross --targets=windows/amd64,darwin/amd64,linux/amd64 --output=sparta

# Make sure to set the correct access rights to the binaries.
(cd build && sudo chmod 666 sparta-linux-amd64 sparta-darwin-amd64 sparta-windows-amd64.exe)

# Prepare the windows executable to be packaged with the icon.
~/go/bin/fyne package -icon assets/icon-512.png -executable build/sparta-windows-amd64.exe -os windows -name Sparta
cp build/fyne.syso fyne.syso

# Package the darwin executable as an application.
~/go/bin/fyne package -os darwin -appID com.github.jacalz.sparta -name Sparta -icon assets/icon-1024.png -executable build/sparta-darwin-amd64
mv Sparta.app build/Sparta.app

# Build the windows binary again to incorporate the logo.
sudo ~/go/bin/fyne-cross --targets=windows/amd64 --output=sparta
rm -rf usr fyne.syso

# Clean up in the build folder (Sparta.app is the one to use).
rm -f build/sparta-darwin-amd64 build/fyne.syso
sudo chmod 666 build/sparta-windows-amd64.exe

# Export the variable to where the ndk is located.
# export ANDROID_NDK_HOME=~/Android/Sdk/ndk/21.0.6113669/

# Build the Android apk using the Android SDK.
# ~/go/bin/fyne package -os android -appID com.github.jacalz.sparta -name Sparta -icon assets/icon-512.png
# mv sparta.apk build/sparta.apk

# Change directory to the build folder.
cd build/

# Lastly we want to compress all the binaries.
tar -cJf sparta-linux-amd64.tar.gz sparta-linux-amd64
zip sparta-darwin-amd64.zip Sparta.app
zip sparta-windows-amd64.zip sparta-windows-amd64.exe

# Final cleanup in the build folder.
rm -rf Sparta.app sparta-linux-amd64 sparta-windows-amd64.exe

