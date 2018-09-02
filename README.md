# AUR Package Status #

A small command-line application to check for AUR package status installed on your machine.
This is only useful for ArchLinux (and derived distros).

### Installation ###

* `yaour -S aur-pkg-status`
* Or use your preferred AUR helper script

### Build from source ###
* Don't trust the pre-built binary? You can build it from the source as following:
  * Clone the repo: `git clone git@github.com:tinhtruong/aur-pkg-status.git`
  * Make sure you have Go install.
  * `cd aur-pkg-status`
  * `./build.sh`
  * The binary is generated at `src/github.com/tinhtruong/aur-pkg-status/aur-pkg-status`

### Usage ###
`aur-pkg-status -status=<status>`

Status can be `updated`, `removed` or `all`. Default to `updated` when status option is not specified

### Sample output ###

`aur-pkg-status -status=all`

| Package Name   	| Installed Version  	| Latest Version  	|
|---	|---	|---	|
| brave-bin  	| 0.19.48-1  	| 0.23.73-1  	|
| pcmciautils  	| 018-7  	| Removed from AUR  	|
| visual-studio-code-bin  	| 1.26.1-2  	| 1.26.1-2  	|

