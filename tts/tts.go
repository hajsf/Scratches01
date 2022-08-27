package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func SpeakText(text string) {
	ttsdll := syscall.NewLazyDLL("base/tts.dll")
	speak := ttsdll.NewProc("rapidSpeakText")
	ptr, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		fmt.Println("Error:", err)
	}
	speak.Call(uintptr(unsafe.Pointer(ptr)))

	newTTS, err := syscall.LoadDLL("base/System.Speech.dll")
	if err != nil {
		fmt.Println("Error:", err)
	}

	ss, err := newTTS.FindProc("System.Speech.Synthesis.SpeechSynthesizer")
	if err != nil {
		fmt.Println("Error:", err)
	}
	ss.Call(uintptr(unsafe.Pointer(ptr)))
}

func main() {
	SpeakText("hi")
}
