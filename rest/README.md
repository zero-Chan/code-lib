# REST - Representational State Transfer (具象状态传输)

# 描述

## 架构设计
- 建造者(Builder): 把具体的数据来源（channel）构造成executor
- 执行者(Executor): 业务逻辑执行者
- 数据通道(RestChannel): 数据通道基类
- rest服务(RestServer): 服务启动器
