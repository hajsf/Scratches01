package main

import (
	"context"
	"fmt"
)

func Connect() {
	fmt.Println("Connected")
	if client.IsConnected() {
		passer.logs <- "Reconnecting client"
		client.Disconnect()
	}

	if client.Store.ID == nil {
		// No ID stored, new login
	GetQR:
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			//	panic(err)
			passer.logs <- "Can not connect with WhatApp server, try again later"
			fmt.Println("Sorry", err)

		}

		for evt := range qrChan {
			switch evt.Event {
			case "success":
				{
					passer.logs <- "success"
					fmt.Println("Login event: success")
				}
			case "timeout":
				{
					passer.logs <- "timeout/Refreshing"
					fmt.Println("Login event: timeout")
					goto GetQR
				}
			case "code":
				{
					fmt.Println("new code recieved")
					fmt.Println(evt.Code)
					passer.logs <- evt.Code
				}
			}
		}
	} else {
		// Already logged in, just connect
		passer.logs <- "Already logged"
		fmt.Println("Already logged")
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}
}
