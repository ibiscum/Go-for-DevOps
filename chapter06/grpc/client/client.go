package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ibiscum/Go-for-DevOps/chapter06/grpc/proto/qotd"
)

// Client is a client to the Quote of the day server.
type Client struct {
	client pb.QOTDClient
	conn   *grpc.ClientConn
}

// New is the constructor for Client. addr is the server's [host]:[port].
func New(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		client: pb.NewQOTDClient(conn),
		conn:   conn,
	}, nil
}

// QOTD retrieves a quote of the day. If wantAuthor is not set, will randomly choose the author
// of a quote.
func (c *Client) QOTD(ctx context.Context, wantAuthor string) (author, quote string, err error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
	}

	resp, err := c.client.GetQOTD(ctx, &pb.GetReq{Author: wantAuthor})
	if err != nil {
		return "", "", err
	}
	return resp.Author, resp.Quote, nil
}
