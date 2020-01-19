# Changelog

## 0.3.0 - First beta
- Major UI rework for better design visuals and more options by using tabs.
- Added a settings page to give the user more control over the application.
 - Added support for changing between ligth and dark themes.
 - Added support for changing the password in settings.
 - Added support for deleteing all activity data.
- Added a button to cancel the form input and clear all data inside the form.
- Clean the text in the form on submit.
- Packaged icons are generlla in higher definition.
- Delete the password value from memory after it has been used.
- Added support for showing error and preventing login if entered password is incorrect.
- Big code cleanups and lots of code refactoring. 

## 0.2.0 - Forth alpha
- Insecure passwords now print an error when the length is lower than 8 characters.
- The Fyne toolkit was updated from 1.2.0 to in development version of 1.2.2 for bug fixes to enable new features.
- Login on the press of the return key (possible thanks to bug fixes in Fyne 1.2.2).
- Updates to gui sizes, layouts and error messages.
- Enable the use of all input fields in the new activity view.
- Binaries are now compiled with `fyne-cross` 1.3.2 with Go 1.12.14 instead of 1.12.12 which brings fixes for the `runtime` package along with fixes for MacOS binaries being rejected by Gatekeeper.

## 0.1.2 - Third alpha
- Fixed the need to restart the app to see new activities if program is opened for the first time.
- Add script (release.sh) to automate building and packaging of release binaries. This fixes the packaging of windows binaries with logos and makes releases much faster to do.
