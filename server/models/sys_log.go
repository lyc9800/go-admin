package models

import "gorm.io/gorm"

type SysLog struct {
	gorm.Model
	Browser      string `gorm:"column:browser;type:varchar(255)" json:"browser"`
	Ip           string `gorm:"column:ip;type:varchar(255)" json:"ip"`
	ClassMethod  string `gorm:"column:class_method;type:varchar(255)" json:"class_method"`
	HttpMethod   string `gorm:"column:http_method;type:varchar(255)" json:"http_method"`
	RequestParam string `gorm:"column:request_param;type:text" json:"request_param"`
	RequestUri   string `gorm:"column:request_uri;type:varchar(255)" json:"request_uri"`
	Result       string `gorm:"column:result;type:text" json:"result"`
	StatusCode   int    `gorm:"column:status_code;type:int(11)" json:"status_code"`
	UseTime      int64  `gorm:"column:use_time;type:varchar(255)" json:"use_time"`
	Country      string `gorm:"column:country;type:varchar(255)" json:"country"`
	Region       string `gorm:"column:region;type:varchar(255)" json:"region"`
	Province     string `gorm:"column:province;type:varchar(255)" json:"province"`
	City         string `gorm:"column:city;type:varchar(255)" json:"city"`
	Isp          string `gorm:"column:isp;type:varchar(255)" json:"isp"`
}

func GetLogList(keyword string) *gorm.DB {
    tx := DB.Debug().Model(&SysLog{}).Select(
        "id, ip, browser, class_method, http_method, request_param, request_uri, result, status_code, use_time, country, region, province, city, isp, created_at, updated_at, deleted_at")
    
    if keyword != "" {
        tx = tx.Where(
            "ip LIKE ? OR request_uri LIKE ? OR class_method LIKE ? OR result LIKE ? OR browser LIKE ?",
            "%"+keyword+"%",
            "%"+keyword+"%", 
            "%"+keyword+"%",
            "%"+keyword+"%",
            "%"+keyword+"%",
        )
    }
    
    return tx
}