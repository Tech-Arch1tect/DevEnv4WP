## DevEnv4WP

#### What is it?

A simple tool to provision a local WP dev environment using Docker.

#### Security

This tool focusses on ease-of-use. DO NOT use DevEnv4WP in production, it is local development only and NOT secure. (e.g. the MariaDB root password is `password`). You have been warned!

#### Features

- Automatically provision multiple local WP sites (inculding database, Nginx configuration, etc)
- Custom PHP ini configuration
- Configurable PHP version per-site (7.4, 8.0, 8.1 and 8.2 currently supported)
- Mailpit configured automatically to intercept emails
- PHPMyAdmin configured with automatic root login

#### Requirements

- Modern Docker version (with support for `docker compose`)

#### Stateless

This software holds no state. Everytime the environment is started all provisioning measures are taken (e.g. wp is installed if not already, db created if not already etc.). This works well, however it does not clean up after it's self. e.g. if one.wordpress.local is provisioned, and then is removed from the configuration file, it will be removed from the nginx configuration however the DB will not be removed and neither will the files on-disk. To cleanup & remove or reset the environment, the easiest method is `devenv4wp stop` & `rm -rf data/`.

#### Backstory

My aim with this project was to support all major OS's (Windows through WSL2, Mac including M1/2 and Linux) and create a framework to easily support other architectures & PHP versions when required. Currently all docker images used within this project (third party included) support both AMD64 and ARM64.

#### TODO

- ~~Add configurable bind address~~
- Add apache support (.htaccess support is important in some cases)
- Review code base (this was originally written for necessity / quickly and without much consideration of clean code)
- Add `fix-ownership` command for the rare case ownership may be incorrect.
- Unit tests
- Documentation
- Probably a lot of stuff I can't think of at the moment
