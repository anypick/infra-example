package dao

import (
	"context"
	"database/sql"
	"example/src/dto"
	"github.com/anypick/infra-mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func GetBySiteCode(siteCode string, ctx context.Context) (dto.SiteDto, error) {
	query := `select id, site_id, site_code, site_name, create_time, update_time from site where site_code = ?`
	site := dto.SiteDto{}
	err := basesql.ExecuteContext(nil, func(runner *basesql.Runner) error {
		var row *sql.Row
		if runner.Tx == nil {
			row = runner.Db.QueryRow(query, siteCode)
		} else {
			row = runner.Tx.QueryRow(query, siteCode)
		}
		err := row.Scan(&site.Id, &site.SiteId, &site.SiteCode, &site.SiteName, &site.CreateTime, &site.UpdateTime)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return site, err
	}
	return site, nil
}

func BatchInsert(siteDtos []dto.SiteDto, ctx context.Context) error {
	insertSql := `insert into site (site_id, site_code, site_name) values (?, ?, ?)`
	err := basesql.ExecuteContext(ctx, func(runner *basesql.Runner) error {
		for i, site := range siteDtos {
			result, er := runner.Tx.Exec(insertSql, site.SiteId, site.SiteCode, site.SiteName)
			// 故意报错，测试事务
			if i == 3 {
				return errors.New("出现错误")
			}
			if er != nil {
				return er
			}
			logrus.Info(result.LastInsertId())
		}
		return nil
	})
	return err
}
