# AUR Package Status #

A small command-line application to check for AUR package status installed on your machine.
This is only useful for ArchLinux (and derived distros).

### Installation ###

* `yaour -S aur-pkg-status`
* Or use your preferred AUR helper script

### Build from source ###
* Don't trust the pre-built binary? You can build it from the source as following:
** Clone the repo: `git clone git@github.com:tinhtruong/aur-pkg-status.git`
** `cd aur-pkg-status`
** `./build.sh`
** The binary is generated at `src/github.com/tinhtruong/aur-pkg-status/aur-pkg-status`

### Sample output ###
    ┌───────────────────────────────┬─────────────────────┬──────────────────┐
    │  Package Name                 │  Installed Version  │  Latest Version  │
    ├───────────────────────────────┼─────────────────────┼──────────────────┤
    │  brave-bin                    │  0.19.48-1          │  0.23.73-1       │
    │  google-cloud-sdk             │  175.0.0-1          │  213.0.0-1       │
    │  insomnia                     │  5.9.6-1            │  6.0.2-1         │
    │  jdk                          │  9.0.1-1            │  10.0.2-1        │
    │  pyenv                        │  1.2.2-1            │  1.2.7-1         │
    │  xcursor-openzone             │  1.2.5-2            │  1.2.6-2         │
    │  xfwm-axiom-theme             │  1-2                │  1-4             │
    │  zuki-themes                  │  3.24.2-1           │  3.26.1-1        │
    └───────────────────────────────┴─────────────────────┴──────────────────┘