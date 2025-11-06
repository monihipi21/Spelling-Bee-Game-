# Speeling Bee Game
## Monika Szczepańska CO.SDH3-B

## Overview

This project is a distributed, text-based implementation of the New York Times Spelling Bee game written in Go.  
It consists of three main components:
- A **client** that handles user interaction.
- A **server** that validates submitted words and manages scoring.
- A **dictionary** component that provides fast word lookups using a JSON dataset.

The client communicates with the server via gRPC, making this a distributed system (though all parts can run locally).  
The game logic follows the assignment specification:
- The player receives 7 random letters.
- One of them is the centre letter (which must appear in all words).
- Words must be at least 4 letters long.
- Pangrams (using all 7 letters) receive a 7-point bonus.
- Scoring:
    - A 4-letter word => 1 point
    - Words longer than 4 letters => word length points


## Design Patterns Used

The project implements three design patterns, chosen because they naturally fit the problem and simplify the structure of a distributed word game.

### 1. Singleton Pattern
**File:** `internal/directory/dictionary.go`

**Why:**  
This pattern ensures that the directory (loaded from .json file) is initialized only once and shared across the whole application.
Since the directory file is large, using a single shared instance improves performance and prevents repeated file reads.
The pattern is implemented using `sync.Once` and the `GetInstance()` method, which initializes the directory when first accessed.


### 2. Strategy Pattern

**File:** `internal/game/score_strategy.go`

**Why:**
This pattern is used to make the scoring system flexible.
Different scoring stategies can be applied withour modifying the core game logic.
The game uses the `ScoringStrategy` interface inside the `Game.AddScore()` method to calculate points.


### 3. Decorator Pattern

**File:** `internal/game/decorator.go`

**Why:**
The Decorator pattern allows adding extra behavior to the scoring process without modifying the base BasicStrategy class.
This makes the code more flexible. 
The `LoggingDecorator` wraps the main `Game` instance and logs method calls such as word validation and score updates.


## How to Run

**Download dependencies:**

` go mod download`

**Start the server:**

`go run cmd/server/main.go`

**Start the client:**

`go run cmd/client/main.go`

**Regenerate gRPC files only if needed:**

`protoc --go_out=. --go-grpc_out=. api/spellingbee/v1/spellingbee.proto`
