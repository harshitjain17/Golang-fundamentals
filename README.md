# Golang Fundamentals

Welcome to the **Golang Fundamentals** repository! This project is a hands-on exploration of Go programming through a series of exercises and mini-projects. It serves as a practical guide for learning key features of the Go language such as data structures, interfaces, concurrency, and more.


# 📦 Projects & Topics Covered

### 🃏 Cards Project

A command-line Go application that simulates operations on a standard 52-card deck. This was the foundational project used to practice Go fundamentals like custom types, slices, functions, file operations, and basic testing.

**Features:**
- Create and print a 52-card deck
- Deal a hand and view remaining cards
- Save/load the deck from a file
- Shuffle using Fisher-Yates algorithm
- Unit tests to ensure correctness

**Key Files:**
- `deck.go`: Deck logic (create, shuffle, save/load)&#8203;:contentReference[oaicite:0]{index=0}
- `deck_test.go`: Unit tests&#8203;:contentReference[oaicite:1]{index=1}
- `main.go`: Demo of full flow&#8203;:contentReference[oaicite:2]{index=2}

**Run Instructions:**
```bash
go run main.go deck.go
go test
```
---

### 🕸️ Goroutines & Channels

Learned how to use goroutines and channels to build concurrent Go programs. A sample program was created to check the status of multiple websites in parallel, with repeated checks at intervals using goroutines.

**Concepts Practiced:**
- Launching concurrent functions with `go` keyword
- Communicating between goroutines via channels
- Infinite loops with time delays for re-checking URLs

**Key File:**
- `main.go` (website checker)

---

### 🔁 Interfaces

Understood and implemented interfaces in Go using two main examples:

**1. Greeting Bots Example**  
Demonstrates polymorphism using a `bot` interface with different implementations: `englishBot` and `spanishBot`.

**2. Shape Area Calculator**  
Demonstrates how structs like `triangle` and `square` implement a `shape` interface to calculate area differently.

---

### 🧱 Structs & Pointers

Built a `person` struct with nested `contactInfo`, showcasing:

- Value vs reference types  
- Pointer receivers  
- Updating struct fields  
- Behavior of slices (reference types)

**Key File:**
- `main.go` (person struct and pointer methods)

---

### 🗺️ Maps

Practiced creating, modifying, and iterating over Go maps. Also handled:

- Checking if a key exists  
- Deleting key-value pairs  
- Iterating through map entries

**Key File:**
- `main.go` (map operations)

---

### 🧠 Summary

This repo serves as a comprehensive journey through Go's core features, blending theory with practical implementation. From creating card games to handling concurrency with goroutines and channels, it provides a strong foundation for building robust applications in Go.
