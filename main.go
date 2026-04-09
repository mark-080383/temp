package main

import (
    // beego 相關
    _ "github.com/beego/beego/v2/core/logs"
    _ "github.com/beego/beego/v2/server/web"
    _ "github.com/beego/beego/v2/server/web/context"
    _ "github.com/beego/beego/v2/task"

    // GORM 相關
    _ "gorm.io/driver/sqlserver"
    _ "gorm.io/gorm"

    // 其他套件
    _ "github.com/fogleman/gg"
    _ "github.com/google/uuid"
    _ "github.com/jinzhu/now"
    _ "github.com/satori/go.uuid"
    _ "github.com/signintech/gopdf"
    _ "github.com/xuri/excelize/v2"

    // golang.org/x 套件
    _ "golang.org/x/text/encoding/traditionalchinese"
    _ "golang.org/x/text/encoding/unicode"
    _ "golang.org/x/text/transform"
)

func main() {}
