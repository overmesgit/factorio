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
	totalStorage int32
}

func NewStorage() *Storage {
	return &Storage{itemByType: make(map[nodemap.ItemType]chan *grpc.Item)}
}

var storageFull = errors.New("storage full")

func (s *Storage) Add(item nodemap.ItemType) error {
	store := s.itemByType[item]
	if store == nil {
		store = make(chan *grpc.Item, 100)
		s.itemByType[item] = store
	}

	select {
	case store <- &grpc.Item{Type: string(item)}:
	default:
		return storageFull
	}
	return nil
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
