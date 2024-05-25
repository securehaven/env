# Env: Taming Environment Variables in Go

Tired of wrestling with environment variables in your Go apps? This nifty little library makes working with environment variables a breeze.

- **Type Like a Boss:** Nail down the exact data type you need, no more wrestling with unexpected values.
- **Defaults? You Got It:** No sweat, set a fallback value to keep your app running smoothly.

## Getting Started

Just a quick `go get` to start:

```sh
go get -u github.com/securehaven/env
```

## Usage

1. **Add the import**

    ```go
    import (
        "github.com/securehaven/env"
    )
    ```

2. **Access variables**

    - **Grab It, Always:** Need a variable, but not sure if it's always there? `env` get it for you, with a default in case it's missing.

        ```go
        port := env.Get[uint16]("PORT", 8000)
        ```

    - **Got It All Figured Out?:** You're the captain now! Decide exactly when to use defaults by handling the returned error.

        ```go
        port, err := env.GetStrict[uint16]("PORT", 8000)

        if !errors.Is(err, ErrNotFound) {
            panic(err)
        }
        ```

    - **Never Miss a Beat:** This variable is mission-critical? `env` will make sure you get it, or your app politely (but firmly) lets you know there's a problem.

        ```go
        apiKey := env.MustGet[string]("API_KEY")
        ```
