package main

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := Queue{}


	for i:= 0; i < 6; i++ {
		q.Enqueue(i)
		if val, ok := q.Peek(); ok != nil {
			if val != i {
				t.Error(ok)
			}
		}
	}
}

func TestDequeue(t *testing.T){
	q := Queue{}


	for i:= 0; i < 6; i++ {
		q.Enqueue(i)
		if val, ok := q.Peek(); ok != nil {
			if val != i {
				t.Error(ok)
			}
		}
	}
	//Dequeuing
	for i:= 6; i > 0; i-- {
		if val, ok := q.Dequeue(); ok != nil {
			if val != i {
				t.Error(ok)
			}
		}
	}
	
	if !q.IsEmpty(){
		t.Error("An empty queue should return true; is empty")
	}
	// should return nil for empty list
	if _, e := q.Dequeue(); e == nil {
		t.Error("Should have been nil when dequeuing an empty list")
	}
}

