# Section-01-configs


## 需要的配置清单

```yaml
settings:
  application:
    # dev开发环境 test测试环境 prod线上环境
    mode: dev
    # 服务器ip，默认使用 0.0.0.0
    host: 0.0.0.0
    # 服务名称
    name: testApp
    # 端口号
    port: 9000 # 服务端口号
    readtimeout: 1
    writertimeout: 2
    # 数据权限功能开关
    enabledp: false
  logger:
    # 日志存放路径
    path: temp/logs
    # 控制台日志
    stdout: true
    # 日志等级
    level: all
    # 业务日志开关
    enabledbus: true
    # 请求日志开关
    enabledreq: false
    # 数据库日志开关 dev模式，将自动开启
    enableddb: false
    # 自动任务日志开关 dev模式，将自动开启
    enabledjob: false
  jwt:
    # token 密钥，生产环境时及的修改
    secret: istone-9527
    # token 过期时间 单位：秒
    timeout: 3600
  database:
    # 数据库类型 mysql
    driver: mysql
    # 数据库连接字符串 mysql 缺省信息 charset=utf8&parseTime=True&loc=Local&timeout=1000ms
    source: user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local&timeout=1000ms

```