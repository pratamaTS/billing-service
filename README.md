# Billing Engine Service

This project implements a simple billing engine for a loan system using Golang and the Gin framework.

---

## Assumptions

- Loan duration is fixed at 50 weeks with a flat 10% interest rate.
- Borrowers can only pay one week's installment at a time.
- Payments must follow the schedule (sequential, no skipping ahead).
- Data is stored in-memory using a global map protected by `sync.Mutex`.

---

## How to Run

```bash
go mod tidy
go run main.go
