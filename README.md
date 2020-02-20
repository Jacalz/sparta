<p align="center">
  <br /><img
    src="assets/sparta-card-readme.png"
    alt="Sparta – Sport and Rehearsal Tracking Application"
  />
</p>

# Sparta

Sparta is a Sport and Rehearsal Tracking Application. It lets the user write down and save all sport activities safely and privately on the computer. No tracking and no collection of user data, activities are simply saved on the computer without outside interference. Sparta uses AES-256 encryption to keep all your data hidden from any spying eyes.

## Requirements

Sparta is built using the following Go packages:

- [fyne](https://github.com/fyne-io/fyne) (version 1.2.2 or later)
- [wormhole-william](https://github.com/psanford/wormhole-william) (version 1.0.1 or later)

## Downloads

Please visit the [release page](https://github.com/Jacalz/sparta/releases) for the downloading the latest release.
Versions for Linux, MacOS and Windows are available, with an Android version comaing in the future.

## Folder structure:
- **assets/ :** Storage of icons and other assets used in the project.
- **bundled/ :** Images bundled in to the source code.
- **file/ :** Common code for file handling in the application and home of the file package.
  - **file/encrypt :** Cryptographic functions used for encryption etc.
  - **file/parse :** Contains adapted versions of parse functions from `strconv` for extracting numbers from strings.
- **gui :** Contains all the files containing the code used for creating and running the graphical user interface.
- **share :** All specific file sharing code for updaing saved exercises across multiple devices.
  
## License
- Sparta is licensed under `GNU AFFERO GENERAL PUBLIC LICENSE Version 3` and created by [Jacalz](https://github.com/jacalz).
- All assets are produced by [Jacalz](https://github.com/jacalz) and licensed under the `Creative Commons By-NC-SA 4.0` license.
