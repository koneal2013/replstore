# key-value store

In this exercise we ask you to write a command line REPL (read-eval-print loop) that drives a simple in-memory key/value storage system. This system should also allow for nested transactions. A transaction can then be committed or aborted.

Submitting Code: Please avoid posting your code publicly, as we will reuse this exercise for other candidates. You have commit access to this github repository, so you should submit your code as a merge request(s). Make sure it's documented how to run your repl if it's something other than `go run main.go`.

We realize that your time is valuable. Please avoid spending more than 4 hours on this exercise. We expect those with a lot of go expertise to get further, and 100% completion is not necessary.

This repository has built-in CI that will run any go tests that you define. You can change CI behavior, including the image that is used to run the tests, by editing the `.github/workflows/go.yml` file if you'd like.

## Example Run

```
$ my-program
> WRITE a hello
> READ a
hello
> START
> WRITE a hello-again
> READ a
hello-again
> START
> DELETE a
> READ a
Key not found: a
> COMMIT
> READ a
Key not found: a
> WRITE a once-more
> READ a
once-more
> ABORT
> READ a
hello
> QUIT
Exiting...
```

We recommend solving simpler parts of the problem before adding on more
advanced things like transactions.

## Command

    READ   <key> Reads and prints, to stdout, the val associated with key. If
           the value is not present an error is printed to stderr.
    WRITE  <key> <val> Stores val in key.
    DELETE <key> Removes all key from store. Future READ commands on that key
           will return an error.
    START  Start a transaction.
    COMMIT Commit a transaction. All actions in the current transaction are
           committed to the parent transaction or the root store. If there is no
           current transaction an error is output to stderr.
    ABORT  Abort a transaction. All actions in the current transaction are discarded.
    QUIT   Exit the REPL cleanly. A message to stderr may be output.

## Other Details

* For simplicity, all keys and values are simple ASCII strings delimited by whitespace. No quoting is needed.
* All errors are output to stderr.
* Commands are case-insensitive.
* As this is a simple command line program with no networking, there is only one “client” at a time. There is no need for locking or multiple threads.

## Future Enhancements

During our call after the interview, we will discuss how you would add additional enhancements, such as:

* Accept network connections from multiple clients, speaking this same protocol as described above.
* Make sure multiple clients can safely mutate data.
* How to persist this data to disk.
* What sort of metrics you would add to decide if this service was healthy.
