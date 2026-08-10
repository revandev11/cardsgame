# Deck of Cards 

A lightweight and clean **Go (Golang)** application that manages a deck of playing cards, supporting actions like generating decks, shuffling, splitting, and saving/loading from the file system.

---

## 🚀 Features

* **Create New Deck**: Generates a standard set of cards combination based on suits (Spades, Hearts, Diamonds, Clubs) and values.
* **Shuffle (`shuffle`)**: Randomizes the order of cards in the deck.
* **File I/O (`saveToFile` / `newDeckFromFile`)**: Persists the current deck state to a local file and restores it.
* **Deal Cards (`deal`)**: Splits a deck into a hand and remaining cards based on a given hand size.
* **Unit Testing**: Fully covered with unit tests to ensure core logic reliability.

---

## 🛠️ Getting Started

### Prerequisites
* **Go** (version 1.20 or higher)

### Running the Application

1. Clone or download the repository:
   ```bash
   git clone [https://github.com/your-username/cards.git](https://github.com/your-username/cards.git)
   cd cards
