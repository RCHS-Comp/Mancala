# Mancala
A library to help put mancala wherever you want it.

## Usage
**Before use, be sure to reset the board!**

```
var Game Board    // Create a new board
Game.Reset()      // Reset the board
```

### Moving

*(When playing, the left side will go first)*

To make a move, just tell what piece you want to move

```
Game.Move(5)      // Moves the 5th piece of the board
Game.Move(6)      // Moves the 6th piece of the board
// Game.Move(0)   // This is an invalid move, it will just return without doing anything
// Game.Move(7)   // Also an invalid move, it will also do nothing
```

### Scoring

To get the score of each side, just grab it from the side you want

```
Game.Right.Points  // Has the right's points
Game.Left.Points   // Has the left's points
```

### Pieces

The pieces are represented as an array with the length of 6

For an array that looks like this:

```
1 2 3 4 5 6 (P1)
1 2 3 4 5 6 (P2)
```

The board will look like this:

```
(P1)  (P2)
1        6
2        5
3        4
4        3
5        2
6        1
```

To get each side, just get the pieces from the side

```
Game.Left.Pieces   // Might look like (1, 2, 3, 4, 5, 6)
Game.Right.Pieces  // Might look like (4, 4, 4, 4, 4, 4)
```

**Note: the arrays are [6]int and not []int**

If you want to print the board, you could try something like this

```
// Be sure to import "fmt" before using this
for i := 0; i < 6; i ++ {    // Repeats 6 times, going from 0 to 5
    fmt.Println(Game.Left.Pieces[i], Game.Right.Pieces[5 - i])
}
```

### Current player

To get the current player, get the current side of the board

```
Game.Side          // True is right, False is left
```
