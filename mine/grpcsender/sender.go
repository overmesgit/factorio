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

func GrpcItemToItem(item *pb.Item) basic.Item {
	var ingredientList []*basic.Item
	for _, ingredient := range item.Ingredients {
		grpcItem := GrpcItemToItem(ingredient)
		ingredientList = append(ingredientList, &grpcItem)
	}

	return basic.Item{
		ItemType:    basic.ItemType(item.Type),
		Id:          item.Id,
		Parents:     item.Parents,
		Ingredients: ingredientList,
	}
}

func ItemToGrpcItem(item basic.Item) pb.Item {
	var ingredientList []*pb.Item
	for _, ingredient := range item.Ingredients {
		grpcItem := ItemToGrpcItem(*ingredient)
		ingredientList = append(ingredientList, &grpcItem)
	}
	return pb.Item{
		Type:        string(item.ItemType),
		Id:          item.Id,
		Parents:     item.Parents,
		Ingredients: ingredientList,
	}
}

func (s Sender) SendItem(adjNode basic.Node, forSend basic.Item) error {
	sugar.Sugar.Infof("Send items. ForSend %v", forSend)

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

	item := ItemToGrpcItem(forSend)
	r, err := c.ReceiveResource(ctx, &item)
	if err != nil {
		return err
	}

	sugar.Sugar.Infof("Sent item %v. Resp %v.", forSend, r)
	return nil
}

func (s Sender) AskForItem(
	prevNode basic.Node, itemType basic.ItemType,
) (basic.Item, error) {
	sugar.Sugar.Infof("Ask for item %v %v", prevNode, itemType)

	conn, err := grpc.Dial(
		fmt.Sprintf("r%vc%v:8080", prevNode.Row, prevNode.Col),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return basic.Item{ItemType: basic.NoItem}, errors.New(
			fmt.Sprintf(
				"did not connect: %v", err,
			),
		)
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetResource(ctx, &pb.Item{Type: string(itemType)})
	if err != nil {
		return basic.Item{ItemType: basic.NoItem}, err
	}

	return GrpcItemToItem(r), nil
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
