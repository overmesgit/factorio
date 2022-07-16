package grpcsender

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers/basic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Sender struct {
}

var _ basic.Sender = Sender{}

func NewSender() Sender {
	return Sender{}
}

func (s Sender) SendItem(adjNode basic.Node, forSend basic.ItemType) error {
	sugar.Sugar.Infof("Send items. Current store. %v forSend %v", forSend)

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

	r, err := c.ReceiveResource(ctx, &pb.Item{Type: string(forSend)})
	if err != nil {
		return err
	}

	sugar.Sugar.Infof("Sent item %v. Resp %v.", forSend, r)
	return nil
}

func (s Sender) AskForItem(
	prevNode basic.Node, itemType basic.ItemType,
) (basic.ItemType, error) {
	sugar.Sugar.Infof("Ask for item %v %v", prevNode, itemType)

	conn, err := grpc.Dial(
		fmt.Sprintf("r%vc%v:8080", prevNode.Row, prevNode.Col),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return basic.NoItem, errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetResource(ctx, &pb.Item{Type: string(itemType)})
	if err != nil {
		return basic.NoItem, err
	}

	return basic.ItemType(r.Type), nil
}

func (s Sender) AskForNeedItem(nextNode basic.Node) (basic.ItemType, error) {
	sugar.Sugar.Infof("Ask for needed item %v", nextNode)

	conn, err := grpc.Dial(
		fmt.Sprintf("r%vc%v:8080", nextNode.Row, nextNode.Col),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return basic.NoItem, errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	item, err := c.NeededResource(ctx, &pb.Empty{})
	if err != nil {
		return basic.NoItem, err
	}
	return basic.ItemType(item.Type), err
}
