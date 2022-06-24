# 开发目标流程
<!-- 1. gin生成项目(没有脚手架),目录构造 -->
<!-- 1. 简单路由 -->
2. 配置文件
    https://github.com/spf13/viper
3. controller与service的关联
2. db
    * gorm迁移
3. redis
4. 日志
5. auth
6. websocket
7. docker
8. 整理,markdown
9. 测试提供

10. go微服务实现,账号鉴权系统
***
# GO

先简单编写,再分层  
帧同步匹配,转发而已

1. 鉴权(外部服务)  
2. 登录(分配标识token,没有鉴权就默认分配一个即可)
3. 大厅
    1. 房间列表
    2. 房间搜索
4. 房间(相同channel)
    1. 加入

5. websocket 
    1. 接收,推送消息    

1. channel内所有人管理
    * 房间
    * 对战开始