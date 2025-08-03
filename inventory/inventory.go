package inventory

type Item struct {
	Name     string
	Quantity int
	MinAlert int
}

type Store struct {
	items []Item
}

func (s *Store) Add(item Item) {
	s.items = append(s.items, item)
}

func (s *Store) AdjustQuantity(name string, delta int) {
	for i, it := range s.items {
		if it.Name == name {
			s.items[i].Quantity += delta
			return
		}
	}
}

func (s *Store) Items() []Item {
	return append([]Item(nil), s.items...)
}
