# Distributed Queue Evaluation in Go

This repository contains implementations and benchmarks of various Go-based distributed queue systems evaluated for handling high-load scenarios. The goal is to buffer user requests during peak load, process them fairly, and maintain low latency with high throughput. Below is a brief explanation of each queue system, their folder structures, environment setup instructions, dependencies, and links to their respective repositories for further details.

## Queue Systems Evaluated

### 1. Asynq
**Description**: Asynq is a Redis-backed distributed task queue for Go, designed for scalability and ease of use. It supports task scheduling, retries, priorities, monitoring via CLI and Web UI, and integration with Prometheus. Asynq is ideal for production-grade systems needing robust task management.

**Folder Structure**:
```
asynqImp/
├── main.go
└── tasks/
    ├── enqueueTask.go
    ├── dequeueTask.go
    └── peekTask.go
```

**Environment Setup**:
- Go installed (v1.18+ recommended)
- Redis running locally (Asynq uses Redis as its backing store)

**Dependencies**:
```bash
# Start Redis using Docker
docker run -p 6379:6379 redis

# Install the Go SDK
go get github.com/hibiken/asynq
```

**For More Info**: [Asynq GitHub Repository](https://github.com/hibiken/asynq)

---

### 2. xsync (MPMCQueue)
**Description**: xsync provides a high-performance, lock-free, in-memory multi-producer, multi-consumer (MPMC) queue optimized for concurrent Go applications. It uses atomic operations to minimize contention, making it suitable for high-throughput, low-latency systems without persistence needs.

**Folder Structure**:
```
xsyncImp/
├── main.go
└── tasks/
    ├── enqueueTask.go
    ├── dequeueTask.go
    └── peekTask.go
```

**Environment Setup**:
- Go installed (v1.18+ recommended)
- No external dependencies (in-memory queue)

**Dependencies**:
```bash
# Install the xsync package
go get github.com/puzpuzpuz/xsync/v3
```

**For More Info**: [xsync GitHub Repository](https://github.com/puzpuzpuz/xsync)

---

### 3. goq
**Description**: goq is a minimal, secure distributed job queue with TLS 1.3 and QUIC support, designed for fast job dispatching and fault tolerance. It was dropped from full benchmarking due to high resource consumption at low volumes but is included for completeness.

**Folder Structure**:
```
goqImp/
├── main.go
└── tasks/
    ├── job_script.sh
    ├── enqueueTask.go
    ├── dequeueTask.go
    └── peekTask.go
```

**Environment Setup**:
- Go installed (v1.18+ recommended)
- Unix-based system (Linux or macOS recommended)

**Dependencies**:
```bash
# Clone and build goq
mkdir -p $GOPATH/src/github.com/glycerine
cd $GOPATH/src/github.com/glycerine
git clone https://github.com/glycerine/goq
cd goq
make

# Set up GOQ_HOME and PATH
export GOQ_HOME=$HOME
export PATH=$GOPATH/bin:$PATH
# Add to ~/.bashrc or ~/.zshrc for persistence
echo 'export GOQ_HOME=$HOME' >> ~/.bashrc
echo 'export PATH=$GOPATH/bin:$PATH' >> ~/.bashrc
source ~/.bashrc

# Start the goq job server
goq serve &
```

**For More Info**: [goq GitHub Repository](https://github.com/glycerine/goq)

---

### 4. Faktory
**Description**: Faktory is a language-agnostic background job system with a centralized work server, supporting job retries, expiration, and a Web UI for monitoring. It simplifies client-side logic but does not support traditional queue operations like peek.

**Folder Structure**:
```
FaktoryImp/
├── main.go
└── tasks/
    ├── enqueueTask.go
    └── dequeueTask.go
```

**Environment Setup**:
- Go installed (v1.18+ recommended)
- Faktory server running

**Dependencies**:
```bash
# Start Faktory using Docker
docker run -d -p 7419:7419 -p 7420:7420 -p 7421:7421 --name faktory contribsys/faktory

# Install the Go worker library
go get -u github.com/contribsys/faktory_worker_go
```

**For More Info**: [Faktory GitHub Repository](https://github.com/contribsys/faktory)

---

### 5. Machinery
**Description**: Machinery is a flexible distributed task queue supporting multiple brokers (Redis, RabbitMQ, AWS SQS) and result backends. It offers task chaining and retry logic but consumes more resources and may encounter I/O timeouts under high load.

**Folder Structure**:
```
Machinery/
├── config/
│   └── config.go
├── server/
│   └── server.go
├── tasks/
│   ├── dequeueTask.go
│   ├── enqueueTask.go
│   └── peekTask.go
└── main.go
```

**Environment Setup**:
- Go installed (v1.18+ recommended)
- RabbitMQ server running

**Dependencies**:
```bash
# Option 1: Install RabbitMQ natively (Linux)
sudo apt install rabbitmq-server
sudo systemctl start rabbitmq-server

# Option 2: Run RabbitMQ via Docker
docker run -d --hostname rabbitmq-host --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

# Install Machinery
go get github.com/RichardKnop/machinery/v1
```

**For More Info**: [Machinery GitHub Repository](https://github.com/RichardKnop/machinery)

---

## Conclusion
Based on benchmark results evaluating latency for enqueue, peek, and dequeue operations under loads from 10 to 10,000 operations, the following conclusions were drawn:

- **xsync**: Best for high-throughput, low-latency in-memory queuing. It outperformed others in raw performance but requires manual overflow handling due to its fixed-size queue.
- **Asynq**: Offers the best balance of performance and production-ready features like retries, scheduling, and monitoring, making it ideal for background job processing.
- **Machinery**: Suitable for complex distributed systems with multi-backend support but consumes more resources and may fail under high load due to I/O timeouts.
- **Faktory**: Excels in polyglot environments with its centralized server and Web UI but lacks traditional queue operations like peek, functioning more as a job dispatcher.
- **goq**: Dropped from full evaluation due to excessive resource consumption, even at low volumes, making it less practical for high-load scenarios.

This repository provides practical implementations of each queue system, enabling developers to experiment and choose the best tool for their specific use case, whether prioritizing throughput, operational features, or language interoperability.