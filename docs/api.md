# API Documentation

## Overview

This document provides information on the available API endpoints.

### Base URL

The API is accessible at: `http://localhost:8080`

### Endpoints

#### 1. **Process Input Endpoint**

- **URL**: `/process`
- **Method**: `POST`
- **Description**: Processes a list of input strings and returns the calculated results for each input.

##### Request

- **Body**: JSON array of objects, each containing an `input` field with the string with required values.
  **Example Request Body**:
  ```
  [
    {
      "input": "1000 1300 1500 4 0 9 9 17 17 22 22 23"
    },
    {
      "input": "1300 1500 1700 7 8 19 9 20 10 21 11 22 0 23 20 22 0 21"
    }
  ]
  ```

##### Response

- **Success Response**:
  - **Code**: `200 OK`
  - **Content**: JSON object containing the results of the processed inputs.
    **Example Success Response**:
  ```
  {
    "results": [29500, 130000]
  }
  ```
  Where each item of results is an integer value of results of the processed inputs.
- **Error Responses**:
  - **Method Not Allowed**:
    - **Code**: `405`
    - **Content**: `"Method Not Allowed"`
  - **Invalid Input**:
    - **Code**: `400`
    - **Content**: `"Invalid input"`
  - **Empty Input**:
    - **Code**: `400`
    - **Content**: `"Empty input"`
  - **Processing Error**:
    - **Code**: `400`
    - **Content**: _Error message_

### Usage Example

To call the `/process` endpoint, you can use a tool like `curl` or Postman.
Here is an example using `curl`:

```
curl -X POST http://localhost:8080/process -H "Content-Type: application/json" -d '[
    {"input": "1000 1300 1500 4 0 9 9 17 17 22 22 23"},
    {"input": "1300 1500 1700 7 8 19 9 20 10 21 11 22 0 23 20 22 0 21"}
    ]'
```
