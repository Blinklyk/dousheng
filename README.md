项目展示
基于gin + gRPC 完成所有接口实现，并进行压测，具体如下：
github 地址：https://github.com/Blinklyk/dousheng
This content is only supported in a Feishu Docs
配置要求
- go 1.16.1+
- MySQL
- 搭建Redis环境
- 七牛云oss
- 抖音客户端软件
  安装步骤
- 下载源码
- 配置MySQL、OSS、Redis等相关参数
- 启动服务
- 在客户端配置相关服务器地址
  技术说明
  WEB框架：gin
  ORM框架：GORM
  RPC框架：gRPC
  缓存：Redis
  鉴权：Jwt
  OSS：七牛云
  数据库：MySQL
  日志：Zap + lumberjack
  配置：Viper
  项目结构：
  .
  ├── README.md
  ├── config
  ├── config.yaml
  ├── controller
  │   ├── comment.go
  │   ├── common.go
  │   ├── favorite.go
  │   ├── feed.go
  │   ├── publish.go
  │   ├── relation.go
  │   └── user.go
  ├── global
  │   ├── app.go
  │   └── constants.go
  ├── go.mod
  ├── go.sum
  ├── initialize
  │   ├── config.go
  │   ├── gorm.go
  │   ├── init.go
  │   ├── internal
  │   │   ├── gorm.go
  │   │   └── logger.go
  │   ├── log.go
  │   ├── redis.go
  │   └── session.go
  ├── main.go
  ├── model
  │   ├── comment.go
  │   ├── favorite.go
  │   ├── follow.go
  │   ├── follower.go
  │   ├── request
  │   │   ├── comment.go
  │   │   ├── favorite.go
  │   │   ├── feed.go
  │   │   ├── publish.go
  │   │   ├── relation.go
  │   │   └── user.go
  │   ├── response
  │   │   ├── comment.go
  │   │   ├── common.go
  │   │   ├── favorite.go
  │   │   ├── feed.go
  │   │   ├── publish.go
  │   │   ├── relation.go
  │   │   └── user.go
  │   ├── user.go
  │   └── video.go
  ├── pb
  │   ├── build.sh
  │   ├── rpcComment
  │   │   ├── rpcComment.pb.go
  │   │   ├── rpcComment.proto
  │   │   └── rpcComment_grpc.pb.go
  │   ├── rpcFavorite
  │   │   ├── rpcFavorite.pb.go
  │   │   ├── rpcFavorite.proto
  │   │   └── rpcFavorite_grpc.pb.go
  │   ├── rpcFollow
  │   │   ├── rpcFollow.pb.go
  │   │   ├── rpcFollow.proto
  │   │   └── rpcFollow_grpc.pb.go
  │   ├── rpcUser
  │   │   ├── rpcUser.pb.go
  │   │   ├── rpcUser.proto
  │   │   └── rpcUser_grpc.pb.go
  │   └── rpcVideo
  │       ├── rpcVideo.pb.go
  │       ├── rpcVideo.proto
  │       └── rpcVideo_grpc.pb.go
  ├── public
  │   └── bear.mp4
  ├── repository
  │   └── db_init.go
  ├── router.go
  ├── server-comment
  │   └── main.go
  ├── server-favorite
  │   └── main.go
  ├── server-follow
  │   └── main.go
  ├── server-user
  │   └── main.go
  ├── server-video
  │   └── main.go
  ├── service
  │   ├── comment.go
  │   ├── common.go
  │   ├── favorite.go
  │   ├── feed.go
  │   ├── publish.go
  │   ├── relation.go
  │   └── user.go
  ├── storage
  │   └── logs
  │       └── douyin.log
  └── utils
  ├── directory.go
  ├── hash.go
  ├── jwt.go
  ├── oss.go
  ├── rpcdto
  │   ├── user.go
  │   └── video.go
  ├── snowflake.go
  └── verify
  └── verify.go


架构图：
[Image]
数据库图：
[Image]
版本Tag
v1.0
- 单机架构，基本实现所有接口逻辑
  v1.1
- Viper配置，实现修改yaml即可配置环境，便于多人协作
- 引入zap + lumberjack 日志框架，主要记录调试，错误，sql信息
- 数据库连接池配置
  v1.1.1
- 修改配置
  v1.1.2 - v1.1.4
- 引入DTO + varify
- 修改DTO思路
  v1.1.5
- cover_url 处理
- play_url 防重复优化
  v1.1.6
- 性能测试，goroutine 简单优化
  v1.1.7
- 分析sql表，创建数据库索引
  v1.2
