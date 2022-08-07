// Code generated by goctl. DO NOT EDIT.
package types

type GoodsSearchPcReq struct {
	Category    string `json:"category,optional"`
	GoodsName   string `json:"goodsName,optional"`
	LastGoodsId int64  `json:"lastGoodsId,default=0"`
	Size        int    `json:"size"`
}

type GoodsSearchPcResp struct {
	IsNext    int         `json:"isNext"`
	GoodsList []GoodsPcVO `json:"goodsList"`
}

type GoodsPcVO struct {
	GoodsId   int64  `json:"goodsId"`
	GoodsName string `json:"goodsName"`
	Title     string `json:"title"`
	Company   string `json:"company"`
	Category  string `json:"category"`
	Logo      string `json:"logo"`
	Content   string `json:"content"`
	SalePrice int    `json:"salePrice"`
	MinPrice  int    `json:"minPrice"`
	MaxPrice  int    `json:"maxPrice"`
	Stock     int    `json:"stock"`
	Sale      int    `json:"sale"`
	IsAttr    int    `json:"isAttr"`
	Status    int    `json:"status"`
	Deleted   int    `json:"deleted"`
	CreateAt  int64  `json:"createAt"`
	UpdateAt  int64  `json:"updateAt"`
}

type GoodsSearchMtReq struct {
	Category    string `json:"category,optional"`
	GoodsName   string `json:"goodsName,optional"`
	LastGoodsId int64  `json:"lastGoodsId"`
	Size        int    `json:"size"`
}

type GoodsSearchMtResp struct {
	IsNext    int         `json:"isNext"`
	GoodsList []GoodsMtVO `json:"goodsList"`
}

type GoodsMtVO struct {
	GoodsId   int64  `json:"goodsId"`
	GoodsName string `json:"goodsName"`
	Title     string `json:"title"`
	Company   string `json:"company"`
	Category  string `json:"category"`
	Logo      string `json:"logo"`
	Content   string `json:"content"`
	SalePrice int    `json:"salePrice"`
	MinPrice  int    `json:"minPrice"`
	MaxPrice  int    `json:"maxPrice"`
	Stock     int    `json:"stock"`
	Sale      int    `json:"sale"`
	IsAttr    int    `json:"isAttr"`
	Status    int    `json:"status"`
	Deleted   int    `json:"deleted"`
}
