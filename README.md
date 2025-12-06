# MiAntivirus

MiAntivirus is a graphical interface for ClamAV that allows you to scan your computer for viruses and easily update the virus definition database through a user-friendly interface.

## Contribution

1. Report Bugs or Suggestions:

https://github.com/mugomes/miantivirus/issues

2. Send translations for more languages:

Download the .po file, translate it
Submit it via Pull Request

3. If you want, give MiAntivirus a "star".

4. Support the project financially through GitHub Sponsors or other forms of support.

## Support

- https://github.com/sponsors/mugomes
- https://www.mugomes.com.br/apoie.html

## Official MiAntivirus link

- https://github.com/mugomes/miantivirus

### Links to Third-Party Resources Used

- https://gambas.sourceforge.net
- https://www.clamav.net
- https://docs.gtk.org/gtk3/
- https://github.com/polkit-org/polkit/

## Installation

### DEB

Double-click the deb package and click install, or run the command in the terminal:

```bash
sudo dpkg -i miantivirus*.deb
sudo apt install -f
```

### AppImage

A portable version will be released soon.

### Source

**Warning:** Don't do this. It's much easier to just double click a .deb or .appimage.

Download the source code and run it:

```bash
/usr/bin/gbc3 -e -a -g -t miantivirus/
gba3 miantivirus/

mv miantivirus/miantivirus.gambas miantivirus/miantivirus

chmod +x miantivirus/miantivirus

./miantivirus/miantivirus
```

### Integrity

To verify integrity, copy the hash provided next to the downloaded release and use [MiCheckHash](https://github.com/mugomes/micheckhash/releases) (graphical interface) to verify its integrity, or use the terminal.

Example:

```bash
echo "sha256:7e4f667108f9ab58ccb2419269d4f533851f756c69f325d6d039808c8710e5b5 miantivirus_1.0.16-0ubuntu1_all.deb" | sha256sum -c
```

If "Success" is displayed, the file was downloaded correctly.

## Usage

MiAntivirus allows you to scan multiple files and folders; just click "Add" and choose which ones you want to add. You can add both types.

Before scanning, I always recommend updating the ClamAV database. To do this, click "Update Database" in the Tools menu. Once the update is complete, a "Completed" message will be displayed.

In Options, in the Tools menu, you can enable and disable scanning features and add and remove folders and files that you want to ignore during the scan.

When you click "Scan," ClamAV will load, this process may take a while (this slowness in loading the database is due to ClamAV itself), after loading, the selected folders and files will be checked. Once the analysis is complete, you can remove infected files or folders if any viruses are detected.

**Warning:** ClamAV may present some false positives, so consider whether you really want to delete them. This deletion is permanent and cannot be recovered. Therefore, be careful when deleting a file or folder.

To check for new MiAntivirus updates, click Check for Updates in the About menu.

## Limitations/Bugs

The software may contain some limitations or bugs, so it is very important to use the official contact channels to report bugs.

## License

The MiAntivirus is provided under:

[SPDX-License-Identifier: GPL-2.0-only](https://spdx.org/licenses/GPL-2.0-only.html)

Beign under the terms of the GNU General Public License version 2 only.

All contributions to the MiAntivirus are subject to this license.
