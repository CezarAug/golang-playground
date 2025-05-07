A really simple API now with Gin

Testing random functionalities, do not try in real environments


Pending tasks:
 - Separate different contexts into their own files
 - Do not expose GORM objects directly
 - Update tests, and organize them properly (can they be in a test package?)
   - Also in tests, use another DB
 - Create a better notion of environments, and have this application configurable for that
 - Dockerize

For future investigations:
 - Auth, how to add some simple auth
 - Logging: improve it and have it ready for actual usage
 - Metrics: how they can be collected
 - Testing: Check how to perform unit tests, also is there any BDD to use?
 - OWASP: There are dependencies in the project, so vulnerabilities also have to be monitored


### Dependencies

- GIN (github.com/gin-gonic/gin)
 - Playing around with a proper web framework instead of doing everything manually as the previous study

- GORM (ORM + Driver) (gorm.io/gorm + gorm.io/driver/postgres) 

- Validator (https://pkg.go.dev/github.com/go-playground/validator/v10)
   - Course suggestion was go get gopkg.in/validator.v2, but the usage in 2025 seems to be pretty low

- Testify (go get github.com/stretchr/testify)
   - Assertions in tests
