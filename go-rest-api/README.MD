

### Used libraries / packages

https://github.com/gorilla/mux (Routing)
 - It was the course suggestion, so far the advantage is to limit the method type directly in the routes class. (e.g.: Route "/whatever" can only accept GET requests, that can be done with a simple call) Without it, every handler function needs to first validate the method manually and then proceed.

https://gorm.io/index.html gorm (ORM)
 
https://github.com/gorilla/handlers (CORS)