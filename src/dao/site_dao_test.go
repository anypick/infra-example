package dao

import (
	"context"
	"example/src/dto"
	_ "example/testx"
	"fmt"
	"github.com/anypick/infra-mysql"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetBySiteCode(t *testing.T) {
	var (
		dto dto.SiteDto
		err error
	)
	err = basesql.DbWithQuery(func(runner *basesql.Runner) error {
		dto, err = GetBySiteCode("S1001", basesql.WithValueContext(context.TODO(), runner))
		return nil
	})

	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(dto)
}

func TestBatchInsert(t *testing.T) {
	dtos := make([]dto.SiteDto, 0)
	for i := 10008; i <= 100010; i++ {
		dtos = append(dtos, dto.SiteDto{SiteId: int64(i), SiteCode: fmt.Sprintf("S%d", i), SiteName: fmt.Sprintf("name%d", i)})
	}
	err := basesql.DbTxRunner(func(runner *basesql.Runner) error {
		err := BatchInsert(dtos, basesql.WithValueContext(context.Background(), runner))
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	})

	if err != nil {
		logrus.Error(err)
	}
}
