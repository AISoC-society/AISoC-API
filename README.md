# AISoC credits API:

## Compilation:

### Compile Time Dependencies:

1. `go` 1.18
1. `make`

### Runtime Dependencies:

1. `sqlite`

### Steps:

```bash
$ make
```

## Running:

1. Copy `.env-example` provided in the project root and create a corresponding `.env` file.
1. Fill `.env` with the required fields.
1. Compile the application and run the binary.

## Contributing:

1. Fork this repository.
1. Make all your changes against the main branch in a patch branch.
1. Run `make check` to make sure all tests and linting checks pass.
1. Follow standard golang practices to reduce heap allocations! This is done to ensure that the compiler can inline most of our code for good performance.
1. Run `make allocs` to make sure no new allocations are made on the heap that are unecessary.
1. All commit messages should adhere to [conventional commits.](https://www.conventionalcommits.org/en/v1.0.0/)
