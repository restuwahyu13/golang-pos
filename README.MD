# Golang Point Of Sale

Example Golang API backend rest implementation mini project Point Of Sale using Gin Framework and Gorm ORM Database.

## Command

- ### App Lifecyle

```sh
$ go run main.go || make go run
```

- ### Docker Lifecycle

```sh
$ docker-compose up -d --build
```
## Endpoint

| **Nama**        | **Route**                  | **Method** |
| --------------- | -------------------------- | ---------- |
| **User**        |                            |            |
|                 | */api/v1/auth/ping*        | *GET*      |
|                 | */api/v1/auth/register*    | *POST*     |
|                 | */api/v1/auth/login*        | *POST*     |
| **Merchant**    |                            |            |
|                 | */api/v1/merchant/ping*    | *GET*      |
|                 | */api/v1/merchant*         | *POST*     |
|                 | */api/v1/merchant*         | *GET*      |
|                 | */api/v1/merchant/:id*     | *GET*      |
|                 | */api/v1/merchant/:id*     | *DELETE*   |
|                 | */api/v1/merchant/:id*     | *UPDATE*   |
| **Outlet**      |                            |            |
|                 | */api/v1/outlet/ping*      | *GET*      |
|                 | */api/v1/outlet*           | *POST*     |
|                 | */api/v1/outlet*           | *GET*      |
|                 | */api/v1/outlet/:id*       | *GET*      |
|                 | */api/v1/outlet/:id*       | *DELETE*   |
|                 | */api/v1/outlet/:id*       | *PUT*      |
| **Customer**    |                            |            |
|                 | */api/v1/outlet/ping*      | *GET*      |
|                 | */api/v1/customer*         | *POST*     |
|                 | */api/v1/customer*         | *GET*      |
|                 | */api/v1/customer/:id*     | *GET*      |
|                 | */api/v1/customer/:id*     | *DELETE*   |
|                 | */api/v1/customer/:id*     | *PUT*      |
| **Product**     |                            |            |
|                 | */api/v1/product/ping*     | *GET*      |
|                 | */api/v1/product*          | *POST*     |
|                 | */api/v1/product*          | *GET*      |
|                 | */api/v1/product/:id*      | *GET*      |
|                 | */api/v1/product/:id*      | *DELETE*   |
|                 | */api/v1/product/:id*      | *PUT*      |
| **Role**        |                            |            |
|                 | */api/v1/role/ping*        | *GET*      |
|                 | */api/v1/role*             | *POST*     |
|                 | */api/v1/role*             | *GET*      |
|                 | */api/v1/role/:id*         | *GET*      |
|                 | */api/v1/role/:id*         | *DELETE*   |
|                 | */api/v1/role/:id*         | *PUT*      |
| **Supplier**    |                            |            |
|                 | */api/v1/supplier/ping*    | *GET*      |
|                 | */api/v1/supplier*         | *POST*     |
|                 | */api/v1/supplier*         | *GET*      |
|                 | */api/v1/supplier/:id*     | *GET*      |
|                 | */api/v1/supplier/:id*     | *DELETE*   |
|                 | */api/v1/supplier/:id*     | *PUT*      |
| **Transaction** |                            |            |
|                 | */api/v1/transaction/ping* | *GET*      |
|                 | */api/v1/transaction*      | *POST*     |
|                 | */api/v1/transaction*      | *GET*      |
|                 | */api/v1/transaction/:id*  | *GET*      |
|                 | */api/v1/transaction/:id*  | *DELETE*   |
|                 | */api/v1/transaction/:id*  | *PUT*      |