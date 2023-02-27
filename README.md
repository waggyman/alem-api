# alem-api
Backend API for ALEM (Academic and Learning Management) - Written in GO LANG

## Installing Steps
To use this repository you need to have:
- GO Compiler in your PC
- MongoDB in your local or Atlas

### Steps
- Clone this repository using `git clone git@github.com:waggyman/alem-api.git`
- Copy the `.env.example` to `.env`
- Fill the value of env variables in your `.env` file
- Run the app with `go run .` in your terminal 
- After you ran it successfully, you can check it with running:
```
curl --location --request GET 'localhost:8000/teachers'
```