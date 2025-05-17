# sales_analysis_service

## 🔧 Technologies Used

- **Language:** Go (Golang)
- **Web Framework:** Gorilla Mux
- **ORM:** GORM

---

## ▶️ How to Run the Project

### 🔧 Prerequisites

- **Go** version 1.18 or above
- A running **MySQL** instance
- Ensure `dbconfig.toml` and `serviceconfig.toml` exist in the expected locations

### **install Required dependencies**

```bash
go mod tidy
```

### **Run the application**:

```bash
go run main.go
```

---

## Implemented Features

### Error Handling

Graceful management of potential errors, including input validation, JSON decoding, and database operation failures.

### Periodic Refresh

- A background job reads data from `salesData.csv` and refreshes the database based on a configurable interval.
- Frequency is set in a TOML config file (`./toml/serviceconfig.toml`).
- Default fallback is 6 hours if not configured or invalid.
- Supports overwriting or appending to the existing data, while managing duplicates.

### Normalized Database Schema

- Separate tables for `OrderDetails`, `ProductDetails`, and `CustomerDetails`.
- Foreign key constraints ensure data integrity.

---

## Endpoint

**POST** `/getSalesRevenue`

---

## Description

Fetches sales revenue data from the backend system. The data returned depends on the type of request:

- Total revenue
- Revenue by product
- Revenue by category
- Revenue by region

---

## Request

### JSON Body

```json
{
  "reqType": "REV01",
  "startDate": "2024-01-01",
  "endDate": "2024-01-31"
}
```

## Response

### Success Response Example

```json
{
  "totalRevenue": 15230.5,
  "productRevenueArr": [
    {
      "productName": "iPhone 15",
      "productRevenue": 7200.0
    }
  ],
  "categoryRevenueArr": [
    {
      "categoryName": "Electronics",
      "categoryRevenue": 8500.0
    }
  ],
  "regionRevenueArr": [
    {
      "regionName": "North America",
      "regionalRevenue": 3200.5
    }
  ],
  "status": "SUCCESS"
}
```

## Error Response

Returned when a request fails due to validation, decoding, or processing errors.

### JSON Structure

```json
{
  "status": "ERROR",
  "errorCode": "RGRA01",
  "errorMessage": "Invalid request payload"
}
```

## `/refreshData`

### Description

Triggers a background process to periodically refresh data from the `salesData.csv` file based on the configured interval.

### Request

**Method:** `POST`
**Endpoint:** `/refreshData`
**Payload:** -

### Behavior

- Reads frequency from `./toml/serviceconfig.toml` (key: `FREQUENCY`)
- If invalid or missing, defaults to **6 hours**
- Internally calls `GetCsvData()` to update the database
- Runs in a continuous loop in a background goroutine
