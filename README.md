## sensupluginssensu

## Commands
 * checkSensuSilences

## Usage

### checkSensuSilences
This will check to see if any checks have been silenced without an expiration. If
it finds any they will be sent as the check output and it will return a `critical` condition.
If no checks are found than an `ok` condition will be returned and no output will be
produced per standard unix convention.

Ex. ` sensupluginssensu checksensu`

## Installation

1. godep go build -o bin/sensupluginssensu
1. chmod +x sensupluginsensu (*nix only)
1. cp sensupluginssensu /usr/local/bin

## Notes
