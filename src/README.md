# apex-network

<!-- ABOUT THE PROJECT -->
## About The Project
This a simple dice game


### Built With
1. [Golang](https://go.dev)
2. [Gin-gonic](https://github.com/gin-gonic/gin)
3. [PostgreSQL](https://www.postgresql.org/)
4. [GORM](https://gorm.io/)
5. REST
6. [Mockery](https://pkg.go.dev/github.com/knqyf263/mockery) for testing


<!-- GETTING STARTED -->
## Getting Started
Clone the project from GitHub using the command
```sh
git clone git@github.com/tejiriaustin/apex-network.git
```

### Prerequisites

* Go
    ```sh
    https://go.dev
    ```

### Starting the project

1. run command
    ```sh
    go mod tidy
    ```
   to index files and dependencies
2. Create file and name `.env`
3. Contact the admin of this project to get your personal credentials
4. Run command to start the application
    ```sh
   make app
   ```
5.Test on postman

## Routes avaliable
- Create player
```
    URL: {{url}}/game/create-player
    Body: {
      "first_name": "Tejiri",
      "last_name": "Dev" 
      }
```

- Start Game
```
    URL: {{url}}/game/start/:player_id
```
- End Game
```
    URL: {{url}}/game/end/:player_id
```
- Roll Dice
```
    URL: {{url}}/game/roll-dice/:player_id
```
- Fetch Balance
```
    URL: {{url}}/game/balance/:player_id
```
- Fund Wallet
```
    URL: {{url}}/game/fund-wallet/:player_id
```
- List Transactions
```
    URL: {{url}}/game/transactions/:player_id
```
- Is Playing
```
    URL: {{url}}/game/is-playing/:player_id
```

[Golang-URL]: https://go.dev 
