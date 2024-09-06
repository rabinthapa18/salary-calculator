# Test Cases

This folder contains the function to test all the test cases in txt file to ceck functionality of the salary calculator application.

## Files

- `main.go`: This file contains the main logic for reading and processing test cases.
- `test_cases.txt`: This file contains the test cases in a comma-separated format where each line represents a test case with an input and the expected output.

## Usage

To run the test cases, execute the `main.go` file. It will read the test cases from `test_cases.txt`, process each test case, and print the results.

### Running the Tests

1. Ensure you have Go installed on your system.
2. Navigate to the root folder of the project.
3. Run the following command:

   ```sh
   go run testCases/main.go
   ```

### Test Case Format

Each line in `test_cases.txt` should be in the following format:

```text
<input>,<expected_output>
```

- `<input>`: The input string to be processed.
- `<expected_output>`: The expected integer output.

> Note: There should be no spaces between the input, comma, and expected output. two examples are given in the txt file.
