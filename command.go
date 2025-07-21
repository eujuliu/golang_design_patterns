package main

import "fmt"

type Command interface {
	execute()
}

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Device interface {
	on()
	off()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()

}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
