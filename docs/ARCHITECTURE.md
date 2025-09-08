# Project Architecture: PrivatPhoto

## 1. High-Level Overview

PrivatPhoto is a cross-platform solution for syncing media files between a mobile device (Android) and a user's personal server. The architecture is built on a classic client-server model with a focus on security and privacy. User data never leaves the user's controlled infrastructure.
+-------------------+ HTTPS (TLS 1.3) +-----------------------+
| | <-------------------> | |
| Android Client | | Go Backend Server |
| (Java/Kotlin) | API Calls | (Chi Router) |
| | | |
+-------------------+ +-----------------------+
| |
| Sync Logic | Persistence
| |
+-------------------+ +-----------------------+
| Android Media | | SQLite Database |
| Store | | |
+-------------------+ +-----------------------+

## 2. Technology Stack

### Backend (Server)
*   **Language:** Go (Golang) 1.21+
*   **Web Framework:** [Chi](https://github.com/go-chi/chi) (lightweight, idiomatic HTTP router)
*   **Database:** SQLite (using [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) driver)
*   **Authentication:** JWT (JSON Web Tokens) using [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
*   **Password Hashing:** `golang.org/x/crypto/bcrypt`
*   **TLS:** Standard library `crypto/tls`

### Client (Android)
*   **Language:** Java (with possible future migration to Kotlin)
*   **Network Client:** [Retrofit 2](https://square.github.io/retrofit/) with [Gson](https://github.com/google/gson) for JSON (de)serialization
*   **HTTP Client:** [OkHttp](https://square.github.io/okhttp/) (underlying Retrofit)
*   **Image Loading:** [Glide](https://github.com/bumptech/glide) or [Coil](https://coil-kt.github.io/coil/) (for gallery display)

### General
*   **Version Control:** Git / GitHub
*   **CI/CD:** GitHub Actions (to be set up later for automated tests and builds)
*   **Communication Protocol:** HTTPS/1.1

## 3. Data Model (SQLite Schema)

### Table `users`
| Field           | Type          | Description                                                                 |
|-----------------|---------------|-----------------------------------------------------------------------------|
| `id`            | INTEGER PK AI | Unique user identifier                                                     |
| `login`         | TEXT UNIQUE   | Unique login name (e.g., email or username)                                |
| `password_hash` | TEXT NOT NULL | Password hash created with bcrypt (min cost: 12)                           |
| `email`         | TEXT          | User email (optional, for notifications and 2FA)                           |
| `two_fa_enabled`| BOOLEAN       | Flag indicating if two-factor authentication is enabled for this user      |
| `two_fa_secret` | TEXT          | Secret for generating TOTP codes (if 2FA is enabled)                       |
| `created_at`    | DATETIME      | Account creation date and time (DEFAULT CURRENT_TIMESTAMP)                 |

### Table `files`
| Field        | Type          | Description                                                      |
|--------------|---------------|------------------------------------------------------------------|
| `id`         | INTEGER PK AI | Unique file identifier                                          |
| `user_id`    | INTEGER FK    | Reference to the file owner (`users.id`)                        |
| `filename`   | TEXT NOT NULL | Original filename (e.g., `IMG_20231001.jpg`)                   |
| `size`       | INTEGER       | File size in bytes                                              |
| `upload_time`| DATETIME      | Upload date and time (DEFAULT CURRENT_TIMESTAMP)               |
| `local_path` | TEXT          | Path to the file on the server (relative to the uploads root)   |

*Foreign Key: `FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE`*

## 4. API Endpoints Specification

**Base URL:** `https://<server-ip>:<port>/api`

### Authentication

| Method | Endpoint           | Description                                                                 | Request Body                                                              | Response                                                                 |
|--------|--------------------|-----------------------------------------------------------------------------|---------------------------------------------------------------------------|--------------------------------------------------------------------------|
| `POST` | `/auth/login`      | First stage of authentication. Always returns `requires2fa` flag.          | `{"login": "string", "password": "string"}`                              | `{"success": bool, "requires2fa": bool, "token": "string"}`             |
| `POST` | `/auth/verify-2fa` | Second stage of authentication (if required). Issues a final sessionToken. | `{"token": "string (from login)", "code": "string (6-digit)"}`           | `{"success": bool, "sessionToken": "string"}`                           |

### Files

All endpoints below require the header: `Authorization: Bearer <sessionToken>`

| Method | Endpoint                 | Description                                 | Request Body                     | Response                                                           |
|--------|--------------------------|---------------------------------------------|----------------------------------|--------------------------------------------------------------------|
| `GET`  | `/files`                 | Get a list of all user files.               | -                                | `[{"id": int, "filename": "string", "size": int, "upload_time": "string"}]` |
| `POST` | `/files/upload`          | Upload a new file.                          | `multipart/form-data` (field `file`) | `{"success": bool, "fileId": int}`                                |
| `GET`  | `/files/download/{id}`   | Download a file by its ID.                  | -                                | `Binary File Data`                                                |
| `DELETE`| `/files/{id}`           | Delete a file by its ID.                    | -                                | `{"success": bool}`                                               |

## 5. Network & Deployment

*   **Protocol:** HTTPS (TLS 1.3) for encrypting all traffic.
*   **Ports:** The server by default listens on port `8443` (can be changed in configuration).
*   **Network Interaction:** The first stage uses direct TCP connection. The user must set up port forwarding on their router from an external port to `localhost:8443` of their server PC.
*   **Domain Names:** The first stage assumes the use of an IP address. In the future, support for DynDNS may be added for users with dynamic IPs.

## 6. Security Considerations

*   **Authentication:** Multi-stage (login/password + 2FA if necessary).
*   **Passwords:** Hashed using `bcrypt` (cost ≥ 12).
*   **Transport:** All traffic is encrypted with TLS.
*   **Secrets:** 2FA secrets and JWT signing keys are stored on the server in encrypted form or in secure environment variables.
*   **Sessions:** Short-lived JWT tokens (sessionToken) are used for API access.

---

*This is a living document. The architecture may change as development progresses and new information is obtained. Last update: 08.09.2025*
