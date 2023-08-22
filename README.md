# Gojob: Your Job Openings Management Solution

Welcome to `gojob` – your advanced RESTful API service built with the power of Go (
Golang), Gin, and Gorm.  
Unleash the potential of your job openings management with this robust and cutting-edge
solution.

Explore the **Swagger documentation** and _dive into_ the
details [here](http://swagger.junio.dev:8080/swagger/index.html).

## Table of Contents

- [Table of Contents](#table-of-contents)
	- [Features](#features)
	- [Requirements](#requirements)
	- [Setup](#setup)
	- [Usage](#usage)
	- [Contributing](#contributing)
	- [License](#license)

## Features

Revolutionize your job openings management with `gojob`:

- Effortlessly retrieve organization data from GitHub API.
- Execute a diverse range of operations on the acquired data.
- Obtain comprehensive user lists, complete with their emails and team affiliations.
- Leverage Go's unparalleled capabilities for swift, dependable, and concurrent tasks.

## Requirements

Ensure a seamless setup with the following prerequisites:

- Docker - [Install Guide](https://docs.docker.com/get-docker/)
- Go 1.21+ - [Install Guide](https://golang.org/doc/install)

## Setup

For docker-compose setup, just run:

```sh
docker-compose up
```

**For local setup, follow these steps:**

1. Begin your journey by cloning the repository:

```sh
git clone https://github.com/devjunio/gojob
cd gojob
```

2. Search for the necessary dependencies:

```sh
go mod download
```

3. Establish a solid foundation by crafting a .env file in the project's root directory.
   Duplicate the .env.example file to achieve this:

```sh
cp .env.example .env
```

4. Tailor the environment variables in the .env file to match your distinctive
   requirements.
   Fortify your solution by building it:

```sh
go build -o gojob
```

5. Ignite the engine and set your solution into motion:

```sh
./gojob
```

_Your server will be operational at http://localhost:8080._

## Usage

Unleash the realm of possibilities with gojob! Our platform offers an extensive array of
operations, empowering you to navigate the job landscape effortlessly. With gojob, you
can:

- Retrieve Job Openings: Seamlessly access a diverse range of job opportunities, enabling
  you to explore your career path effectively.
- Execute Rest Operations: Harness the power of Restful API operations on the data you
  obtain, enabling you to manipulate and analyze information with ease.
- Craft Your Unique Experience: Customize your journey with gojob, tailoring your
  interactions to match your unique preferences and aspirations.

Explore, innovate, and redefine your job search experience with gojob today.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for ways to get started.

## License

Distributed under the GPL 3.0 License. See [LICENSE](LICENSE) for more information.
