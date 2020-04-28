package services

import (
	"context"
	"fmt"
)

type ProdService struct{}

func (p *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (res *ProdResponse, err error) {
	return &ProdResponse{ProdStock: request.ProdId}, nil
}

func (p *ProdService) GetProdStocks(ctx context.Context, request *QuerySize) (res *ProdListResponse, err error) {
	prodInfo := make([]*ProdResponse, 0, int(request.Size))

	for i := 0; i < int(request.Size); i++ {
		prodInfo = append(prodInfo, &ProdResponse{ProdStock: int32(i)})
	}
	fmt.Println(prodInfo)
	res = new(ProdListResponse)
	res.ProdInfo = prodInfo
	return res, err
}
