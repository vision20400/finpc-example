package grpc

import (
	"context"
	"database/sql"

	// external packages
	log "github.com/sirupsen/logrus"
	"github.com/speps/go-hashids"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Trading struct {
	TradingServer
}

func (s *Trading) GetStockList(ctx context.Context, empty *emptypb.Empty) (*StockListResp, error) {

	db := ctx.Value(DBSession).(*sql.DB)
	hash := ctx.Value(HashID).(*hashids.HashID)

	rows, err := db.Query("SELECT id, code, name, total_stock_count FROM stocks ORDER BY code;")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	list := make([]*Stock, 0, 10)

	for rows.Next() {
		var id int64
		var code, name string
		var totalStockCount uint32
		if err := rows.Scan(&id, &code, &name, &totalStockCount); err != nil {
			log.Fatal(err)
		}

		enc, _ := hash.EncodeInt64([]int64{id})

		list = append(list, &Stock{
			Id:              enc,
			Code:            code,
			Name:            name,
			TotalStockCount: totalStockCount,
		})
	}

	return &StockListResp{
		StockList: list,
	}, nil
}
