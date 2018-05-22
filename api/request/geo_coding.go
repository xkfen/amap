package request

/**
	该api用于地理编码
 */

 // 地理编码响应参数
type GenCodingResp struct {
	// 返回结果状态值:0 表示请求失败；1 表示请求成功。
	Status int `json:"status"`
	// 返回结果数目:结果个数
	Count int `json:"count"`
	// 返回状态说明当 status 为 0 时，info 会返回具体错误原因，否则返回“OK”
	Info string `json:"info"`
	// 地理编码信息列表
	GeoCodes []GeoCodes `json:"geocodes"`
}

type GeoCodes struct {
	// 结构化地址信息:省份＋城市＋区县＋城镇＋乡村＋街道＋门牌号码
	FormattedAddress string `json:"formatted_address"`
	// 地址所在的省份
	Province string `json:"province"`
	// 地址所在的城市名
	City string `json:"city"`
	// 城市编码
	CityCode string `json:"city_code"`
	// 地址所在的区
	District string `json:"district"`
	// 地址所在的乡镇
	Township string `json:"township"`
	// 街道
	Street string `json:"street"`
	// 门牌
	Number string `json:"number"`
	// 区域编码
	Adcode string `json:"adcode"`
	// 坐标点
	Location string `json:"location"`
	// 匹配级别
	Level string `json:"level"`
}

func GenCoding(key , address string)(resp GenCodingResp, err error){


}