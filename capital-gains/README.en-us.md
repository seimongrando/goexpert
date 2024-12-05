# Capital Gains Tax Calculator

This project calculates capital gains taxes based on buy and sell stock operations. It processes operations provided in JSON format, applies rules for profit, loss, and tax exemptions, and returns the computed tax values.

---

## Technical and Architectural Decisions

1. **Code Organization**:
   - The project is divided into three main packages:
      - `logic`: Contains the core logic for capital gains calculations.
      - `console`: Manages user interaction and input/output operations.
      - `integration`: Contains integration tests.

2. **Referential Transparency**:
   - Auxiliary functions in the `logic` package are designed to be pure, ensuring predictability, reusability, and ease of testing.

3. **Simplicity**:
   - The code is kept simple and modular, favoring readability and maintainability.

4. **Separation of Concerns**:
   - Each layer of the program has a clear responsibility:
      - Processing logic (`logic`).
      - User communication (`console`).
      - Integration tests (`integration`).

---

## Justification for Frameworks or Libraries

1. **`mockgen`**:
   - Used to create mocks for the `processor` interface, simplifying unit tests and ensuring coverage for various scenarios.

2. **`testify`**:
   - Used to enhance test assertions, making them more readable and reducing redundant code.

These libraries were chosen for their wide usage, extensive documentation, and compatibility with Go projects of any size.

---

## Instructions to Compile and Execute

1. Clone the repository:
   ```bash
   git clone https://github.nubank.com/capital-gains
   cd capital-gains
   ```

2. Compile and run the program:
   ```bash
   go run main.go
   ```

3. Provide input in JSON format directly in the console:
   ```json
   [{"operation":"buy", "unit-cost":10.00, "quantity":10000}, {"operation":"sell", "unit-cost":50.00, "quantity":5000}]
   ```

4. Press **Enter** twice to process.

---

## Instructions to Generate the Executable

1. To generate the executable:
```bash
go build -o capital-gains
```
This will create a binary named capital-gains in the current directory.

2. Run the program
```bash
./capital-gains < input.txt
```

---

## Instructions to Run Tests

1. Run all tests (unit and integration):
   ```bash
   go test ./...
   ```

2. To run only integration tests:
   ```bash
   go test ./integration
   ```

---

## Additional Notes

- **Code Quality**:
   - The code is modular and follows best practices, such as separation of concerns and clear implementation.

- **Extensibility**:
   - It is easy to add new features or adapt calculations to specific rules.

- **Error Messages**:
   - All error messages are clear and user-oriented, helping debug any issues.

---
