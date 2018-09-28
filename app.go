// Template generated by reactGen

package main

import (
	"strconv"
	"myitcv.io/react"
	"myitcv.io/react/jsx"
)

// The top-level application component
type AppDef struct {
	react.ComponentDef
}

// The required build function for the top-level application component
func App() *AppElem {
	return buildAppElem()
}

// The required render function for the top-level application component
// For simplicity, I use the jsx.HTML function to generate several
// components from a single HTML string.
func (a AppDef) Render() react.Element {
	jsxArray := jsx.HTML(`
<p>This is a port of the <a href="https://reactjs.org/tutorial/tutorial.html">React Tic-Tac-Toe Tutorial</a> to 
<a href="https://github.com/gopherjs/gopherjs">GopherJS</a> using the GopherJS React bindings from 
<a href="https://github.com/myitcv">Paul Jolly</a>.</p>
<p> Note that I have performed a few of the suggested extensions mentioned at the end of the tutorial as well.  
I also have a example of how to use the JSX features of the GopherJS React bindings </p>
`)
	
	return react.Div(nil,
		react.H1(nil,
			react.S("React Tic-Tac-Toe Tutorial with GopherJS"),
		),
		react.P(nil,
			jsxArray...,
		),
		Game(),
	)
}

// Definition of the Square component
//
// Note that no state is held at the Square-level itself.
type SquareDef struct {
	react.ComponentDef
}

// Properties for the Square component
//
// position: the location of the square on the grid labeled from upper left to lower right starting with 0
//           ((0, 1, 2), (3, 4, 5), (6, 7, 8))
//
// selected: the location of the last move on the board (used for comparing with current position and
//           hilighting the Square component using CSS if the current position number matches)
//
// value: the current value of the Square ("X", "O", "")
//
// game: a reference to the upper-level GameDef component that is used for managing the click event
type SquareProps struct {
	position int
	selected int
	value    string
	game     GameDef
}

// The "build" function for the Square component
func Square(p SquareProps) *SquareElem {
	return buildSquareElem(p)
}

// The Square component's Render function.
//
// If the position and selected properties are equal, we highlight the component using CSS.
//
// The button created has its CSS class name and a "onclick" properties set.
func (sq SquareDef) Render() react.Element {
	buttonClass := "square"
	props := sq.Props()
	if props.position == props.selected {
		buttonClass = "square-selected"
	}
	return react.Button(&react.ButtonProps{ClassName: buttonClass,
		OnClick: SquareClick{props.game, props.position},
	},
		react.S(props.value),
	)
}

// The Board component definition
//
// Note that no state is held at the Board-level itself.
type BoardDef struct {
	react.ComponentDef
}

// The Board component's properties
//
// game: the higher-level GameDef option that will handle the click functionality
//
// squares: the array displayed by the board
//
// selected: the Square on the board that should be selected
type BoardProps struct {
	game     GameDef
	squares  [numSquares]string
	selected int
}

// The Board build function
func Board(bp BoardProps) *BoardElem {
	return buildBoardElem(bp)
}

// The helper function for creating an individual square for based on the Board's
// properties and an index and selected index information.
func (b BoardDef) renderSquare(index int, selected int) react.Element {
	return Square(
		SquareProps{
			position: index,
			selected: selected,
			game:     b.Props().game,
			value:    b.Props().squares[index],
		})
}

// The Board component's Render function.
//
// This creates the 3x3 grid of Square components declaratively.
func (b BoardDef) Render() react.Element {
	return react.Div(nil,
		react.Div(&react.DivProps{ClassName: "board-row"},
			b.renderSquare(0, b.Props().selected),
			b.renderSquare(1, b.Props().selected),
			b.renderSquare(2, b.Props().selected),
		),
		react.Div(&react.DivProps{ClassName: "board-row"},
			b.renderSquare(3, b.Props().selected),
			b.renderSquare(4, b.Props().selected),
			b.renderSquare(5, b.Props().selected),
		),
		react.Div(&react.DivProps{ClassName: "board-row"},
			b.renderSquare(6, b.Props().selected),
			b.renderSquare(7, b.Props().selected),
			b.renderSquare(8, b.Props().selected),
		),
	)
}

// The Game Component definition
//
// All of the state of the tic-tac-toe game is held at this level and then
// distributed to the children components through properties.
type GameDef struct {
	react.ComponentDef
}

// The maximum capacity for the slices holding the Game History data
const historyCapacity = 10
// The number of squares in the grid
const numSquares = 9
// The number of indicies (col, row) for the "move" array
const numIndicies = 2

// The History structure records the state of the board and the move selected
// for this particular turn.
type History struct {
	squares [9]string
	move    [numIndicies]int
}

// The overall game state (a GopherJS React declaration)
//
// history: Keeps a history of the moves for the game
//
// stepNumber: the step number for the current state of the board
//
// xIsNext: helps the Game know whose turn it is
type GameState struct {
	history    []History
	stepNumber int
	xIsNext    bool
}

// The build function for the Game component
func Game() *GameElem {
	return buildGameElem()
}

