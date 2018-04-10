package display

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

// ComponentOptionFactory is essentially any function that returns a
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
	duration time.Duration,
	easing EasingFunc) string {
	// TODO(lbayes): Make a hash from these values instead.
	return fmt.Sprintf("%v:%v:%v:%v:%v", option, start, finish, duration, easing)
}

// Transition is a helper that allows us to define and name Transitions in order
// to later apply them as Traits.
func Transition(option ComponentOptionAssigner,
	start interface{},
	finish interface{},
	durationMs time.Duration,
	easingFunc EasingFunc) ComponentOption {

	key := transitionToKey(option, start, finish, durationMs, easingFunc)
	log.Println("KEY:", key)

	log.Println("Transition created!")

	optionValue := reflect.ValueOf(option)

	var startTime time.Time
	var totalDistance float64

	var update = func(d Displayable) {
		elapsedTimeSeconds := time.Since(startTime).Seconds()
		percentComplete := float64(elapsedTimeSeconds) / float64(durationMs)

		// log.Println("ELAPSED", elapsedTimeSeconds)
		// log.Println("PERCENT", percentComplete)
		// log.Println("totalDistance", totalDistance)

		newValue := start.(float64) + (totalDistance * percentComplete)
		dValue := reflect.ValueOf(d)

		// TODO(lbayes): Clean up this mess!
		applicators := optionValue.Call([]reflect.Value{reflect.ValueOf(newValue)})
		applicators[0].Call([]reflect.Value{dValue})
		// d.Invalidate()
		// log.Println("ERROR:", errValue[0])
	}

	var listen = func(d Displayable) {
		startTime = time.Now()
		totalDistance = (finish.(float64) - start.(float64))
		d.OnEnterFrame(update)
		update(d)
	}

	return func(d Displayable) error {
		listen(d)
		return nil
	}
}
