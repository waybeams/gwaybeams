package events

// Gesture Notifications (past tense)
const CharEntered = "CharEntered"
const EnterKeyReleased = "EnterKeyReleased"
const KeyEntered = "KeyEntered"
const KeyPressed = "KeyPressed"
const KeyReleased = "KeyReleased"
const Moved = "Moved"
const Pressed = "Pressed"
const Released = "Released"

// Spec Notifications (past tense)
const Blurred = "Blurred"
const Clicked = "Clicked"
const DragEnded = "DragEnded"
const DragStarted = "DragStarted"
const Entered = "Entered"
const Exited = "Exited"
const Focused = "Focused"
const FrameEntered = "FrameEntered"
const Hovered = "Hovered"
const Submitted = "Submitted"
const TextChanged = "TextChanged"

// Navigation Requests (present tense)
const MoveBackward = "MoveBackward"
const MoveDown = "MoveDown"
const MoveForward = "MoveForward"
const MoveLeft = "MoveLeft"
const MoveNext = "MoveNext"
const MovePrevious = "MovePrevious"
const MoveRight = "MoveRight"
const MoveUp = "MoveUp"

// Spec Lifecycle
const Configured = "Configured"
const Created = "Created"
const DrawCompleted = "DrawCompleted"
const Invalidated = "Invalidated"
const LayoutCompleted = "LayoutCompleted"

var AllEvents = []string{
	// Gesture Notifications
	CharEntered,
	EnterKeyReleased,
	Moved,
	Pressed,
	Released,
	KeyEntered,
	KeyPressed,
	KeyReleased,

	// Spec Notifications
	Blurred,
	Clicked,
	DragEnded,
	DragStarted,
	Entered,
	Exited,
	Focused,
	FrameEntered,
	Hovered,
	Submitted,
	TextChanged,

	// Navigation Requests
	MoveBackward,
	MoveDown,
	MoveForward,
	MoveLeft,
	MoveNext,
	MovePrevious,
	MoveRight,
	MoveUp,

	// Spec Lifecycle
	Configured,
	Created,
	DrawCompleted,
	Invalidated,
	LayoutCompleted,
}
