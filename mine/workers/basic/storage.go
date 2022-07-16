package basic

import (
	"errors"
	"fmt"
)

type Storage struct {
	itemByType   map[ItemType]chan ItemType
	totalStorage int
}

func NewStorage() Storage {
	return Storage{itemByType: make(map[ItemType]chan ItemType), totalStorage: 100}
}

var storageFull = errors.New("storage full")

func (s *Storage) Add(item ItemType) error {
	store := s.itemByType[item]
	if store == nil {
		store = make(chan ItemType, s.totalStorage)
		s.itemByType[item] = store
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

func (s *Storage) Get(itemType ItemType) (ItemType, error) {
	ch, ok := s.itemByType[itemType]
	if !ok {
		return NoItem, errors.New(string("storage is empty " + itemType))
	}

	select {
	case forSend := <-ch:
		return forSend, nil
	default:
	}
	return NoItem, errors.New(string("storage is empty " + itemType))
}

func (s *Storage) GetAnyItem() (ItemType, error) {
	for _, ch := range s.itemByType {
		select {
		case forSend := <-ch:
			return forSend, nil
		default:
		}
	}
	return NoItem, errors.New("storage is empty")
}

func (s Storage) String() string {
	return fmt.Sprintf("%v", s.GetItemCount())
}
