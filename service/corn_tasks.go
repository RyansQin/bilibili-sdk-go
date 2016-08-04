package service

import (
	"strconv"
)

var (
	_INDEX_SORTS = []int{
		1, 13, 3, 129, 5, 4, 119, 36,
	}
)

type SortRankInfo struct {
	SortName string `json:"sort_name"`
	Videos   []interface{} `json:"videos"`
}

type IndexInfoTask struct {
	CornTask
	app *BiliBiliApplication
}

func (i *IndexInfoTask) Run() error {
	retInfo := make([]SortRankInfo, 0, len(_INDEX_SORTS))
	for _, sortId := range _INDEX_SORTS {
		back, err := i.app.Client.Rank.SortRank(sortId, 1, 8, "hot")
		if err != nil {
			return err
		}
		videos := make([]interface{}, 0, len(back.List))
		for i := 0; i < len(back.List); i++ {
			videos = append(videos, back.List[strconv.Itoa(i)])
		}

		sortRank := SortRankInfo{SortName:back.Name, Videos:videos}

		retInfo = append(retInfo, sortRank)

		sortCacheName := SORT_TOP_CACHE + strconv.Itoa(sortId)

		i.app.Cache.SetCache(sortCacheName, sortRank)

	}
	i.app.Cache.SetCache(INDEX_CACHE, retInfo)
	return nil
}

type BangumiInfoTask struct {
	CornTask
	app *BiliBiliApplication
}

func (i *BangumiInfoTask) Run() error {
	ret, err := i.app.Client.Bangumi.GetIndex()
	if err != nil {
		return err
	}
	i.app.Cache.SetCache(BANGUMI_CACHE, ret)
	return nil
}

type BangumiListTask struct {
	CornTask
	app *BiliBiliApplication
}

func (i *BangumiListTask) Run() error {
	ret, err := i.app.Client.Bangumi.GetWeekList("2")
	if err != nil {
		return err
	}
	i.app.Cache.SetCache(BANGUMI_LIST_CACHE, ret)
	return nil
}

type TopRankTask struct {
	CornTask
	app *BiliBiliApplication
}

func (i *TopRankTask) Run() error {
	ret, err := i.app.Client.Rank.SortRank(0, 1, 10, "hot")
	if err != nil {
		return err
	}
	i.app.Cache.SetCache(ALL_RANK_CACHE, ret)
	return nil
}

