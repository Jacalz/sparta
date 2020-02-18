<p align="center">
  <br /><img
    src="assets/icon-256.png"
    alt="Sparta – Sport and Rehearsal Tracking Application"
  />
</p>

# Sparta

Sparta is a Sport and Rehearsal Tracking Application for logging all your sport activities safely and privately on the computer. No tracking and no collection of user data, we just save your activities for you. Sparta utilizes Military Grade Encryption to keep all your data and sport activities away from spying eyes.

## Requirements

Sparta is built using the following Go packages:

- fyne (version 1.2.2 or later)
- wormhole-william (version 1.0.1 or later)

## Downloads

Please visit the [release page](https://github.com/Jacalz/sparta/releases) for the downloading the latest release.

## Folder structure:
- **assets/ :** Storage of icons and other assets used in the project.
- **bundled/ :** Images bundled in to the source code.
- **file/ :** Common code for file handling in the application and home of the file package.
  - **file/encrypt :** Cryptographic functions used for encryption etc.
  - **file/parse :** Contains adapted versions of parse functions from `strconv` for extracting numbers from strings.
- **gui :** Contains all the files containg the code used for creating and running the graphical user interface.
- **share :** All specific file sharing code for updaing saved exercises across multiple devices.
  
## License
- Sparta is licensed under `GNU AFFERO GENERAL PUBLIC LICENSE Version 3` and created by [Jacalz](https://github.com/jacalz).
- All assets are produced by [Jacalz](https://github.com/jacalz) and licensed under the `Creative Commons By-NC-SA 4.0` license.
