# Asynchronous programming concurrency and parallelism
## callbacks
Use callbacks to do simple asynchronous tasks.

## promises
Example

```ts
function appendAndReadPromise(path: string, data: string): Promise<string> {
  return appendPromise(path, data)
    .then(() => readPromise(path))
    .catch(e => console.log(e))
}
```

Callback example

```ts
function appendAndRead(
  path: string,
  data: string,
  callback: (error: Error | null, result: string | null) => void
) {
  appendFile(path, data, error => {
    if (error) {
      return callback(error, null)
    }
    readFile(path, (error, result) => {
      if (error) {
        return callback(error, null)
      }
      callback(null, result)
    })
  })
}
```

