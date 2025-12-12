# 🛍️ Go Minimal E-Commerce Backend

- This project is a robust E-Commerce backend built with the Go Standard Library, applying DDD (Domain-Driven Design) and Hexagonal Architecture.

- It aims to build a solid and scalable architecture without relying on external web frameworks (like Gin or Echo), focusing on the core capabilities of Go.
  🏗️ Project Architecture

- This project follows the Encapsulated Core architectural style. The application package acts as the Core of the business, encompassing domain entities, use cases, and ports (interfaces). The adapter layer connects to this core like a plugin.

## 📂 Folder Structure

```txt
my-shop/
├── cmd/
│ └── api/ # 🚀 Entrypoint: Assembly and execution
│
├── internal/
│ └── product/ # 📦 [Product Context]
│ │
│ ├── adapter/ # 🔌 [Adapter Layer] (Outside World)
│ │ │ # Depends on the Core (application) and implements specific technologies.
│ │ ├── http/ # - [Primary] Web Handler, Router
│ │ ├── storage/ # - [Secondary] DB Implementation (Postgres)
│ │ └── infra/ # - [Secondary] Infra Implementation (IDGen)
│ │
│ └── application/ # 💎 [Application Core] (The Hexagon)
│ │ # The aggregate of business logic; zero external dependencies.
│ │
│ ├── domain/ # 🧠 [Pure Domain]
│ │ ├── product.go # - Aggregate Root & Business Logic
│ │ └── vo_money.go # - Value Objects
│ │
│ ├── port/ # 🚪 [Ports] (Interfaces)
│ │ ├── repo.go # - Repository Interface (Driven Port)
│ │ └── usecase.go # - (Optional) Service Interface (Driving Port)
│ │
│ └── service/ # ⚙️ [Service] (Use Cases)
│ └── service.go # - Business flow & Transaction management
│
└── sqlc.yaml # Type-Safe SQL configuration

```

## 🏛️ Design Decisions

### 1. Encapsulated Application Core

We have grouped domain, service, and port under the application directory.

    Why? To increase cohesion by gathering "everything related to business logic" into a single Core folder (application).

        This clearly defines the adapter vs. application boundary.

        The adapter layer (technology) imports the application layer (business), while the reverse direction is strictly forbidden.

### 2. Explicit Ports Package

Interfaces are isolated in the application/port package.

    Why? To prevent dependency confusion that can occur when interfaces are mixed within domain or service packages.

        Repository Interface: Defined in application/port/repo.go.

        Implementation: adapter/storage imports the port package to implement it.

        Usage: application/service imports the port package to use it.

### 3. Pure Domain Entities

The application/domain package consists of pure Go code with zero external dependencies.

    Why? To ensure that core business rules remain unaffected even if frameworks or database technologies change.

### 4. Standard Library & Tools

No Framework: We focus on the essence of HTTP communication using net/http and ServeMux.

Type-Safe SQL: We use sqlc to generate type-safe Go code from SQL queries.

Structured Logging: We use slog for structured JSON logging.

## 🚀 Getting Started

### 1. Run with Air (Hot Reload)

```bash
air
```

### 2. Run Tests

```bash
go test ./...
```
