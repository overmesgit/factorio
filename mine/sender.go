package mine

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/nodemap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) SendItem(adjNode *pb.Node, forSend *pb.Item) error {
	Sugar.Infof("Send items. Current store. %v forSend %v", MyStorage.GetItemCount(), forSend)

	conn, err := grpc.Dial(
		fmt.Sprintf("r%vc%v:8080", adjNode.Row, adjNode.Col),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SendResource(ctx, forSend)
	if err != nil {
		return err
	}

	Sugar.Infof("Sent item %v. Resp %v.", forSend, r)
	return nil
}

func (s *Sender) AskForItem(
	prevNode *pb.Node, itemType nodemap.ItemType, store bool,
) (*pb.Item, error) {
	Sugar.Infof("Ask for item %v %v", prevNode, itemType)

	conn, err := grpc.Dial(
		fmt.Sprintf("r%vc%v:8080", prevNode.Row, prevNode.Col),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GiveResource(ctx, &pb.Item{Type: string(itemType)})
	if err != nil {
		return nil, err
	}
	if store {
		err := MyStorage.Add(nodemap.ItemType(r.Type))
		if err != nil {
			Sugar.Warnf("can't store aquared item %v %v", r, err)
		}
	}

	return r, nil
}

func (s *Sender) AskForNeedItem(nextNode *pb.Node) (*pb.Item, error) {
	Sugar.Infof("Ask for needed item %v", nextNode)

	conn, err := grpc.Dial(
		fmt.Sprintf("r%vc%v:8080", nextNode.Row, nextNode.Col),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.NeededResource(ctx, &pb.Empty{})
}
