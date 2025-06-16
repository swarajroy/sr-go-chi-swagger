Sample go webserver using demonstrating the inefficiencies whilst loading up swaggeer

- Go Chi Router
- Air for hot reloading

  - Prerequisites

    - install air via https://github.com/air-verse/air using go install github.com/air-verse/air@latest
    - do an air init in the root of the project or tweak the generated .air.toml configuration file with pre_cmd = ["make gen-docs"] as it exists in the commit history
    - install cli binary swag as that is referred to in the Makefile go install github.com/swaggo/swag/cmd/swag@latest
    - follow more for https://github.com/swaggo/http-swagger

  - Whenever you type air at the project root you'll see the server boot up and also the docs getting generated

Try accessing swagger at

- http://localhost:8080/v1/swagger/index.html
- see it fail
