# Changelog

## 0.4.0 - Second beta
- The password inside the password entry is now cleared from memory after login.
- Internal optimization to reduce draw calls when showing exercises directly on start.
- Improved imput checking in password changer to avoid changing password to an invalid password.
- Fixed a lot of potential security issues reported by [gosec](https://github.com/securego/gosec).
- Settings page is now separated in to groups for better clarity.
- Make the settings scrollable on smaller screens.
- Added option to reset all settings to their default values.
- Added strict input checking using regular expressions for the add activity view.
- Store reps and sets as unsigned ints (possible due to regex input checking).
- Use internal fyne preferences api for settings storage (breaking change).
- Move exercises into the fyne config path (breaking change).
- Switch from xml to json files for exercises storage (breaking change).
  - Json is about 3x faster than xml and should show slight speed improvements.
- Added end-to-end encrypted local network exercise sharing support between devices (using [wormhole-william](https://github.com/psanford/wormhole-william)).
- Changed default theme to light.
- Upgraded to a newer version of golang.org/x/net.
- Upgraded fyne to v1.2.2, which brings the following changes for Sparta:
  - Horizontal scrolling when holding down the shift key.
  - Pasting unicode characters could cause a panic.
  - Shortcuts are now handled by the event queue - fixed possible deadlock.
  - When auto scaling check the monitor in the middle of the window, not top left.
  - Update scale on Linux to be "auto" by default (and numbers are relative to 96DPI standard).
  - Scale calculations are now relative to system scale - the default "1" matches the system.
- Package MacOS releases as `.zip` files instead of `tar.gz`.


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
