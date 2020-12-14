module github.com/panda8z/istone

go 1.14

replace (
	github.com/panda8z/istone/cmd => ./cmd
	github.com/panda8z/istone/src/pkg/app => ./src/pkg/app
	github.com/panda8z/istone/src/pkg/captcha => ./src/pkg/captcha
	github.com/panda8z/istone/src/pkg/casbin => ./src/pkg/casbin
	github.com/panda8z/istone/src/pkg/config => ./src/pkg/config
	github.com/panda8z/istone/src/pkg/cronjob => ./src/pkg/cronjob
	github.com/panda8z/istone/src/pkg/dbservice => ./src/pkg/dbservice
	github.com/panda8z/istone/src/pkg/dto => ./src/pkg/dto
	github.com/panda8z/istone/src/pkg/global => ./src/pkg/global
	github.com/panda8z/istone/src/pkg/jwt => ./src/pkg/jwt
	github.com/panda8z/istone/src/pkg/log => ./src/pkg/log
	github.com/panda8z/istone/src/pkg/logger => ./src/pkg/logger
	github.com/panda8z/istone/src/pkg/middleware => ./src/pkg/middleware
	github.com/panda8z/istone/src/pkg/models => ./src/pkg/models
	github.com/panda8z/istone/src/pkg/search => ./src/pkg/search
	github.com/panda8z/istone/src/pkg/security => ./src/pkg/security
	github.com/panda8z/istone/src/pkg/tools => ./src/pkg/tools
	github.com/panda8z/istone/src/pkg/trace => ./src/pkg/trace
	github.com/panda8z/istone/src/pkg/utils => ./src/pkg/utils

	github.com/panda8z/istone/src/admin/models => ./src/admin
)

require (
	github.com/alibaba/sentinel-golang v0.6.1
	github.com/casbin/casbin/v2 v2.11.2
	github.com/casbin/gorm-adapter/v3 v3.0.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2
	github.com/mojocn/base64Captcha v1.3.1
	github.com/opentracing/basictracer-go v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pingcap/log v0.0.0-20201112100606-8f1e84a3abc8
	github.com/robfig/cron/v3 v3.0.1
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9
	gorm.io/driver/mysql v0.3.0
	gorm.io/driver/postgres v0.2.6
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.8
	sourcegraph.com/sourcegraph/appdash v0.0.0-20190731080439-ebfcffb1b5c0
	sourcegraph.com/sourcegraph/appdash-data v0.0.0-20151005221446-73f23eafcf67 // indirect
)
