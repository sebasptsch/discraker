# Discraker

A Moonraker-Discord Bridge built in Go.

Has a much smaller total size than any alternatives built in Python or JavaScript, suited for applications where memory or disk space is limited.

This is still very much a work in progress and barely has any commands or features implemented but the structure, updating mechanism and everything else is all present.

## Install

The easiest way to install and setup Discraker is to use the included script to create the required `release_info.json` and the systemd service. This script can be found in the `./scripts` folder. This only has to be run once and once you add the reccomended section to the Moonraker updater config updates will be managed for you.

## Config

Discraker uses a TOML based config stored alongside the rest of the klipper config files in the default location. By default a skeleton config will be created with critical values missing.