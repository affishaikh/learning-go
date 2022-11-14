package main

type Value interface {
	Value() Value
}

type StringValue struct {
	val string
}

type BooleanValue struct {
	val bool
}

func (s StringValue) Value() StringValue {
	return StringValue{s.val}
}


func (b BooleanValue) Value() BooleanValue {
	return BooleanValue{b.val}
}