<p align="center">
  <br /><img
    src="internal/assets/sparta-card-readme.png"
    alt="Sparta – Sport and Rehearsal Tracking Application"
  />
</p>

# Sparta

Sparta is a Sport and Rehearsal Tracking Application. It lets the user write down and save all sport activities safely and privately on the computer. No tracking and absolutely zero collection of any user data. Activities are simply saved on the computer without outside interference. Sparta uses top of line `AES-256` encryption, state of the art key derivation using `Argon2` and end-to-end encrypted sharing over the network using `wormhole-william` to keep all your data hidden from any spying eyes.

## Requirements

Sparta is built using the following Go packages:

- [fyne](https://github.com/fyne-io/fyne) (version 1.3.0 or later)
- [wormhole-william](https://github.com/psanford/wormhole-william) (version 1.0.1 or later)
- [x/crypto](https://golang.org/x/crypto)

## Downloads

Please visit the [release page](https://github.com/Jacalz/sparta/releases) for the downloading the latest release.
Versions for Linux, MacOS and Windows are available, with an Android version possibly coming in the future.

Systems that have [Go](https://golang.org) and the [required prequsites for Fyne](https://fyne.io/develop/) installed can alternatively install using `go get`:
```bash
go get github.com/Jacalz/sparta
```

## Contributing

Contributions are strongly appreciated. Everything from creating bug reports to contributing code will help the project a lot, so please feel free to help in any way, shape or form that you feel comfortable doing. We as a community can strive towards making this project even better. If you want to contribute code, the folder structure below will hopefully help you know where to start looking.

### TODO
- [ ] Create a custom widget for displaying exercises in a more pleasant way.
- [ ] Add the option to delete individual exercises.
- [ ] Add import of gpx files from smart training watches and other training equipment.
- [ ] Possibly an Android version and maybe an iOS version if it is requested.

### Folder Structure
- **internal/assets/ :** Logos, artwork and assets bundled in the source code.
- **internal/crypto/ :** Cryptographic functions and wrappers to simplify password hashing and encryption/decryption.
  - **internal/crypto/argon2/ :** Wrapper around `golang.org/x/crypto/argon2` for simplified use inside **internal/crypto**.
  - **internal/crypto/validate/ :** Handles functions for validating usernames and passwords.
- **internal/file/ :** Common code for file and data handling inside the application.
  - **internal/file/parse :** Contains adapted wrappers and functions for parsing numbers and urls from strings.
- **internal/gui :** Graphical interface code for controlling look and function in the application window using `fyne.io/fyne`.
  - **internal/gui/widgets :** Custom widget adaptations to extend and simplify functionality on top of standard widgets.
- **internal/sync :** Implementation of end-to-end encrypted file sharing over a local network using `github.com/psanford/wormhole-william`.
  
## License
- Sparta is licensed under `GNU AFFERO GENERAL PUBLIC LICENSE Version 3` and created by [Jacalz](https://github.com/jacalz).
- All assets are produced by [Jacalz](https://github.com/jacalz) and licensed under the `Creative Commons By-NC-SA 4.0` license.
