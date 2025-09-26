package now_timer

import "time"

func (n *nowTimerImpl) Now() time.Time {
	return n.f()
}
