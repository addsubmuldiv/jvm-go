package base

import "ch10/rtda"

// Branch 这里就只是把当前线程的 PC 加上了一个偏移量
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
