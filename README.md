<p align="center">
  <br /><img
    width="256"
    src="assets/github-readme-logo.png"
    alt="Sparta – Sport and Rehearsal Tracking Application"
  />
</p>

# Sparta

Sparta, a Sport and Rehearsal Tracking Application with a focus for privacy and security.
The application uses encryption and locally stored files to keep your activities private and hidden from spying eyes.

## Cross compiling using fyne-cross
Run the following command from the root of the repo:
```bash
sudo ../../bin/fyne-cross --targets=windows/amd64,darwin/amd64,linux/amd64 --output=sparta ./src/
```

## Packaging the logo
Run the following command for bundling the icon in a windows binary:
```bash
(cd src && ../../../bin/fyne package -icon ../assets/bundled-logo.png -executable ../sparta-windows-amd64.exe -os windows -name sparta)
```

Run the following command for bundling the icon in a macos binary:
```bash
(cd src && ../../../bin/fyne package -icon ../assets/bundled-logo.png -executable ../sparta-darwin-amd64 -os darwin -name sparta)
```

Run the following command for bundling the icon in a linux binary and installer:
```bash
(cd src && ../../../bin/fyne package -icon ../assets/bundled-logo.png -executable ../sparta-linux-amd64 -os linux -name sparta)
```

## Folder structure:
- assets/ : Storage of icons used in the project.
- src/ : The place where all application code is placed and home of the simple main.go file.
  - src/bundled : Images bundled in to the source code.
  - src/file : Common code for file handling in the application.
    - src/file/encrypt : Cryptographic functions used for encryption etc.
  - src/gui : Contains all code used for and around the graphical user interface.
  
## License
- Sparta is licensed under `GNU AFFERO GENERAL PUBLIC LICENSE Version 3` and created by [Jacalz](https://github.com/jacalz).
- All assets are produced by [Jacalz](https://github.com/jacalz) and licensed under the `Creative Commons By-NC-SA 4.0` license.
