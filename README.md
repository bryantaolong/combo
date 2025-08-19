<a href="https://github.com/bryantaolong/combo">
  <img width="60px" height="60px" src="./assets/logo.png" align="right"  alt=""/>
</a>

# Combo

**Combo** is a command-line To Do application developed in **Go**, supporting group management, task management, marking tasks as done, and color-coded output for easy task status distinction.

---

## ✨ Features

* Create/Delete/List groups
* Add/Delete/List/Clear tasks
* Mark tasks as done
* Support for default and custom groups
* Color-coded output to differentiate task status:

    * Completed tasks: Green
    * Pending tasks: Default color
    * Warnings/Errors: Red
    * Informational messages: Yellow or Blue

---

## 🚀 Installation & Build

### 1. Install Go

Make sure Go 1.20+ is installed and environment variables are set:

```bash
go version
```

### 2. Clone the repository

```bash
git clone https://github.com/bryantaolong/combo.git
cd combo
```

### 3. Build the executable

#### Windows

```powershell
go build -o combo.exe
```

#### Linux / macOS

```bash
go build -o combo
```

> Optional: Add the executable to your system PATH for global usage.

---

## 🎮 Usage

### Help

```bash
combo --help
combo todo --help
combo group --help
```

### Group Management

```bash
# Create a group
combo group add work

# List all groups
combo group list

# Delete a group (default group cannot be deleted)
combo group delete work
```

### Task Management

```bash
# Add tasks
combo todo add "Write daily report" -g work
combo todo add "Default task 1"

# List tasks
combo todo list -g work
combo todo list

# Mark task as done
combo todo done 1 -g work

# Delete task
combo todo delete 1 -g work

# Clear all tasks in a group
combo todo clear -g work
```

---

## Example

```powershell
# Create a group
combo group add work

# Add tasks
combo todo add "Write daily report" -g work
combo todo add "Meeting preparation" -g work

# List tasks
combo todo list -g work

# Mark task as done
combo todo done 1 -g work

# Delete task
combo todo delete 2 -g work

# Delete group
combo group delete work
```

---

## Project Structure

```
combo/
├─ cmd/           # Cobra command modules
│  ├─ root.go
│  ├─ todo/
│  │  ├─ add.go
│  │  ├─ delete.go
│  │  ├─ list.go
│  │  ├─ done.go
│  │  └─ clear.go
│  └─ group/
│     ├─ group.go
│     ├─ add.go
│     ├─ list.go
│     └─ delete.go
├─ storage/       # Data structures and storage logic
├─ main.go
└─ README.md
```

---

## 💾 Data Storage

* Tasks are stored in a JSON file located in the user directory (customizable)
* Each group contains a list of tasks
* Each task has:

    * `ID`: unique identifier
    * `Content`: task content
    * `Done`: completion status

---

## 📚 Color Output

* Implemented using [`fatih/color`](https://github.com/fatih/color)
* Completed tasks: Green
* Warnings/Errors: Red
* Informational messages: Yellow or Blue

---

## License

[MIT License](./LICENSE)
