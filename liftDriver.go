package liftDriver

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
#include "elev.h"
*/
import "C"

//should we create a button_type enum in go??
//type buttonType int

func LiftDriver_Initialize() bool {
	return int(C.elev_init()) != 0
}

func liftDriver_setMotorDirection(direction int) {
	C.elev_set_motor_direction(C.elev_motor_direction_t(direction))
}
func liftDriver_SetButtonLamp(button int, floor int, onOrOff bool) {
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(onOrOff))
}
func liftDriver_SetFloorIndicator(floor int) {
	C.elev_set_floor_indicatior(C.int(floor))
}

func liftDriver_SetDoorLamp(onOrOff bool) {
	C.elev_set_door_open_lamp(C.int(onOrOff))
}

func liftDriver_GetButtonSignal(button int, floor int) bool {
	return int(C.elev_get_button_signal(C.elev_button_type(button)))
}

func liftDriver_GetFloor() int {
	return int(C.elev_get_floor_sensor_signal())
}

//func liftDriver_DetectButtonEvent()
