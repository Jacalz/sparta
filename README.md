<p align="center">
  <br /><img
    src="assets/sparta-card-readme.png"
    alt="Sparta – Sport and Rehearsal Tracking Application"
  />
</p>

# Sparta

Sparta is a Sport and Rehearsal Tracking Application. It lets the user write down and save all sport activities safely and privately on the computer. No tracking and absolutely zero collection of any user data. Activities are simply saved on the computer without outside interference. Sparta uses AES-256 encryption to keep all your data hidden from any spying eyes.

## Requirements

Sparta is built using the following Go packages:

- [fyne](https://github.com/fyne-io/fyne) (version 1.2.2 or later)
- [wormhole-william](https://github.com/psanford/wormhole-william) (version 1.0.1 or later)

## Downloads

Please visit the [release page](https://github.com/Jacalz/sparta/releases) for the downloading the latest release.
Versions for Linux, MacOS and Windows are available, with an Android version coming in the future.

## Contributing

Contributions are strongly appreciated. Everything from creating bug reports to contributing code will help the project a lot, so please feel free to help in any way, shape or form that you feel comfortable doing. We as a community can strive towards making this project even better. If you want to contribute code, the folder structure below will hopefully help you know where to start looking.

### Folder structure:
- **assets/ :** Storage for icons and other assets used throughout the project.
- **assets/ :** Logos and artwork along with logos bundled in to the source code.
- **crypto/ :** Cryptographic functions for hashing along with encryption and decryption.
- **file/ :** Common code for file handling in the application.
  - **file/parse :** Contains adapted versions of parse functions from `strconv` for extracting numbers from strings.
- **gui :** All the code for creating the interface along with functions that run on button presses are to be found here.
  - **gui/widgets :** Custom widget adaptations to extend and simplify functionality.
- **sync :** All specific file sharing code for end-to-end encrypted file sharing over a local network.
  
## License
- Sparta is licensed under `GNU AFFERO GENERAL PUBLIC LICENSE Version 3` and created by [Jacalz](https://github.com/jacalz).
- All assets are produced by [Jacalz](https://github.com/jacalz) and licensed under the `Creative Commons By-NC-SA 4.0` license.
