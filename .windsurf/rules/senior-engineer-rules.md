---
trigger: always_on
---
# Senior Engineer AI Assistant Rules
A set of core rules to ensure the AI assistant generates system design, code, and tests that meet principal/senior engineer standards for scalability and best practices.

## 1. System Design
- **Scalability First**: Design architectures capable of handling growth in traffic, data, and complexity (e.g., stateless services, caching layers, asynchronous processing).
- **Separation of Concerns**: Apply Clean Architecture, Hexagonal Architecture, or Domain-Driven Design (DDD) principles to decouple business logic from infrastructure and delivery mechanisms.
- **Modularity**: Design well-bounded domains and modular backends.
- **API Design**: Adhere to RESTful or gRPC best practices. Ensure APIs are versioned, idempotent, clearly documented, and secure.

## 2. Go Backend Engineering
- **Idiomatic Go**: Follow Go conventions (Effective Go, standard project layout). Avoid over-engineering; leverage the standard library effectively.
- **Concurrency**: Utilize goroutines and channels safely. Prevent race conditions and goroutine leaks by managing contexts appropriately.
- **Error Handling**: Implement robust, structured error handling. Wrap errors with meaningful context (`fmt.Errorf("...: %w", err)`) and avoid silent failures.
- **Interfaces & DI**: Use small, focused interfaces and Dependency Injection to make the code highly testable and decoupled.
- **Performance**: Write allocation-conscious code and avoid premature optimization, but design for performance from the ground up.

## 3. Mini-Program Frontend Engineering
- **Component-Based Architecture**: Build reusable, isolated, and stateless UI components.
- **State Management**: Manage global and local state efficiently to minimize re-renders and optimize memory footprint.
- **Performance**: Optimize asset loading, implement lazy-loading where appropriate, and ensure a fluid user experience.
- **Robust UI**: Always implement loading states, skeleton screens, and graceful error handling for network or processing failures.

## 4. Testing & Quality Assurance
- **Testable Code**: Ensure all logic is easily testable without requiring heavy external dependencies.
- **Comprehensive Coverage**: Write meaningful unit tests, integration tests, and table-driven tests (for Go). Use appropriate mocking frameworks (e.g., `gomock`, `testify`).
- **Edge Cases**: Explicitly handle and test edge cases, boundary conditions, nil pointers, and concurrent access scenarios.

## 5. General Best Practices
- **Security**: Validate and sanitize all inputs. Protect against common vulnerabilities (SQLi, XSS, CSRF).
- **Observability**: Incorporate structured logging, metrics, and tracing considerations into the code structure.
- **Simplicity & Maintainability**: Favor simple, highly readable code over complex, "clever" solutions (KISS principle).
- **Documentation**: Document "why" over "what" in comments, ensure exported types/functions have GoDoc comments, and maintain clear system documentation.

