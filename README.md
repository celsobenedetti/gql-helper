# GraphQL Helper

A sandbox project to hack around utils to convert `query.gql` files into HTTP requests.

## MVP

1. Read `query.gql` file
2. Read `vars.json` file
3. Make HTTP request with correct headers, data and vars

## Improvements

- [ ] Generate equivalent `curl` request for sharing
- [ ] `-i` flag for searching through folder for all `*.gql` requests and `vars.json`
