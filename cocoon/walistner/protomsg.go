package main

import d "github.com/jhump/protoreflect/dynamic"

type ProtoMessage interface {
	ProtoReflect() d.Message
}

func Conv() {
	x := new(ProtoMessage)
	y := *x.New()
}
