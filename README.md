# Clip Board

## Copy and Paste Across the Internet

### Quick Start

#### Dependencies

```sh
npm install
```

#### Development

```sh
make live
```

Navigate to `http://localhost:7331` on browser.

#### Production

```sh
make prod && cd prod
CLPBRD_ADDR=:8000 CLPBRD_DSN=./clipboard.db ./main
```

Navigate to `http://localhost:8000` on browser.

### To Do

- [x] ~~Update client with new data (web socket, server sent events, polling)~~
