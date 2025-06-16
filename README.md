Sample go webserver using demonstrating the inefficiencies whilst loading up swaggeer
- Go Chi Router
- Air for hot reloading 
  - Prerequisites
    - install air via https://github.com/air-verse/air using go install github.com/air-verse/air@latest
    - tweak the generated .air.toml configuration file with pre_cmd = ["make gen-docs"]

  - Whenever you type air at the project root you'll see the server boot up and also the docs getting generated

Try accessing swagger at 
- http://localhost:8080/v1/swagger/index.html
- see it fail
