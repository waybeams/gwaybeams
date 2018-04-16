package display

import (
	"fmt"
	"reflect"
	"time"
)

// ComponentOptionAssigner is essentially any function that returns a
// ComponentOption, but we can't make Go's type system play nice with
// the fact that these outer functions may have any interface at all.
// The main point is, that any function you can use to apply a
// Displayable property in a Component declaration and be dropped in
// here by reference.
type ComponentOptionAssigner interface{}

type EasingFunc func(float64) float64

func transitionToKey(
	option ComponentOptionAssigner,
	start interface{},
	finish interface{},
	duration int,
	easing EasingFunc) string {
	// TODO(lbayes): Make a hash from these values instead.
	return fmt.Sprintf("%v:%v:%v:%v:%v", option, start, finish, duration, easing)
}

// Transition is a helper that allows us to define and name Transitions in order
// to later apply them as Traits.
func Transition(b Builder, option ComponentOptionAssigner,
	start interface{},
	finish interface{},
	durationMs int,
	easingFunc EasingFunc) ComponentOption {

	clock := b.Clock()

	// TODO(lbayes): Store transitions and try to match existing instances after recompose.
	// key := transitionToKey(option, start, finish, durationMs, easingFunc)
	optionValue := reflect.ValueOf(option)

	var startTime time.Time
	var totalDistance float64
	var unsub Unsubscriber

	var update = func(d Displayable) {
		elapsedTimeMs := clock.Since(startTime).Nanoseconds() / int64(time.Millisecond)

		var percentComplete float32

		if elapsedTimeMs == 0 {
			percentComplete = 0.0
		} else {
			percentComplete = float32(elapsedTimeMs) / float32(durationMs)
		}

		if elapsedTimeMs > (int64(durationMs) * 1) {
			unsub()
			return
		}

		// TODO(lbayes): Can't assume transitioned values are float64
		newValue := start.(float64) + (totalDistance * easingFunc(float64(percentComplete)))
		dValue := reflect.ValueOf(d)

		// TODO(lbayes): Clean up this mess!
		applicators := optionValue.Call([]reflect.Value{reflect.ValueOf(newValue)})
		applicators[0].Call([]reflect.Value{dValue})
	}

	return func(d Displayable) error {
		startTime = clock.Now()
		totalDistance = (finish.(float64) - start.(float64))

		unsub = d.Builder().OnEnterFrame(func(e Event) {
			update(d)
		})

		// Trigger the handler with the component instance:
		update(d)
		return nil
	}
}
