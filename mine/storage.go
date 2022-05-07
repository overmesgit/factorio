package mine

import (
	"errors"
	"github.com/overmesgit/factorio/grpc"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/nodemap"
)

var MyStorage = NewStorage()

type Storage struct {
	itemByType   map[nodemap.ItemType]chan *grpc.Item
	totalStorage int
}

func NewStorage() *Storage {
	return &Storage{itemByType: make(map[nodemap.ItemType]chan *grpc.Item), totalStorage: 100}
}

var storageFull = errors.New("storage full")

func (s *Storage) Add(item nodemap.ItemType) error {
	store := s.itemByType[item]
	if store == nil {
		store = make(chan *grpc.Item, s.totalStorage)
		s.itemByType[item] = store
	}

	select {
	case store <- &grpc.Item{Type: string(item)}:
	default:
		return storageFull
	}
	return nil
}

func (s *Storage) GetCount(itemType nodemap.ItemType) int {
	val, ok := s.itemByType[itemType]
	if !ok {
		return 0
	}
	return len(val)
}

func (s *Storage) GetItemCount() []*pb.ItemCounter {
	res := make([]*pb.ItemCounter, 0)

	for itemType, ch := range s.itemByType {
		res = append(res, &pb.ItemCounter{
			Type:  string(itemType),
			Count: int64(len(ch)),
		})
	}
	return res
}

func (s *Storage) isFull(itemType nodemap.ItemType) bool {
	ch, ok := s.itemByType[itemType]
	if !ok {
		return false
	}
	return len(ch) >= s.totalStorage
}

func (s *Storage) Get(itemType nodemap.ItemType) *pb.Item {
	ch, ok := s.itemByType[itemType]
	if !ok {
		return nil
	}

	select {
	case forSend := <-ch:
		return forSend
	default:
	}
	return nil
}

func (s *Storage) GetItemForSend() *pb.Item {
	for _, ch := range s.itemByType {
		select {
		case forSend := <-ch:
			return forSend
		default:
		}
	}
	return nil
}
