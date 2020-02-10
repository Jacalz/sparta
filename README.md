<p align="center">
  <br /><img
    src="assets/icon-256.png"
    alt="Sparta – Sport and Rehearsal Tracking Application"
  />
</p>

# Sparta

Sparta is a Sport and Rehearsal Tracking Application for logging all your sport activities safely and privately on the computer. No tracking and no collection of user data, we just save your activities for you. Sparta utilizes Military Grade Encryption to keep all your data and sport activities away from spying eyes.

## Requirements

Sparta depends on the following packages.

- Fyne (version 1.2.2 or later)

## Downloads

Please visit the [release page](https://github.com/Jacalz/sparta/releases) for the downloading the latest release.

## Folder structure:
- **assets/ :** Storage of icons and other assets used in the project.
- **src/ :** The location where all application code is placed and the home of the simple main.go file.
  - **src/bundled :** Images bundled in to the source code.
  - **src/file :** Common code for file handling in the application and home of the file package.
    - **src/file/encrypt :** Cryptographic functions used for encryption etc.
    - **src/file/parse :** Contains adapted versions of parse functions from `strconv` for extracting numbers from strings.
    - **src/file/settings :** The code that handles saving of application settings.
  - **src/gui :** Contains all the files containg the code used for creating and running the graphical user interface.
  
## License
- Sparta is licensed under `GNU AFFERO GENERAL PUBLIC LICENSE Version 3` and created by [Jacalz](https://github.com/jacalz).
- All assets are produced by [Jacalz](https://github.com/jacalz) and licensed under the `Creative Commons By-NC-SA 4.0` license.
