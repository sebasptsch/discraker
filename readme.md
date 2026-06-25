# Discraker

A Moonraker-Discord Bridge built in Go.

Has a much smaller total size than any alternatives built in Python or JavaScript, suited for applications where memory or disk space is limited.

This is still very much a work in progress and barely has any commands or features implemented but the structure, updating mechanism and everything else is all present.

## Config

Discraker uses a TOML based config stored alongside the rest of the klipper config files in the default location. By default a skeleton config will be created with critical values missing.