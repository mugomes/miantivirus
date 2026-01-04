# MiAntivirus

MiAntivirus is a graphical interface for ClamAV that allows you to scan your computer for viruses and easily update the virus definition database through a user-friendly interface.

## Screenshots

<img height="300" alt="image" src="https://github.com/user-attachments/assets/7b2299fa-7515-4be8-9b00-604ed03b73da" />
<img height="300" alt="image" src="https://github.com/user-attachments/assets/9e362326-73e4-4441-8b90-15bfd789fce7" />

## Contribution

1. Report Bugs or Suggestions: https://github.com/mugomes/miantivirus/issues

2. Send translations for more languages: Download the .po file, translate it and submit it via Pull Request

3. If you want, give MiAntivirus a "star".

4. Support the project financially through GitHub Sponsors or other forms of support.

## Support

- https://github.com/sponsors/mugomes
- https://mugomes.github.io/apoie.html

## Official MiAntivirus link

- https://github.com/mugomes/miantivirus

### Official Author link

- https://mugomes.github.io

### Links to Third-Party Resources Used

- https://go.dev
- https://fyne.io
- https://www.clamav.net
- https://github.com/polkit-org/polkit/

## Pronunciation

MiAntivirus is a Portuguese name.

The "Mi" part comes from the Portuguese word "Meu" (meaning "My").

Its correct pronunciation is:

"Mee-ahn-chee-VEE-roos"

- Mi: sounds like "mee"
- An: nasal sound, like "ahn"
- ti: pronounced "chee" in Portuguese
- virus: pronounced "VEE-roos", not "vai-rus"

You can read it smoothly as: Mee-ahn-chee-VEE-roos

## Installation

### DEB

Double-click the deb package and click install, or run the command in the terminal:

```bash
sudo dpkg -i miantivirus*.deb
sudo apt install -f
```

### Integrity

To verify integrity, copy the hash provided next to the downloaded release and use [MiCheckHash](https://github.com/mugomes/micheckhash/releases) (graphical interface) to verify its integrity, or use the terminal.

Example:

```bash
echo "12f95d1ba9b46f5713d8010963c4c782e315b7985027c44c6e292ede69454301 miantivirus_2.0.0_all.deb" | sha256sum -c
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
