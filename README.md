# Go Task Queue Benchmarking Project

This project benchmarks and demonstrates different task queue systems in Go. It includes implementations using **Asynq**, **xsync (MPMCQueue)**, **goq**, **Faktory**, and **Machinery**. Each system supports enqueue, dequeue, and peek operations where possible.

---

## Table of Contents

* [Features](#features)
* [Queue Implementations](#queue-implementations)
* [Environment Setup](#environment-setup)
* [Usage](#usage)
* [Folder Structure](#folder-structure)
* [Notes & Caveats](#notes--caveats)

---

## Features

* Asynchronous task enqueueing
* Peek tasks (where supported)
* Dequeue tasks (manual or worker-based)
* Supports Redis, RabbitMQ, or in-memory queues
* Demonstrates production-grade solutions and lightweight alternatives

---

## Queue Implementations

### 1. Asynq

* Redis-backed, production-grade task queue
* Supports retries, scheduling, and task prioritization
* Uses JSON-based payloads for cross-system portability

**Notes:**

* Best for production-grade background job processing
* Flexible queue priority system

---

### 2. xsync (MPMCQueue)

* Lock-free, in-memory, multi-producer multi-consumer queue
* Very fast, no external dependencies
* Optional shadow structure to safely peek while preserving FIFO order

**Notes:**

* Best for low-latency in-memory processing
* Peeking requires additional data structure to preserve FIFO order

---

### 3. goq

* CLI-based job management for Unix-like systems
* Executes shell commands asynchronously
* Peek and dequeue implemented via CLI commands

**Notes:**

* Simple to set up and use
* Consumes higher system resources under load
* Not recommended for large-scale queues

---

### 4. Faktory

* Language-agnostic background job system
* Supports retries, scheduling, and monitoring via web UI
* Workers handle job execution; manual dequeue can be implemented if needed

**Notes:**

* Suitable for cross-language background job processing
* Peek is not directly supported
* Web UI provides real-time monitoring

---

### 5. Machinery

* RabbitMQ-based distributed task queue
* Supports asynchronous processing with multiple workers
* Requires task registration
* Custom implementations can simulate peek/dequeue by directly accessing RabbitMQ

**Notes:**

* Ideal for distributed task processing
* Resource-intensive for large payloads
* Provides reliable delivery using RabbitMQ

---

## Environment Setup

1. Install **Go 1.18+**
2. Install required libraries using `go get`
3. Setup Redis (for Asynq), RabbitMQ (for Machinery), or Faktory server as required
4. Run worker processes first for queue-backed systems

---

## Usage

* Each queue implementation includes utilities for:

  * Enqueueing tasks
  * Dequeueing tasks (if supported)
  * Peeking tasks (if supported)
* Use the appropriate client/server setup for each implementation
* Ensure background workers are running for queue systems that require them

---

## Folder Structure

```
project-root/
│
├── asynqImp/
│   ├── main.go
│   └── tasks/
│       ├── enqueueTask.go
│       ├── dequeueTask.go
│       └── peekTask.go
│
├── xsyncImp/
│   └── main.go
│
├── goqImp/
│   ├── main.go
│   └── tasks/
│       ├── job_script.sh
│       ├── enqueueTask.go
│       ├── dequeueTask.go
│       └── peekTask.go
│
├── faktoryImp/
│   ├── main.go
│   └── tasks/
│       ├── enqueueTask.go
│       └── dequeueTask.go
│
├── machineryImp/
│   ├── config/
│   ├── server/
│   ├── tasks/
│   └── main.go
```

---

## Notes & Caveats

* **Asynq**: Best suited for production-grade background job processing; flexible queue priorities
* **xsync MPMCQueue**: Ideal for low-latency in-memory use; peeking requires a shadow structure
* **goq**: Simple but resource-intensive; not ideal for high-volume tasks
* **Faktory**: Cross-language support with web UI monitoring; peek not supported
* **Machinery**: Distributed processing with RabbitMQ; custom code needed for peek/dequeue; can be resource-intensive for large datasets

---

This project serves as a reference for selecting the right task queue in Go depending on use case, performance needs, and infrastructure availability.
