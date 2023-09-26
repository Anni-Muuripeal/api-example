## Code Architecture

- cmd
    * main.go (start point)

- model
    * model.go (Customer struct)

- pkg
    - api
        * api.go (handlers, request, response)
        * api_test.go

    - dao
        * dao.go (CustomerList, AddCustomer)
        * dao_test.go

    - service
        * service.go
        * service_test.go
