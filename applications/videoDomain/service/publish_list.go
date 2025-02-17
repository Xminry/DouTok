package service

import (
	"context"
	"fmt"
	"strconv"

	tools "github.com/TremblingV5/DouTok/applications/videoDomain/misc"
	"github.com/TremblingV5/DouTok/applications/videoDomain/typedef"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/misc"
)

type QueryPublishListService struct {
	ctx context.Context
}

func NewQueryPublishListService(ctx context.Context) *QueryPublishListService {
	return &QueryPublishListService{ctx: ctx}
}

func (s *QueryPublishListService) QueryPublishListInHBase(user_id int64) ([]*typedef.VideoInHB, error) {
	user_id_string := strconv.FormatInt(user_id, 10)
	user_id_string = tools.FillUserId(fmt.Sprint(user_id))

	filters := hbaseHandle.GetFilterByRowKeyPrefix(user_id_string)

	video_list, err := HBClient.Scan(
		"publish", filters...,
	)

	list := []*typedef.VideoInHB{}
	if err != nil {
		return list, err
	}

	for _, v := range video_list {
		temp := typedef.VideoInHB{}
		err := misc.Map2Struct4HB(v, &temp)
		if err != nil {
			continue
		}
		list = append(list, &temp)
	}

	return list, nil
}
