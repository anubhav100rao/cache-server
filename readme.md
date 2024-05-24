# In-Memory Cache server

This is a simple in-memory cache server that supports the following operations:

1. SET
2. GET
3. DELETE
4. CLEAR
5. GETALL

Other features include:

1. Cache eviction policies (LRU, LFU, FIFO, LIFO, Random)
2. Cache expiration policies
3. Cleanup of expired keys using a background thread

## How to run

1. Clone the repository
2. Run the following command to start the server:

```bash
go run main.go
```
