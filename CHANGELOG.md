# Changelog

## 0.8.0 - The first stable release
- Fixed wording of "Delete all saved activities" that should have been "Delete all saved exercises".
- Fixed a bug that caused the "Delete all saved exercises" to not work.
- Cleaned up the assets inside `internal/assets`.
  - Dropped icons that were unused for a long time.
- Stopped bundling 256px and 512px icons in the application binaries.
  - This reduces the size of the compiled application.
- Apply a couple fixes and cleanups suggested by staticcheck.
- Cleaned up the Makefile to work better.

## 0.7.1 - Fixes for the fifth beta
- Fix some spelling mistakes.

## 0.7.0 - Fifth beta

- Big refraction of the gui code.
- Reworked and refactored back-end code.
  - Added multi user support.
  - Better security by dropping SHA-256 for Argon2 key derivation.
- A new login screen that works great on mobile and desktop.
- Disable share button when receiving in synchronization.
- Omit json fields that are empty.
  - Should make loading and parsing json a bit faster.
- Use regular expression checks for sync code.
- Update the regular expression checks to be much stricter.
- Improved password and username checks.
- improved error handling and printing errors to the ui.
- Rework the code for the extended entry widget.
- Add an about page with logo and release link.
- Upgrade `wormhole-william` to [v1.0.3](https://github.com/psanford/wormhole-william/releases/tag/v1.0.3).
- Update the Makefile to better work with fyne-cross v2.0.0.
- Use `fyne-cross` [v2.0.0](https://github.com/lucor/fyne-cross/releases/tag/v2.0.0) for building release binaries.
  - This means that binaries are built using Go 1.13.11.
  - Opens up the possibility of Android and OpenBSD releases in the future.
- Upgrade fyne to [v1.3.0](https://github.com/fyne-io/fyne/releases/tag/v1.3.0).
  - We are now using text wrapping where it makes sense.

## 0.6.0 - Forth beta
- Improve the experience when logging in with the wrong account.
- Changed share wording to sync instead.
- Show sync errors with error popups in the GUI.
- Update in-memory login credentials when changing password or username.
- Added support for up/down arrows and enter to execute for username and password change entries in settings.
- How a popup telling the sender if the sync completed.
- Switch from activity wording to exercise.
- Clear the sync code on successful sync.
- In general a better sync experience for users.
- Improved file handling and better error messages (breaking change).
- Fix a couple security issues brought up by gosec.
- Move theme select to beginning to hopefully fix windows weirdness.
- A bunch of spelling fixes.
- Update fyne-cross to v1.4.0.
  - All binaries are now built with Go 1.13.8 over the 1.12.x series.
    - This means that MacOS 10.11 El Capitan is the minimum supported version.
    - Defer uses in the codebase should be around 30% faster.
- Updated fyne to v1.2.3 for a bunch of fixes.
- Switched from shell scripts to makefiles for common operations.
  - Added `make check` command for running `gosec`, `misspell` and `gofmt` for major checks over the whole codebase.

## 0.5.0 - Third beta
- A substantial code refractor of the user interface to cleanup code and remove use of global variables.
- Added an option to change the username in the settings.
- Improve the experience on the first start of the application.
- Better password input checking when changing it inside the settings.
- Switch from AES-256 to AES-512/256 (breaking change).
  - This improves security against length extension attacks.
- Avoid code duplication between share and file operation code.
- Major sort change to sort all exercises after when they occurred instead of when they were added.
- Fix an issue where nothing would be displayed if the exercises file existed but still was empty after several app starts.
- Make sure that the settings reset to light theme and not dark (missed during change of default).
- Switch between login entry widgets using key up and key down.
- Abort and show error if user tries to share with nothing to share.
- Use the correct way to change selected option in the theme switcher.

## 0.4.0 - Second beta
- The password inside the password entry is now cleared from memory after login.
- Internal optimization to reduce draw calls when showing exercises directly on start.
- Improved input checking in password changer to avoid changing password to an invalid password.
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
