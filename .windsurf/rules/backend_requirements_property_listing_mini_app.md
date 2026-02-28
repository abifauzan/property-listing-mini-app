---
description: At the first of every model decision, the AI assistant must consider this document.
---
# Backend Requirements (English) — Property Listing Mini App (DANA Take‑Home)

> Purpose: This document is a **backend-only** requirements translation from the provided assessment PDF.  
> It is intended to be fed into an AI coding agent (e.g., Windsurf) as a single source of truth.

---

## 1. Goal

You must build a backend API that serves property listing data for a simplified **“Property Finder”** mini app. The assessment requires both a mini-program frontend and a backend REST API to serve the data. fileciteturn2file0L3-L6

---

## 2. Allowed Backend Options (Technology & Interface)

### 2.1 Language / Runtime
- Backend MUST be implemented using **GoLang (preferred)** or **Node.js**. fileciteturn2file0L14-L16

### 2.2 API Style
- Backend MUST expose either:
  - a **RESTful** API, **or**
  - a **GraphQL** API  
  to support the frontend. fileciteturn2file0L45-L47

### 2.3 Data Storage
- Backend MAY use a persistent database (**PostgreSQL**, **MongoDB**, **MySQL**), **or**
- Backend MAY use a simple **in-memory data store / JSON file** for this assessment. fileciteturn2file0L14-L16

> For this implementation, the target is **JSON file / in-memory** storage, while keeping requirements compatible with future migration to a persistent DB.

---

## 3. Backend Functional Requirements

### 3.1 Provide Property Listing Data (for the Listing Screen)
- Backend MUST support fetching property data for the listing screen (“Data Fetching: Fetch property data from your Backend API”). fileciteturn2file0L34-L36

### 3.2 Support Search by Title (Listing Screen Search Bar)
- The listing screen includes a **Search Bar** that filters properties by **title** (e.g., typing `"sam"` shows items containing `"sample"`). fileciteturn2file0L29-L31
- Backend MUST provide data that enables this title-based filtering to work end-to-end. fileciteturn2file0L29-L31

> Notes (requirement-level, not implementation):  
> The PDF specifies the UI behavior and does not mandate whether filtering is server-side or client-side. The backend must ensure the data returned makes this behavior possible.

### 3.3 Provide Property Detail by Property ID
- When a user clicks a card in the listing screen, the app navigates to the detail screen and passes the **Property ID**. fileciteturn2file1L8-L10
- Backend MUST support retrieving and returning the **full property detail** for a given Property ID. fileciteturn2file1L8-L11

### 3.4 “Book Now” Does Not Require Backend Processing
- The detail screen contains a fixed **“Book Now”** button that **can be a dummy action** (toast/alert). fileciteturn2file1L12-L13
- Therefore, backend is NOT required to implement booking creation or booking workflows. fileciteturn2file1L12-L13

---

## 4. Data Contract Requirements (Schema / JSON Structure)

The PDF provides a JSON example and states to **use the structure as a baseline**. fileciteturn2file1L14-L18

### 4.1 Top-Level Response Envelope
- Responses MUST follow the example baseline with a root object containing:
  - `data` object
  - `data.propertyListings` array  
  fileciteturn2file1L20-L23

### 4.2 Property Listing Item (minimum fields used on Listing Screen)

Each item in `data.propertyListings[]` MUST include at least: fileciteturn2file1L20-L32
- `documentId` (string) — the Property ID
- `Banner.url` (string) — listing thumbnail/cover image URL
- `Title` (string)
- `Price` (string) — numeric string (e.g., `"750000"`)
- `createdAt` (string) — ISO timestamp string

UI implication (for correctness of client display):
- Price is displayed as **“Rp X.XXX.XXX”** in the listing screen. The backend MUST provide `Price` in a format that can be formatted into this display. fileciteturn2file0L35-L36

### 4.3 Property Detail Item (fields used on Detail Screen)

The detail screen must display: a large hero image, title, formatted price, gallery carousel, and description text. fileciteturn2file1L10-L11

A property detail object MUST support the following fields (based on the provided example object): fileciteturn2file3L23-L60 fileciteturn2file4L24-L40
- `documentId` (string)
- `Title` (string)
- `Price` (string) — numeric string
- `Description` (string)
- `Banner.url` (string) — hero image URL
- `Images` (array of objects), where each item includes:
  - `url` (string)
- `Facilities` (array of strings)
- `Terms` (string)
- `Conditions` (string)

---

## 5. Quality / Evaluation Criteria Related to Backend

### 5.1 Code Quality
- The submission will be evaluated on **clean, readable, maintainable** code structure (including backend). fileciteturn2file2L19-L20

### 5.2 Bonus: Backend Unit Tests
- Bonus points include **unit tests for backend logic**. fileciteturn2file2L23-L26

---

## 6. Submission Requirements That Affect Backend Deliverables

The repository MUST include a `README.md` with:
- Instructions on **how to run the Backend**
- Instructions on how to run the Mini Program
- Screenshots or screen recording of the running application  
fileciteturn2file2L28-L33

The repository link must be sent **within five (5) days** after receiving the email. fileciteturn2file2L33-L34

---

## 7. Out of Scope for Backend (Explicit / Implied)

- Backend does NOT need to implement booking (the “Book Now” action may be dummy on the client). fileciteturn2file1L12-L13
- The PDF does not require authentication/authorization, user management, payments, or admin CRUD tools (not mentioned in functional requirements). fileciteturn2file0L17-L44

---

## 8. Acceptance Checklist (Backend-Only)

A backend implementation is considered compliant if:

### MUST
- Exposes an API (RESTful or GraphQL) to support the frontend. fileciteturn2file0L45-L47
- Supports fetching property listings from the backend API. fileciteturn2file0L34-L36
- Supports retrieving property detail by Property ID (`documentId`). fileciteturn2file1L8-L11
- Uses the baseline JSON structure `data.propertyListings[]` and the required fields listed above. fileciteturn2file1L20-L32

### SHOULD
- Ensure the returned data supports title-based filtering (search bar behavior). fileciteturn2file0L29-L31

### MAY (bonus / optional)
- Include unit tests for backend logic. fileciteturn2file2L23-L26
