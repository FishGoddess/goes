## ✒ 历史版本的特性介绍 (Features in old versions)

### v0.2.5-alpha

> 此版本发布于 2025-06-30

* 自动清理 worker 机制，避免浪费 worker 资源

### v0.2.4-alpha

> 此版本发布于 2025-06-28

* 支持查询可用的 worker 数量
* 懒启动 worker 机制，避免浪费 worker 资源

### v0.2.3-alpha

> 此版本发布于 2025-06-28

* 调整 worker 调度层

### v0.2.2-alpha

> 此版本发布于 2025-06-26

* 抽象出 workers 接口层
* 增加随机调度策略

### v0.2.1-alpha

> 此版本发布于 2025-06-25

* 支持选择 sync.Mutex 作为锁实现
* 单元测试覆盖率提高到 95%

### v0.2.0-alpha

> 此版本发布于 2025-06-25

* 增加异步任务执行器
* 支持轮询调度策略
* 支持设置 worker 任务队列大小
* 支持设置 panic 处理函数
* 支持自定义 sync.Locker 实现
* 单元测试覆盖率提高到 94%

### v0.1.0

> 此版本发布于 2025-06-21

* 增加支持退避策略的自旋锁

### v0.0.1

> 此版本发布于 2023-03-03

* 基础功能完善