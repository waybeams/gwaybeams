package events

// Gesture Notifications (past tense)
const Moved = "Moved"
const Pressed = "Pressed"
const Released = "Released"
const KeyPressed = "KeyPressed"
const KeyReleased = "KeyReleased"

// Component Notifications (past tense)
const Blurred = "Blurred"
const Clicked = "Clicked"
const DragEnded = "DragEnded"
const DragStarted = "DragStarted"
const Entered = "Entered"
const Exited = "Exited"
const Focused = "Focused"
const FrameEntered = "FrameEntered"
const Hovered = "Hovered"

// Navigation Requests (present tense)
const MoveBackward = "MoveBackward"
const MoveDown = "MoveDown"
const MoveForward = "MoveForward"
const MoveLeft = "MoveLeft"
const MoveNext = "MoveNext"
const MovePrevious = "MovePrevious"
const MoveRight = "MoveRight"
const MoveUp = "MoveUp"

var AllEvents = []string{
	// Gesture Notifications
	Moved,
	Pressed,
	Released,
	KeyPressed,
	KeyReleased,

	// Component Notifications
	Blurred,
	Clicked,
	DragEnded,
	DragStarted,
	Entered,
	Exited,
	Focused,
	FrameEntered,
	Hovered,

	// Navigation Requests
	MoveBackward,
	MoveDown,
	MoveForward,
	MoveLeft,
	MoveNext,
	MovePrevious,
	MoveRight,
	MoveUp,
}
