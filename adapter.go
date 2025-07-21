package main

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Client struct {}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer")
	com.InsertIntoLightningPort()
}

type Mac struct {}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac computer")
}

type Windows struct {}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine")
}

type WindowsAdapter struct {
	machine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning into USB")
	w.machine.InsertIntoUSBPort()
}
