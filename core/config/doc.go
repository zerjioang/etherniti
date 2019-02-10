package config

/*
for production deployments, an SSL certificate information is required.
in order to have a modular an extensible design, this information is provided via env variables as folloeing:

* X_ETHERNITI_SSL_CERT_FILE: /path/to/cert/file.pem
* X_ETHERNITI_SSL_KEY_FILE: /path/to/cert/key.pem
*/
