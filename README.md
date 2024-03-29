# Swiggy Technical Assignment

## Problem Statement:

Design a multiplayer card game that supports multiple players (up to 4) and different types of cards (e.g. number cards, action cards, etc.). The game should follow the following rules:

- Each player starts with a hand of 5 cards.

- The game starts with a deck of 52 cards ( a standard deck of playing cards).

- Players take turns playing cards from their hand, following a set of rules that define what cards can be played when.

- A player can only play a card if it matches either the suit or the rank of the top card on the discard pile.

- If a player cannot play a card, they must draw a card from the draw pile. If the draw pile is empty, the game ends in a draw and no player is declared a winner..

- The game ends when one player runs out of cardswho is declared the winner.

- BONUS: Aces, Kings, Queens and Jack are action cards. When one of these is played the following actions occur:
  - Ace(A): Skip the next player in turn
  - Kings(K): Reverse the sequence of who plays next
  - Queens(Q): +2
  - Jacks(J): +4

**NOTE**: actions are not stackable i.e. if Q is played by player 1 then player two draws two cards and cannot play a Q from his hand on that turn even if available

## Setup

- Install go

```
brew install go
```

- Clone the repo

**Username:** `ahampriyanshu`

**Fine Grained PAT:** `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`

```
git clone https://github.com/ahampriyanshu/swiggy-technical-assignment.git cardgame
cd cardgame
```

- Run the application

```
bash build.sh
```

## Preview

- ![fdsf](./screenrecord.gif)

## Evaluation Criteria

- [x] Simple design: Does the code have a clear and simple design? Is it easy to understand and modify?
- [x] Readability: Is the code well-organized and easy to read? Are the naming conventions clear and consistent?
- [x] Modelling: Are the objects and classes used in the code well-designed and appropriate for the problem at hand?
- [x] Maintainability: Is the code easy to maintain and modify? Are there any potential areas of concern or technical debt?
