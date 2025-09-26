package now_timer

import "time"

type nowTimerImpl struct {
	f func() time.Time
}

func New(f func() time.Time) *nowTimerImpl {
	return &nowTimerImpl{
		f: f,
	}
}
