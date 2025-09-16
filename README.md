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

# Transaction Isolation Levels & Read Phenomena

## Introduction
Database transactions allow multiple operations to be executed as a single **atomic** unit.  
However, when multiple transactions run concurrently, they can **interfere** with each other, leading to inconsistent or unexpected results.

To control this, relational databases provide **isolation levels**.  
Isolation levels determine **how visible changes made by one transaction are to others** before they are committed.

---

## Read Phenomena (Concurrency Anomalies)

### 1. Dirty Read
- **Definition**: Reading uncommitted data from another transaction.

**SQL Example** (works in databases that allow dirty reads, e.g., SQL Server with `READ UNCOMMITTED`. PostgreSQL does NOT allow dirty reads at all):

```sql
-- Transaction A
BEGIN;
UPDATE accounts SET balance = balance + 100 WHERE id = 1;

-- Transaction B (with READ UNCOMMITTED in some DBs)
SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
SELECT balance FROM accounts WHERE id = 1;  -- Sees uncommitted update!
```

If Transaction A rolls back, B has read a value that never existed.

---

### 2. Non-Repeatable Read
- **Definition**: Reading the same row twice in the same transaction returns different values.

**SQL Example** (`READ COMMITTED`):

```sql
-- Transaction A
BEGIN;
SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
SELECT balance FROM accounts WHERE id = 1;  -- returns 100

-- Transaction B
BEGIN;
UPDATE accounts SET balance = 200 WHERE id = 1;
COMMIT;

-- Back to Transaction A
SELECT balance FROM accounts WHERE id = 1;  -- returns 200 now (changed)
COMMIT;
```

---

### 3. Phantom Read
- **Definition**: Re-running a query returns different sets of rows due to inserts/updates/deletes by another transaction.

**SQL Example** (`READ COMMITTED` or `REPEATABLE READ` in MySQL):

```sql
-- Transaction A
BEGIN;
SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
SELECT * FROM orders WHERE amount > 100;  
-- returns 3 rows

-- Transaction B
BEGIN;
INSERT INTO orders (id, amount) VALUES (999, 150);
COMMIT;

-- Back to Transaction A
SELECT * FROM orders WHERE amount > 100;  
-- returns 4 rows (phantom row appears)
COMMIT;
```

PostgreSQL `REPEATABLE READ` prevents this because it uses a **consistent snapshot**.

---

### 4. Serialization Anomaly
- **Definition**: Concurrent execution produces a result that couldn’t happen if transactions ran one after another.

**SQL Example** (`READ COMMITTED`):

```sql
-- Assume account 1 has balance = 100

-- Transaction A
BEGIN;
SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
SELECT balance FROM accounts WHERE id = 1;  -- sees 100
-- decides to withdraw 100
UPDATE accounts SET balance = balance - 100 WHERE id = 1;

-- Transaction B (concurrent)
BEGIN;
SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
SELECT balance FROM accounts WHERE id = 1;  -- also sees 100
-- decides to withdraw 100
UPDATE accounts SET balance = balance - 100 WHERE id = 1;

-- Both commit → Final balance = -100 (impossible in serial execution)
```

At `SERIALIZABLE` level, one of these transactions would be rolled back to prevent inconsistency.

---

## ANSI SQL Isolation Levels

| Isolation Level      | Dirty Reads | Non-Repeatable Reads | Phantom Reads | Notes |
|----------------------|-------------|----------------------|---------------|-------|
| **Read Uncommitted** | Possible    | Possible             | Possible      | Fastest, lowest safety. Rarely used. |
| **Read Committed**   | Prevented   | Possible             | Possible      | Default in PostgreSQL, Oracle. |
| **Repeatable Read**  | Prevented   | Prevented            | Possible*     | Default in MySQL InnoDB. (*Postgres eliminates phantoms here.) |
| **Serializable**     | Prevented   | Prevented            | Prevented     | Strongest, ensures full serial execution equivalence. |

---

## PostgreSQL Notes
- **Read Committed** (default): Each statement sees only committed data at the start of execution.
- **Repeatable Read**: Prevents non-repeatable reads and phantoms; queries see a consistent snapshot.
- **Serializable**: Uses **Serializable Snapshot Isolation (SSI)**. If anomalies could occur, one transaction is aborted.

---

## Choosing an Isolation Level

- **Read Uncommitted**: Almost never recommended.
- **Read Committed**: Good balance of performance and correctness for most OLTP workloads.
- **Repeatable Read**: Safer for reporting, analytics, and financial calculations.
- **Serializable**: Use when correctness is critical and anomalies cannot be tolerated (e.g., money transfers, inventory management). Be prepared for **retries**.

---

## Example: Bank Transfer (Deadlock & Isolation)

- **Scenario**: 
  - T1: Transfer $100 from Account A → B.
  - T2: Transfer $50 from Account B → A.

- If both update rows in opposite order, a **deadlock** may occur.
- If running at lower isolation levels, anomalies may cause **double spending**.
- Recommended: 
  - Use **Repeatable Read** or **Serializable**.
  - Implement **retry logic** on deadlock/serialization errors.

---

## Summary
- **Isolation levels** balance between **performance** and **consistency**.
- Always understand which **read phenomena** your application can tolerate.
- For critical systems (like finance), use **Serializable** with retries.
- For high-performance systems, **Read Committed** is often enough with careful design.

---
