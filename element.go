package pairingheap

type Element interface {
  Priority() int64
}

type IntElement int

func (i IntElement) Priority() int64 {
  return int64(i)
}

type Int64Element int64

func (i Int64Element) Priority() int64 {
  return int64(i)
}
