package basic

import (
	"errors"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/mine/workers"
)

type Storage struct {
	itemByType   map[workers.ItemType]chan workers.ItemType
	totalStorage int
}

func NewStorage() Storage {
	return Storage{itemByType: make(map[workers.ItemType]chan workers.ItemType), totalStorage: 100}
}

var storageFull = errors.New("storage full")

func (s *Storage) Add(item workers.ItemType) error {
	store := s.itemByType[item]
	if store == nil {
		store = make(chan workers.ItemType, s.totalStorage)
		s.itemByType[item] = store
	}

	select {
	case store <- item:
	default:
		return storageFull
	}
	return nil
}

func (s *Storage) GetCount(itemType workers.ItemType) int {
	val, ok := s.itemByType[itemType]
	if !ok {
		return 0
	}
	return len(val)
}

func (s *Storage) GetItemCount() []*pb.ItemCounter {
	res := make([]*pb.ItemCounter, 0)

	for itemType, ch := range s.itemByType {
		res = append(
			res, &pb.ItemCounter{
				Type:  string(itemType),
				Count: int64(len(ch)),
			},
		)
	}
	return res
}

func (s *Storage) IsFull(itemType workers.ItemType) bool {
	ch, ok := s.itemByType[itemType]
	if !ok {
		return false
	}
	return len(ch) >= s.totalStorage
}

func (s *Storage) Get(itemType workers.ItemType) (workers.ItemType, error) {
	ch, ok := s.itemByType[itemType]
	if !ok {
		return workers.NoItem, errors.New(string("storage is empty " + itemType))
	}

	select {
	case forSend := <-ch:
		return forSend, nil
	default:
	}
	return workers.NoItem, errors.New(string("storage is empty " + itemType))
}

func (s *Storage) GetAnyItem() (workers.ItemType, error) {
	for _, ch := range s.itemByType {
		select {
		case forSend := <-ch:
			return forSend, nil
		default:
		}
	}
	return workers.NoItem, errors.New("storage is empty")
}

func (s Storage) String() string {
	return fmt.Sprintf("%v", s.GetItemCount())
}
