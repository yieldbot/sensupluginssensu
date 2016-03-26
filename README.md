## Sensu Plugins Sensu

### checkSensuSilences

This will check to see if any checks have been silenced without an expiration. If
it finds any they will be sent as the check output and it will return a `critical` condition.
If no checks are found than an `ok` condition will be returned and no output will be
produced per standard unix convention.
