package Mancala
// Todo: add different modes & better errors
// Maybe clean things up?

const Default = 4

type Side struct {
    Pieces [6]int
    Points int
}
 
type Board struct {
    Side bool   // True = right
    Left Side
    Right Side
    Mode bool   // True = avalanche
}

func (Current Side) IsEmpty() (bool) {
    return (Current.Pieces == [6]int{0, 0, 0, 0, 0, 0})
}

func (Current *Board) Reset() {
    // Dude, I love this song playing
    // Jazz is just so good
    // It's "We had a ball" - Gazzara

    for i := 0; i < 6; i ++ {   // Easier patching in the future
        Current.Left.Pieces[i] = Default
        Current.Right.Pieces[i] = Default
    }

    Current.Left.Points = 0
    Current.Right.Points = 0
    Current.Side = false
}

func (Current *Board) EndMove() {   // Make this easier because why not?
    Current.Side = !(Current.Side)  // If we need to steal be sure to invert first
}


func (Current *Board) Move(Position int) {
    if (Position < 1 || Position > 6) {
        return
    } else {
        Position --
    }

    if Current.Side {
        Position = (5 - Position)
    }

    if (Current.Left.Pieces[Position] == 0 && !(Current.Side)) {
        return
    } else if (Current.Right.Pieces[Position] == 0 && Current.Side) {
        return
    }

    var ModifySide *[6]int
    var Opposite *[6]int
    var Points *int

    if Current.Side {
        ModifySide = &Current.Right.Pieces
        Points = &Current.Right.Points
        Opposite = &Current.Left.Pieces
    } else {
        ModifySide = &Current.Left.Pieces
        Points = &Current.Left.Points
        Opposite = &Current.Right.Pieces
    }

    UntilMoveEnd := ModifySide[Position]
    CurrentPosition := Position
    OtherSide := false

    ModifySide[Position] = 0

    for UntilMoveEnd > 0 {
        CurrentPosition ++

        if !(OtherSide && CurrentPosition == 6) {
            UntilMoveEnd --
        }

        if OtherSide {
            if CurrentPosition == 6 {
                CurrentPosition = -1
                OtherSide = false
            } else {
                Opposite[CurrentPosition] ++
            }
        } else {
            if CurrentPosition == 6 {
                *Points ++
                CurrentPosition = -1
                OtherSide = true
                if UntilMoveEnd <= 0 {
                    // Free move
                    Current.Side = !(Current.Side)    // Inverts it to be inverted again
                }
            } else {
                ModifySide[CurrentPosition] ++
            }
        }

    }

    if !(OtherSide) {
        if ModifySide[CurrentPosition] == 1 {
            ModifySide[CurrentPosition] = 0
            *Points ++
            *Points += Opposite[5 - CurrentPosition]
            Opposite[5 - CurrentPosition] = 0
            // Steal
        }
    }

    EndMove()
}
