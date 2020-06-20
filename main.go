 package main

import "fmt"

type Item struct {
  next *Item
  Value interface{}
}


func (i *Item) Next() *Item {
  if i.next == nil {
    i.next = &Item{}
  }
  return i.next
}

type Ring struct{
  list *Item
  head, tail *Item
}

func New(size int) *Ring {
  r := &Ring{
    list: &Item{},
  }

  r.head = r.list
  r.tail = r.list

  for i:=0; i<size; i++ {
    r.head = r.head.Next()
  }
  
  r.head.next = r.list
  r.head = r.head.Next()
  return r
}

func (r *Ring) Write(i interface{}) {
  r.head.Value = i
  if r.head.Next() == r.tail{
    return
  }

  r.head = r.head.Next()
}

func (r *Ring) Read() interface{} {
  if r.tail == r.head {
    return nil
  }

  r.tail = r.tail.Next()
  return r.tail.Value
}

func main() {

  r := New(5)

  for i:=0; i<10; i++ {
    r.Write(i)
  }
  
  for i:=r.Read(); i != nil;{
    fmt.Println(i)
    i = r.Read()
  }
}