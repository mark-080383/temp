module blweb

go 1.22

require (
	github.com/beego/beego/v2 v2.3.10
	github.com/fogleman/gg v1.3.0
	github.com/google/uuid v1.6.0
	github.com/jinzhu/now v1.1.5
	github.com/satori/go.uuid v1.2.0
	github.com/signintech/gopdf v0.28.0
	github.com/xuri/excelize/v2 v2.9.0
	golang.org/x/text v0.25.0
	gorm.io/driver/sqlserver v1.5.4
	gorm.io/gorm v1.25.12
)

// 以下為間接相依套件，在可連線外網的電腦執行 go mod tidy 後會自動更新
// 目標最新安全版本：
// - golang.org/x/crypto v0.49.0
// - golang.org/x/image v0.38.0  
// - golang.org/x/net v0.52.0
// - golang.org/x/sys v0.33.0
