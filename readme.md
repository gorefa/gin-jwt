```bash
http -v --json POST localhost:8081/login username=admin password=admin

http -v -f GET localhost:8081/auth/refresh_token "Authorization:Bearer xxxtokenxxx"  "Content-Type: application/json"

http -f GET localhost:8081/auth/hello "Authorization:Bearer xxxtokenxxx"  "Content-Type: application/json"
```