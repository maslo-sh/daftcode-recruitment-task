# Daftcode Recruitment Task

## Project Structure

### cmd/

Main entrypoint of the project.

### internal/

Catalog containing all Gin server code and tests of specific endpoints.

### scripts/

Catalog containing scripts for building and running application as a Docker container.

## TBD

- checking rates of exchange to the same value (f. ex. USD - USD); should it give any result?
- adding one non-existing currency to the list of appropriate values (f. ex. __?currencies=GBP,PLN,ABC__); should it return data calculated only for real currencies?