- 服务拆分，后端service部分使用gRPC框架通信

性能测试(基于v1.1.6)
一：压测信息
This content is only supported in a Feishu Docs
二：压测模块
模块：
This content is only supported in a Feishu Docs
三：测压报告
总体流程测试如图所示，总共循环30整体周期，共有480个请求。
[Image]
3.1 关注模块
3.1.1 获取关注列表
[Image]
3.1.2 获取粉丝列表
[Image]
3.1.3 关注
[Image]
3.1.4 取关
[Image]
3.2 点赞模块
3.2.1 获取点赞列表
[Image]
3.2.2 点赞
[Image]
3.2.3 取消赞
[Image]
3.3 评论模块
3.3.1 获取评论列表
[Image]
3.3.2 评论
[Image]
3.3.3 删除评论
[Image]
3.4 视频模块
3.4.1 获取发布列表
[Image]

3.4.2 获取视频流
[Image]
3.5 用户模块
3.5.1 获取用户信息
[Image]
3.5.2 登录
[Image]
3.5.3 注册
[Image]
四 压测问题及分析
整体性能良好，但是在注册时候请求耗时可能有稍微偏高。后续会不断优化。
模块设计
用户模块
1、用户注册
接收参数：用户名、密码
返回参数：用户 ID 与 Token
type RegisterRequest struct {
Username string `json:"username" gorm:"not null; comment:username for register;" form:"username"`
Password string `json:"password" gorm:"not null; comment:password for register" form:"password"`
}

type RegisterResponse struct {
Response
UserId int64  `json:"user_id,omitempty"`
Token  string `json:"token"`
}
2、用户登录
尝试使用session + redis 的鉴权方案，但客户端不支持识别带有set-Cookies的Header的Response，所以最终选择jwt + redis 的搭配。登录时将部分用户信息存入redis中，每次活跃用户经过用户鉴权部分则会自动刷新 authentication 过期时间。
type LoginResponse struct {
Response
UserId int64  `json:"user_id,omitempty"`
Token  string `json:"token"`
}

