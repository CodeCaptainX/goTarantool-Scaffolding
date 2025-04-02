# GoTarantool - Project Generator

## Overview

GoTarantool is a project generator that simplifies creating Go projects with a predefined template. This guide provides setup instructions and usage details.

## Setup Instructions

### 1. Navigate to the Template Directory

Ensure you are in the correct directory where the template files are stored:

```sh
cd /path/to/YourtemplateInThisProject

```

### 2. Update `createProjectStructure`

Modify the `templatePath` variable to properly reference the template directory:

```go
templatePath := filepath.Joifdn("/path/to/YourtemplateInThisProject", strings.TrimPrefix(v, "template/"))
```

This ensures correct path resolution for templates during project generation.

### 3. Build the Project

Compile the Go program to generate an executable:

```sh
go build -o goTarantool main.go
```

### 4. Move the Executable to a Global Path

Move the built executable to `/usr/local/bin/` for global access:

```sh
sudo mv goTarantool /usr/local/bin/
```

### 5. Usage

Once installed, you can generate a new project by running:

```sh
goTarantool anyProjectName
```

This will create a project using the specified name with the predefined template.

---

### 🚀 Enjoy Coding!

Now you can generate Go projects effortlessly. Happy coding! 🎯
