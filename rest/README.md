# REST - Representational State Transfer (具象状态传输)

# 描述

## 架构设计
- 建造者(Builder): 把具体的数据来源（channel）构造成executor
- 执行者(Executor): 业务逻辑执行者
- 数据通道(Cmd): 数据通道基类
- rest服务(RestServer): 服务启动器

### 建造者(Builder)


### 执行者(Executor)
- Executor.Prepare(): 预处理阶段.加载全局 dbCli, 自定义 log 等操作.
- Executor.Exec(): 逻辑处理阶段.
- Executor.Finish(): 构造自定义response返回.
- Executor.RestCmd(): 返回自身的 Cmd 实例指针