package basic

import (
	"errors"
	"fmt"
)

type Storage struct {
	itemByType   map[ItemType]chan Item
	totalStorage int
}

func NewStorage() Storage {
	return Storage{itemByType: make(map[ItemType]chan Item), totalStorage: 100}
}

var storageFull = errors.New("storage full")

func (s *Storage) Add(item Item) error {
	store := s.itemByType[item.ItemType]
	if store == nil {
		store = make(chan Item, s.totalStorage)
		s.itemByType[item.ItemType] = store
	}

	select {
	case store <- item:
	default:
		return storageFull
	}
	return nil
}

func (s *Storage) GetCount(itemType ItemType) int {
	val, ok := s.itemByType[itemType]
	if !ok {
		return 0
	}
	return len(val)
}

func (s *Storage) GetItemCount() []ItemCounter {
	res := make([]ItemCounter, 0)

	for itemType, ch := range s.itemByType {
		res = append(
			res, ItemCounter{
				Type:  string(itemType),
				Count: int64(len(ch)),
			},
		)
	}
	return res
}

func (s *Storage) IsFull(itemType ItemType) bool {
	ch, ok := s.itemByType[itemType]
	if !ok {
		return false
	}
	return len(ch) >= s.totalStorage
}

func (s *Storage) Get(itemType ItemType) (Item, error) {
	ch, ok := s.itemByType[itemType]
	if !ok {
		return Item{ItemType: NoItem}, errors.New(string("storage is empty " + itemType))
	}

	select {
	case forSend := <-ch:
		return forSend, nil
	default:
	}
	return Item{ItemType: NoItem}, errors.New(string("storage is empty " + itemType))
}

func (s *Storage) GetAnyItem() (Item, error) {
	for _, ch := range s.itemByType {
		select {
		case forSend := <-ch:
			return forSend, nil
		default:
		}
	}
	return Item{ItemType: NoItem}, errors.New("storage is empty")
}

func (s Storage) String() string {
	return fmt.Sprintf("%v", s.GetItemCount())
}
