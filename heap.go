package pairingheap

type Heap struct {
  Min Element
  subheaps []Heap
}

func (h Heap) Merge(h2 Heap) Heap {
  if h.Min == nil {
    return h2
  }
  if h2.Min == nil {
    return h
  }
  if h.Min.Priority() < h2.Min.Priority() {
    subheaps := make([]Heap, len(h.subheaps) + 1)
    subheaps[0] = h2
    copy(subheaps[1:], h.subheaps)
    return Heap{h.Min, subheaps}
  }
  subheaps := make([]Heap, len(h2.subheaps) + 1)
  subheaps[0] = h
  copy(subheaps[1:], h2.subheaps)
  return Heap{h2.Min, subheaps}
}

func (h Heap) Insert(elem Element) Heap {
  return h.Merge(Heap{elem, nil})
}

func mergePairs(hs []Heap) Heap {
  if len(hs) == 0 {
    return Heap{}
  } else if len(hs) == 1 {
    return hs[0]
  } else {
    return hs[0].Merge(hs[1]).Merge(mergePairs(hs[2:]))
  }
}

func (h Heap) DeleteMin() Heap {
  return mergePairs(h.subheaps)
}
