# Cards - Go Project

A simple command-line Go application that simulates a standard 52-card deck. It allows for creating, shuffling, dealing, saving to a file, and loading from a file.

## Features

- Create a new 52-card deck
- Print the full deck
- Deal a hand of cards
- Save the deck to a file
- Load the deck from a file
- Shuffle the deck
- Basic unit testing for core functionalities

## Project Structure
```
├── deck.go         # Core logic for deck handling 
├── deck_test.go    # Tests for deck creation and file operations 
├── main.go         # Application entry point and usage demo 
└── go.mod          # Module definition
```

## How to Run

1. **Clone the repo**  
2. **Build and run the app**:
   ```bash
    go run main.go deck.go
3. **Run tests:**:
    ```bash
    go test

## Notes

- Deck is saved as a comma-separated string in a file.
- `shuffle()` uses a basic randomization algorithm.
- Includes error handling for file operations.

## Example Flow

- Creates and prints a full deck
- Deals a 5-card hand and shows the remaining deck
- Saves the deck to `my_cards.txt`
- Loads the deck back from the file
- Shuffles and reprints the deck