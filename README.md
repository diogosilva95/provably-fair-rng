# Provably Fair RNG

A provably fair random number generator implementation in Go using SHA-512.

## Usage

```go
go run main.go
```

## Api

- `games.RNG(clientSeed, serverSeed, nonce, cursor)` - Returns random number 0-9999
- `games.DiceRoll(clientSeed, serverSeed, nonce)` - Returns dice roll 0.00-99.99

## How it works

1. Combines seeds and nonce
2. SHA-512 hash
3. Extract hex values
4. Return random number
