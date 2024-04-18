# think-money-technical-test

## Setup
Either run the following commands:
- Manual:
    - `go mod download`
- Automatic: `make setup`

## Usage
Please note, if you don't have `make` installed, you can open the makefile to find the relevant commands that need to be ran. 

| Command | Description | Flags |
| --- | --- | --- |
| `make setup` | Installs all dependencies for the code | |
| `make run` | Runs a provided main file | path=`{path to main file}` |
| `make run_rest_api` | Runs the wired up rest API in {project root}/main.go | |
| `make run_devtest` | Runs the dev test tool within cli/* | |
| `make run_tests` | Runs all tests within the code repository | additionalParams=`{ any golang 'go run' parameter}` |
| `make run_test_coverage` | Runs all tests, creates test_coverage.out in the root of the project, opens your browser of choice with results of coverage | |
| `make run_benchmarks` | Runs defined benchmarks | |
| `make swag_generate` | Generates the swagger files for the passed main file | |

### Running the REST API
To run the REST API you need to:
- Run `make setup`
- Run `make run_rest_api`
- Wait for the API to start
- Navigate to [the swagger homepage](http://localhost:4000/swagger/index.html)
- Utilize the various endpoints to:
    - Scan items into the checkout
    - Update discounts
    - Create items based on sentences in the format `{ optional[int] - quantity for discount } { [string] - name of item } cost { [int] - cost of discount or single item }`. These can also be stacked by seperating them with ',' or '.'. (Please note that this only works with items already registered), Some examples are:
        - Pineapples cost 50, Waffles cost 30.
        - 2 Pineapples cost 75.
        - Bacon cost 10, 3 Bacon cost 25. Waffles cost 2, 2 Waffles cost 3, Pineapples cost 75.
    - Retrieve the total of the checkout

The configuration for bound IP and port can be changed by opening `{project root}/main.go` and updating the following const variables located near the top of the file:
- DEFAULT_IP - The standard IP to bind to on the system (currently defaulted to "0.0.0.0" - aka, every IP address)
- DEFAULT_PORT - The port to bind to on the system (currently defaulted to 4000)

### Adding items to the catalogue
To support additional SKUs / items / initial discount, you will need to update the `main.go` file found in the root of the project.
- To add additional catalogue items, you can call `myCatalogue.Create(sku, checkout.NewItem(name, cost))` after myCatalogue has been created but before the line that contains `err := r.Run(...)`
    - Please note that if you attempt to overwrite an already existing SKU using the Create function, you will receive an error
- To add additional initial discounts, you can call `myDiscountCatalogue.Create(sku, checkout.NewDiscount(quantity, price))` after myDiscountCatalogue has been created but before the line that contains `err := r.Run(...)`
    - Please note that if you attempt to overwrite an already existing SKU using the Create function, you will receive an error
- To add an initial sentence that you would like to be parsed (as outlined above), you can modify the `INITIAL_PARSED_SENTENCE` const variable declared near the top of the main.go file to contain the sentence you would like to parse.


## Installed packages
- [testify](https://github.com/stretchr/testify)
- [gin](github.com/gin-gonic/gin)

## Final notes
- Why isn't there a use of channels?
    - There wasn't a massive need for them
    - They add complexity to testing, due to the short time deadline an MVP is better than nothing. This can be included later.
- Are there any examples of goroutines?
    - There is an example of goroutines used in the repository bench mark test
- Why does the memory repository only have a mutex for when data is updated?
    - We want to block the system from updating the same "row" at the same time, as otherwise data could be corrupted. This way, reading will always be faster than writing, but we guarentee data integrity.
- I wasn't sure how to pick apart `Implement the code for a checkout system that handles pricing schemes such as "pineapples cost 50, three pineapples cost 130."`
    - After much deliberation I determined that currently I don't know how I would implement the sentence parser to match `{ any number of numbers in word form } { item name } cost { item cost }`. As such, I simplified the problem to allow us to create a solution that will work moving forwards, whilst also providing us with a way to implement this solution at a later date.
- Why was everything developed on 1 branch?
    - As I was working on this by myself, 1 branch is all that was needed. With additional people working alongside me I would have created a skeleton repository first and then would have done small individual changes on a separate branch.
- Why didn't I use mocks for testing?
    - Currently the tests are not written using the best of practices, however as all the components are using a memory store, as this is what the mocks would have most likely utilize, there aren't many issues using it. However, if we were utilizing Postgres, MongoDB or other distributed data store, it would be best to mock the return values of these results (preventing uneccesary calls to an actual database).