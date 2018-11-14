package service

import (
	"github.com/tarm/serial"
)

var Building = newBuildingService()

type buildingService struct {
	conn *serial.Port
}

func newBuildingService() buildingService {
	c := &serial.Config{Name: "COM9", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		//log.SendSlack(err.Error())
		panic(err)
	}
	conn := s
	return buildingService{conn: conn}
}

func (s *buildingService) Bar() error {
	s.write([]byte{0x00})
	return nil
}

// 信号送信
func (s *buildingService) write(b []byte) error {
	_, err := s.conn.Write(b)
	if err != nil {
		//log.SendSlack(err.Error())
		panic(err)
	}
	return nil
}