type LoginRequest struct {
Username string `json:"username" gorm:"not null; comment:username for register;" form:"username"`
Password string `json:"password" gorm:"not null; comment:password for register" form:"password"`
}
3、用户信息
type User struct {
//gorm.Model
ID             int64          `gorm:"primarykey"` // 主键ID
CreatedAt      time.Time      // 创建时间
UpdatedAt      time.Time      // 更新时间
DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`                            // 删除时间
Name           string         `json:"name,omitempty" gorm:"default:testName"`    // TODO
FollowCount    int64          `json:"follow_count,omitempty" gorm:"default:0"`   // 关注数
FollowerCount  int64          `json:"follower_count,omitempty" gorm:"default:0"` // 粉丝数
IsFollow       bool           `json:"is_follow,omitempty" gorm:"default:false"`  // 当前用户是否关注
Username       string         `json:"username" gorm:"comment:username" `         // 登录账号
Password       string         `json:"password" gorm:"comment:password"`          // 登录密码
Videos         []Video        `json:"videos"`                                    // 发布视频列表
FavoriteVideos []Video        `json:"favorite_videos"`                           //`gorm:"many2many:favorite"`                        // 点赞视频列表
}

视频模块
1、视频feed流
不限制用户登录状态，返回投稿时间倒序的视频列表，单次最多返回20个视频
// 传入参数
type FeedRequest struct {
LatestTime string `json:"latest_time,omitempty"`
Token      string `json:"token" form:"token"`
}

type FeedResponse struct {
Response
VideoList []VideoDTO `json:"video_list"`
NextTime  int64      `json:"next_time"`
}

// 视频内容
type VideoDTO struct {
ID            int64           `gorm:"primarykey"`                                  // 主键ID
UserID        int64           `json:"author_id,omitempty"`                         // 发布作者
User          UserInfo        `json:"author,omitempty" gorm:"foreignKey:UserID"`   // user信息
PlayUrl       string          `json:"play_url,omitempty" gorm:"default:testName"`  // 视频地址
CoverUrl      string          `json:"cover_url,omitempty" gorm:"default:testName"` // 封面地址
FavoriteCount int64           `json:"favorite_count" gorm:"default:0"`             // 点赞数量
CommentCount  int64           `json:"comment_count" gorm:"default:0"`              // 评论数量
IsFavorite    bool            `json:"is_favorite" gorm:"default:false"`            // 是否点赞
PublishTime   time.Time       `json:"publish_time" gorm:"comment:投稿时间"`            // 投稿时间
Title         string          `json:"title, omitempty" gorm:"comment:视频说明"`        // 投稿时添加的文字
CommentList   []model.Comment `json:"comment_list,omitempty"`                      // 视频下的评论列表
}

2、视频投稿
传入传出参数：
// 传入参数
type PublishRequest struct {
Token string `json:"token" form:"token"`
Title string `json:"title" form:"title"`
}

// 返回参数
{
StatusCode    Integer
StatusMsg     string or null
}
Controller 层：
- 读取传过来的文件 ， token， title
- 返回响应
  Service 层：
- 使用雪花算法生成随机文件名
- 设置访问路径
- 编写七牛云上传逻辑， 截取第一秒为封面

3、用户投稿列表
点赞模块
- 登录用户对视频进行取消赞，首先检查用户是否登录，登录后才能进行取消赞操作
- 取得客户端对点赞或取消点赞的状态码，获取用户的id以及视频的id
  type FavoriteRequest struct {
  Token      string `json:"token" form:"token"`
  VideoID    string `json:"video_id" form:"video_id"`
  ActionType string `json:"action_type" form:"action_type"`
  }

// 返回参数
{
StatusCode    Integer
StatusMsg     string or null
}
评论模块
// 评论信息
type CommentInfo struct {
ID         int64    `gorm:"primarykey"` // 主键ID
UserID     int64    `json:"user_id" `
User       UserInfo `json:"user" gorm:"foreignKey:UserID"`
VideoID    int64    `json:"video_id"`
Content    string   `json:"content"`
CreateData string   `json:"create_data"`
}

// 评论请求
type CommentRequest struct {
Token       string `json:"token" form:"token"`
VideoID     string `json:"video_id" form:"video_id"`
ActionType  string `json:"action_type" form:"action_type"`
CommentText string `json:"comment_text" form:"comment_text"`
CommentID   string `json:"comment_id" form:"comment_id"`
}

// 评论响应
type CommentActionResponse struct {
Response
Comment CommentInfo `json:"comment"`
}

// 评论列表
关注模块
- 接口主要接收三个参数：当前用户id、目标用户id、动作（关注还是取关）
- 设计关注表和粉丝表，便于查询，分别添加status字段，用于判断是否互相关注。
- 每个用户保存两个集合列表（邻接表和逆邻接表的思想）
    - 我关注了谁
    - 谁关注了我
- 一次关注操作，同时保存两份数据：
    - 当前用户的关注信息。key样式：followee + 当前用户id。value：目标用户id
    - 目标用户的被关注信息。key样式：follower + 目标用户id。value：当前用户id
      // 关注列表请求
      type FollowListRequest struct {
      Token  string `json:"token" form:"token"`
      UserID string `json:"user_id" form:"user_id"`
      }

// 关注列表响应
type FollowListResponse struct {
Response
UserList []UserInfo `json:"user_list"`
}

// 粉丝列表请求参数
type FollowerListRequest struct {
Token  string `json:"token" form:"token"`
UserID string `json:"user_id" form:"user_id"`
}

// 粉丝列表响应
type FollowerListResponse struct {
Response
UserList []UserInfo `json:"user_list"`
}
编译项目到linux
set GOARCH=amd64
set GOOS=linux
go build ./

chmod 777 main
./main
未来展望（持续优化中...）
消息队列
使用Rabbit MQ，异步、解耦、削峰。
创建多个MQ，分别对关系、评论和点赞操作时对所传输过程中保存消息的容器，针对关系建立关注和取关两个队列；评论操作则添加删除评论的队列；对点赞操作添加点赞和取消两个队列。可以大幅度提升并发场景下的处理速度。
容器化
用Docker容器化应用程序，方便与依赖打包、版本控制，并且达到隔离性、可移植、轻量高效和安全性的好处。
gRPC网关配置 + 负载均衡
可以通过可插拔的支持来有效地连接数据中心内和跨数据中心的服务，以实现负载平衡，跟踪，运行状况检查和身份验证。充分发挥分布式架构优点。
[Image]
GRPC服务端创建步骤：
- NettyServer 实例创建；
- 绑定 IDL 定义的服务接口实现类；
- gRPC 服务实例（ServerImpl）构建。
  GRPC服务端service调用流程：
- gRPC 请求消息接入；
- gRPC 消息头和消息体处理；
- 内部的服务路由和调用；
- 响应消息发送。
  Contributors
- 李云开 邮箱：290913796@qq.com
- 刘志君 邮箱：695532288@qq.com
- 高志豪 邮箱：gzhh0216@gmail.com