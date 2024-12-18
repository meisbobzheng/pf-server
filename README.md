# pf-server

First, run `docker compose up --build` to build docker.

Then, run `air` to startup server. Air provides hot reload.

To use, be sure to have direnv installed and aliased in zsh.

For migrations: install golang-migrate

- `make migration [kebab-case-migration-name]`
- `make migrate-up`
- `make migrate-down`
