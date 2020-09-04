
# abel 来历
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;阿贝尔来自英语中abel，感觉和本人很像，所以起名abel。英文中解释：诚实、可靠而且喜欢摸索新事物。 个性严谨，有责任心，情绪稳定。社交能力强，容易相处。渴望了解他人的内心世界，适合与人打交道的工作。富有爱心，家庭责任感强。处理问题较犹豫。[更多动态请关注>>](https://www.api996.cn/p/im)

# 架构说明
- 架构图待补充...
- abel-im采用GO语言开发；
- 底层实现技术gorilla/websocket、存储目前使用mysql(后续会改进)简单存储。
- 采用ETCD服务注册发现
- gRpc方式消息扩散
- abel-im支持集群部署
- 自定义日志实现

# 架构思考和疑虑
- 目前系统还停留在起步阶段，我对go语言理解还不够深技术水平有限。所以在阿贝尔还有很多的改进和优化地方。作为一个高性能的Im即时通讯系统来说分布式部署是至关重要，怎样能做到更好的扩容和应对紧急情况该如何处理，希望能得到网友指点。
- 目前消息存储是实时存储，如果高并发这种100%不行。
- 单台机器这种长连接，连接数问题。
- 突然间网络不稳定，客户端如何保证消息100%接收成功（难道网络连接成功后去获取历史消息，个人认为不是最好的解决方法）
##### 虽然我们的阿贝尔还比较简单，但是我相信通过后续的努力和广大网友提出的宝贵建议，依然能够作出高可用的企业级产品这也是作者终极目标，虽然代码简单但是性能还是绝对的优秀。作者微信号：majun391 欢迎骚扰提出宝贵建议

# 功能说明
#### 已经实现功能
- 1、支持单聊
- 2、支持群聊
- 3、聊天记录存储、历史消息查询
#### 准备实现功能
- 1、发送表情、发送文件
- 2、集群实现
- 3、弹幕功能
- 4、已读消息、未读消息