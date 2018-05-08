package events

// Gesture Notifications (past tense)
const CharEntered = "CharEntered"
const Moved = "Moved"
const Pressed = "Pressed"
const Released = "Released"
const KeyPressed = "KeyPressed"
const KeyReleased = "KeyReleased"

// Control Notifications (past tense)
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

// Control Lifecycle
const Configured = "Configured"
const Created = "Created"
const DrawCompleted = "DrawCompleted"
const LayoutCompleted = "LayoutCompleted"

var AllEvents = []string{
	// Gesture Notifications
	CharEntered,
	Moved,
	Pressed,
	Released,
	KeyPressed,
	KeyReleased,

	// Control Notifications
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

	// Control Lifecycle
	Configured,
	Created,
	DrawCompleted,
	LayoutCompleted,
}
