# bstreee

Binary search tree service with http access

## Usage

search value `http GET localhost:8080/search?val=40`

insert value `http -j POST localhost:8080/insert val:=40`

delete value `http DELETE localhost:8080/delete?val=40`
