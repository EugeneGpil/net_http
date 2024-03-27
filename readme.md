# Conclusions

### Problem

1. Url `/item` will work for request `/item`, but not for `/item/`.
2. Url `/item/` will redirect request `/item` to `/item/`, which is unexpected and should be avoided.

### Solution

For one handler function should be defined 2 routes with urls `/item` and `/items/`.

### Problem

Url `/item/` will also work for any url subtrees, like `/item/foo/bar`, `/item/foo`, which is also unexpected and should be avoided.

### Solution

In route handler the request url should be inspected on exact math with defined in route url.
