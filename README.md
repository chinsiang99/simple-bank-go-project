# simple-bank-go-project
simple bank go project


# Migration in Golang
```zsh
$ brew install golang-migrate
```

## Usage
```
$ migrate -help
Usage: migrate OPTIONS COMMAND [arg...]
       migrate [ -version | -help ]

Options:
  -source          Location of the migrations (driver://url)
  -path            Shorthand for -source=file://path
  -database        Run migrations against this database (driver://url)
  -prefetch N      Number of migrations to load in advance before executing (default 10)
  -lock-timeout N  Allow N seconds to acquire database lock (default 15)
  -verbose         Print verbose logging
  -version         Print version
  -help            Print usage

Commands:
  create [-ext E] [-dir D] [-seq] [-digits N] [-format] [-tz] NAME
           Create a set of timestamped up/down migrations titled NAME, in directory D with extension E.
           Use -seq option to generate sequential up/down migrations with N digits.
           Use -format option to specify a Go time format string. Note: migrations with the same time cause "duplicate migration version" error.
           Use -tz option to specify the timezone that will be used when generating non-sequential migrations (defaults: UTC).

  goto V       Migrate to version V
  up [N]       Apply all or N up migrations
  down [N] [-all]    Apply all or N down migrations
        Use -all to apply all down migrations
  drop [-f]    Drop everything inside database
        Use -f to bypass confirmation
  force V      Set version V but don't run migration (ignores dirty state)
  version      Print current migration version
```

## to create migration file
- note that -seq is to create sequential init_schema is the name of file
> migrate create -ext sql -dir internal/database/migrations -seq init_schema

# Sqlc in golang

## Installation

> brew install sqlc

> sqlc version

> sqlc init

# Database Design
![ERD Diagram](docs/erd-diagram.png)

# Deadlocks in Database Transactions

## What is a Deadlock?

A **deadlock** happens when two or more transactions are waiting for each other to release a lock.  
Since each one is waiting, none of them can proceed — the database detects this and aborts one of the transactions with an error like:

```
deadlock detected
```

This is PostgreSQL’s way of preventing the system from freezing completely.

---

## Example Scenario: Bank Transfer

Imagine two accounts: **A** and **B**.  
Two transfers are happening at the same time:

1. **Tx1**: Transfer 100 from **A → B**  
   - Locks account **A**
   - Then tries to lock account **B**

2. **Tx2**: Transfer 200 from **B → A**  
   - Locks account **B**
   - Then tries to lock account **A**

Now we have a cycle:
- Tx1 holds lock on **A**, waiting for **B**
- Tx2 holds lock on **B**, waiting for **A**

This is a **deadlock**.

---

## How to Avoid Deadlocks

### ✅ 1. Consistent Lock Ordering
Always acquire locks in a **fixed order**.  
In our case:
- Always update the account with the **smaller ID first**, then the larger ID.
- This ensures that two concurrent transfers won’t end up waiting on each other.

```go
if arg.FromAccountID < arg.ToAccountID {
    // lock/update FromAccount first, then ToAccount
} else {
    // lock/update ToAccount first, then FromAccount
}
```

### ✅ 2. Keep Transactions Short
Avoid long-running queries inside a transaction.  
The shorter the transaction, the lower the chance of blocking other transactions.

### ✅ 3. Retry on Deadlock
Even with precautions, deadlocks can still happen.  
It’s common to catch the `deadlock detected` error and retry the transaction automatically.

---

## Summary

- Deadlocks occur when transactions wait on each other’s locks.  
- PostgreSQL resolves this by aborting one of them.  
- Use **consistent lock ordering**, keep transactions short, and **retry on deadlock** to ensure safe and reliable transfers.
