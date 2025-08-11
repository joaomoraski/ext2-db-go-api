# EXT2-DB: A "mini-db" using EXT2 with a Go API

This project combines a Computer Science course assignment that aimed to create an **interactive shell** for running commands on a `.iso` or `.img` disk image using the EXT2 filesystem, with a Go API, leveraging these developed operations as a "database."

[Link to the EXT2 project](https://github.com/joaomoraski/ext2-fs-tools/tree/master) (switch to the `db-engine` branch for this project)

## Architecture

The system consists of two main components that communicate via standard I/O (`stdin`/`stdout`):

1. **Database (`ext2-db-engine`):**  
   A C command-line tool compiled from the original project, which directly manipulates an EXT2 disk image.  
   It handles all low-level operations, such as inode allocation, bitmap manipulation, and block read/write.

2. **Go API:**  
   A simple RESTful API that exposes endpoints for **Create** and **List (with filters)** operations on a "user database."  
   It processes HTTP requests, validates data, and calls the C "database" to persist or retrieve information.

## How to Run the Project

### Prerequisites

* GCC compiler
* `readline` library (`sudo apt-get install libreadline-dev` or equivalent)
* Go (version 1.18 or higher)

### Compilation and Execution

The project uses a `Makefile` in the root of the `ext2-fs-tools` submodule to compile the C engine and a `Makefile` in the `ext2-db-go-api` root to orchestrate everything.

1. **Compile the C engine:**
   ```bash
   cd ext2-db-engine
   make all
   cd ..

1.1 **Generate a new image**

```bash
make generate-ext2
```

1.2 **Creating the table**

```bash
make run
```

Inside the terminal that opens, create the database with the command:

```bash
touch user_record
```

2. **Start the Go API:**
   In the project root, run:

   ```bash
   cd api
   go run .
   ```

   The server will be running at `http://localhost:8080`.

---

## How to Use the API

The API exposes the following endpoints for the `User` entity:

### Create a New User

* **Endpoint:** `POST /users`
* **Description:** Creates a new user record in the `/user_record` file inside the EXT2 image.
* **Request Body (JSON):**

  ```json
  {
      "id": 1,
      "is_active": 1,
      "username": "moraski",
      "email": "moraski@gmail.com"
  }
  ```

### Fetch Users

* **Endpoint:** `GET /users`
* **Description:** Returns a list of all registered users.
* **Query Parameters:**
    * `limit=<number>`: Limits the number of results.
    * `filters=<conditions>`: (Optional) Filters results. Conditions must follow the pattern `field:operator:value`.
        * Supports `=`, `%`
