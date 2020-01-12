# Changelog

## 0.2.0 - Forth alpha
- Insecure passwords now print an error when the length is lower than 8 characters.
- The Fyne toolkit was updated from 1.2.0 to the in development version of 1.2.2 for bug fixes needed to enable new features in Sparta.
- Login on the press of the return key (possible thanks to bug fixes for extended widggets in Fyne 1.2.2).
- Overall updates to gui sizes, layouts and error messages.
- Enable the use of all input fields in the new activity view.
- Binaries are now compiled with fyne-cross 1.3.2 with golang 1.12.14 instead of 1.12.12 which brings fixes for the runtime package along with fixes for MacOS binaries being rejected by Gatekeeper.

## 0.1.2 - Third alpha
- Fixed the need to restart the app to see new activities if program is opened for the first time.
- Add script (release.sh) to automate building and packaging of release binaries. This fixes the packaging of windows binaries with logos and makes releases much faster to do.
