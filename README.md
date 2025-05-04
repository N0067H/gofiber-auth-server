# Go/Fiber Auth Server
Authentication server using Go/Fiber V2 for my study.
 
## Endpoints

| Method | PATH  | details |
| :--- | :--- | :--- |
| GET | /ping | Ping Pong! |
| GET | /restricted | Bearer token required |
| POST | /register | Register |
| POST | /login | Login |

## Dotenv Configuration
```
DB_USER= [DATABASE USERNAME]
DB_PASS= [DATABASE PASSWORD]
DB_NAME= [DATABASE NAME]

JWT_SECRET= [JWT SECRET KEY]
```

## To-do
- User session with JWT
- Password hasing with Argon2
- Architecting (idk detail :p)

## Shit Spagetti Code
Contribute to my blossoming.
