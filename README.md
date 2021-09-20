Becklyn gocruddy
=================

[![CI](https://github.com/Becklyn/gocruddy/actions/workflows/ci.yml/badge.svg)](https://github.com/Becklyn/gocruddy/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/Becklyn/gocruddy/branch/main/graph/badge.svg?token=YO6PLZ30RN)](https://codecov.io/gh/Becklyn/gocruddy)

Made with ❤ by Becklyn

A framework that supports you in creating CRUD APIs using go (golang).

This framework is useful for applications that use [fiber](https://github.com/gofiber/fiber) as router and [gorm](https://gorm.io/) as ORM. 

Installation
------------

```shell
go get -u github.com/Becklyn/gocruddy
```

Usage
-----

```go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Becklyn/gocruddy"
)

type UserCrud struct {
	container gocruddy.Container
}

func (uc *UserCrud)  {
	
}

func main() {
	app := fiber.New()
	
	config := 

	RegisterAdminRoutes(app, c)

	app.Listen(":3000")
}
```

Development
-----------

Set up your local development environment:

````shell
make setup
````

Add new modules / dependencies:

```shell
make install MOD=your.dependency/name
```

Tidy up modules:

```shell
make tidy
```

Test your implementation:

```shell
make test
```

Calculate the code coverage:
(This currently requires a local go installation)

```shell
make cover
```

References
----------

This project makes use of some really great packages. Please make sure to check them out!

| Package                                                                  | Usage          |
| ------------------------------------------------------------------------ | -------------- |
| [github.com/ao-concepts/logging](https:/github.com/ao-concepts/logging)  | Logging        |
| [github.com/ao-concepts/storage](https://github.com/ao-concepts/storage) | DB abstraction |
| [github.com/gofiber/fiber](https://github.com/gofiber/fiber)             | HTTP router    |
| [github.com/stretchr/testify](https://github.com/stretchr/testify)       | Testing        |
| [gorm.io/gorm](https://gorm.io/)                                         | ORM            |
