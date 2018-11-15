package service

import (
	"time"

	"github.com/hal-ms/driver/rpio"
)

var Boss = newBossService()

const (
	freq       = 64000
	deg0usec   = 750
	deg180usec = 2250
)

type bossService struct {
	pin rpio.Pin
}

func newBossService() bossService {
	err := rpio.Open()

	if err != nil {
		defer rpio.Close()
		panic(err)
	}

	pin := rpio.Pin(19)
	pin.Pwm()
	pin.Freq(freq)
	pin.DutyCycle(0, 1280)

	return bossService{pin: pin}
}

func (s *bossService) Pow(deg int) {
	duty := deg0usec + deg*(deg180usec-deg0usec)/180
	dutyLen := uint32(duty * 1280 / 20000)
	pin.DutyCycle(dutyLen, 1280)
	time.Sleep(time.Second / 1280)

	time.Sleep(time.Second)

	duty := deg0usec + 0*(deg180usec-deg0usec)/180
	dutyLen := uint32(duty * 1280 / 20000)
	pin.DutyCycle(dutyLen, 1280)
	time.Sleep(time.Second / 1280)

}