// The function compares the given state of the board with known board configurations
// in which a player has won.
//
// The return value is "X" if "X has won.
// The return value is "O" if "O has won.
// Otherwise, the return value is "".
func calculateWinner(squares [9]string) string {
	lines := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for i := 0; i < len(lines); i++ {
		a := lines[i][0]
		b := lines[i][1]
		c := lines[i][2]
		if squares[a] != "" && squares[a] == squares[b] && squares[a] == squares[c] {
			return squares[a]
		}
	}
	return ""
}

// Creates a [col, row] array based on the board square number.
func rowColumnPosition(i int) [numIndicies]int {
	var row = i / 3
	var col = i % 3
	return [2]int{col, row}
}

// Calculates the board square number from a [col, row] array.
func cellNumber(v [numIndicies]int) int {
	if (v[0] == -1) && v[1] == -1 {
		return -1
	} else {
		return v[0] + 3*v[1]
	}
}

// Generates the initial GameState object for the Game component
//
// The intial state contains a History slice with a single History element having a blank board,
// and a move that is illegal.
func (g GameDef) GetInitialState() GameState {
	history := make([]History, 1)
	history[0] = History{[9]string{"", "", "", "", "", "", "", "", ""}, [numIndicies]int{-1, -1}}
	return GameState{
		history:    history,
		stepNumber: 0,
		xIsNext:    true,
	}
}

// Selects a new position in the Game components History to display
func (g GameDef) jumpTo(index int) {
	newState := GameState{
		history:    g.State().history,
		stepNumber: index,
		xIsNext:    (index % 2) == 0,
	}
	g.SetState(newState)

}

func (g GameDef) HandleClick(index int) {
	state := g.State()
	history := make([]History, state.stepNumber+1, historyCapacity)
	copy(history, state.history[0:state.stepNumber+1])
	current := history[state.stepNumber]
	var squares [9]string
	squares = current.squares

	if calculateWinner(squares) != "" || squares[index] != "" {
		return
	}

	if state.xIsNext {
		squares[index] = "X"
	} else {
		squares[index] = "O"
	}

	// g.SetState(
	newState := GameState{
		history: append(history,
			History{squares: squares,
				move: rowColumnPosition(index)}),
		stepNumber: len(history),
		xIsNext:    !state.xIsNext,
	}
	// )
	// newState := GameState{history, len(history), !state.xIsNext}

	g.SetState(newState)
}

func (g GameState) Equals(v GameState) bool {
	if g.xIsNext != v.xIsNext {
		return false
	}
	if g.stepNumber != v.stepNumber {
		return false
	}
	if len(g.history) != len(v.history) {
		return false
	}

	for i := 0; i < len(g.history); i++ {
		for j := 0; j < numIndicies; j++ {
			if g.history[i].move[j] != v.history[i].move[j] {
				return false
			}
		}
		for j := 0; j < numSquares; j++ {
			if g.history[i].squares[j] != v.history[i].squares[j] {
				return false
			}
		}
	}
	return true
}

func (g GameDef) renderHistoryButtons(history []History) []react.RendersLi {
	var elements []react.RendersLi
	elements = make([]react.RendersLi, len(history))
	for i := 0; i < len(history); i++ {
		var tmpString string
		if i > 0 {
			tmpString = "Go to move #" + strconv.Itoa(i) + " (" + strconv.Itoa(history[i].move[0]) +
				"," + strconv.Itoa(history[i].move[1]) + ")"
		} else {
			tmpString = "Go to game start"
		}
		elements[i] = react.Li(&react.LiProps{Key: strconv.Itoa(i)},
			react.Button(&react.ButtonProps{OnClick: JumpClick{game: g, move: i}},
				react.S(tmpString),
			),
		)
	}
	return elements
}

func (g GameDef) Render() react.Element {
	state := g.State()
	history := state.history
	current := history[state.stepNumber]
	winner := calculateWinner(current.squares)

	var status string
	if winner != "" {
		status = "Winner: " + winner
	} else if state.stepNumber == 9 {
		status = "Draw"
	} else {
		status = "Next player: "
		if state.xIsNext {
			status += "X"
		} else {
			status += "O"
		}
	}
	historyButtons := g.renderHistoryButtons(history)
	return react.Div(&react.DivProps{ClassName: "game"},
		react.Div(&react.DivProps{ClassName: "game-board"},
			Board(BoardProps{
				squares:  current.squares,
				game:     g,
				selected: cellNumber(current.move),
			}),
		),
		react.Div(&react.DivProps{ClassName: "game-info"},
			react.Div(nil, react.S(status)),
			react.Ul(nil, historyButtons...),
		),
	)
}

type JumpClick struct {
	game GameDef
	move int
}

func (jc JumpClick) OnClick(e *react.SyntheticMouseEvent) {
	jc.game.jumpTo(jc.move)
	e.PreventDefault()
}

type SquareClick struct {
	game     GameDef
	position int
}

func (sc SquareClick) OnClick(e *react.SyntheticMouseEvent) {
	sc.game.HandleClick(sc.position)
	e.PreventDefault()
}
