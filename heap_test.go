package pairingheap

import "testing"

func TestInsertIntoEmpty(t *testing.T) {
  e := IntElement(1)
  h := Heap{}
  if h.Insert(e).Min != e {
    t.Errorf("Inserted element is not the minimum.")
  }
}

func testInsertIntoMax(t *testing.T) {
  e1 := IntElement(1)
  e2 := IntElement(2)
  h := Heap{}.Insert(e2)
  if h.Insert(e1).Min != e1 {
    t.Errorf("Inserted element is not the minimum.")
  }
}

func testInsertIntoMin(t *testing.T) {
  e1 := IntElement(1)
  e2 := IntElement(2)
  h := Heap{}.Insert(e1)
  if h.Insert(e2).Min != e1 {
    t.Errorf("Inserted element is the minimum.")
  }
}

func testInsertIntoMaxAndDelete(t *testing.T) {
  e1 := IntElement(1)
  e2 := IntElement(2)
  h := Heap{}.Insert(e2)
  if h.Insert(e1).Min != e1 {
    t.Errorf("Inserted element is not the minimum.")
  }
  if h.DeleteMin().Min != e2 {
    t.Errorf("Inserted element was not deleted.")
  }
}

func testInsertIntoMinAndDelete(t *testing.T) {
  e1 := IntElement(1)
  e2 := IntElement(2)
  h := Heap{}.Insert(e1)
  if h.Insert(e2).Min != e1 {
    t.Errorf("Inserted element is the minimum.")
  }
  if h.DeleteMin().Min != e1 {
    t.Errorf("Inserted element was deleted.")
  }
}
