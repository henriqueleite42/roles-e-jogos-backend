# Roles e Jogos Backend

## How to use this repository

### Context

- Pay attention to the port, the default value used on the `.env.example` file is `3001`

### Installation

- Install [GNU Make](https://www.gnu.org/software/make/)
- Install [Golang](https://go.dev/doc/install)
- Install [Sqlc](https://docs.sqlc.dev/en/stable/overview/install.html)
- Install [AtlasGo](https://atlasgo.io)
- Install [Docker (with docker compose)](https://docs.docker.com/engine/install/)
	- [Docker post install](https://docs.docker.com/engine/install/linux-postinstall/)
- Install [libvips](https://www.libvips.org/install.html)
- Copy `.env.example`, rename it to `.env.docker`, then fill the values

### Database

- On Terminal A, run `make run-db`
- On Terminal B, run `make migrate`
- Close all tabs

### Run

- On Terminal A, run `make es` (`es` from **E**co**S**ystem)
- On Terminal B, run `make run`
- If you ever need to make a change on the API, stop the process on Terminal B and run it again.
