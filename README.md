# password-api

password-api is a simple rest-style backend service that generates passwords.

It accepts the following query params:
* chars (bool) : default "true"
* numbers (bool): default "true"
* specials (bool): default "true"
* length (int): default "64"

It returns a JSON dictionary containing a "password" key.