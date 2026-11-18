## Current Architecture

- Job: Represents a unit of scheduled work.
- Collection: Stores jobs in insertion order and indexes them by ID.
- BoundedQueue: Stores pending jobs in FIFO order.
- JobStore: Defines the storage contract.
- Scheduler Service: Validates and manages job submission.

## Current Limitations

- Storage is in-memory.
- No persistence.
- No concurrent access protection.
- No worker execution yet.