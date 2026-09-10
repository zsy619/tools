package xmath

import (
	"errors"
	"sync"
	"time"
)

const (
	// numberBits 表示每个集群下的每个节点，1 毫秒内可生成的 id 序号的二进制位数（对应 Snowflake ID 的最低位段）。
	numberBits uint8 = 12
	// workerBits 每台机器（节点）的 ID 位数。10 位最大可有 2^10 = 1024 个节点，单节点每毫秒最多生成 2^12 = 4096 个唯一 ID。
	workerBits uint8 = 10
	// 这里求最大值使用了位运算，-1 的二进制表示为 1 的补码。
	workerMax int64 = -1 ^ (-1 << workerBits) // 节点 ID 的最大值，用于防止溢出
	numberMax int64 = -1 ^ (-1 << numberBits) // 同上，表示生成 id 序号的最大值
	timeShift uint8 = workerBits + numberBits // 时间戳向左的偏移量
	workerShift uint8 = numberBits            // 节点 ID 向左的偏移量
	// 41 位字节作为时间戳数值的话，大约 68 年就会用完。
	// 一旦定义并开始生成 ID，请勿修改 epoch，否则可能生成重复 ID。
	epoch int64 = 1525705533000 // epoch 常量对应的时间戳（毫秒）
)

// Worker 雪花算法 ID 生成器实例。
//
// 字段说明：
//   - mu: 互斥锁，用于保证并发安全。
//   - timestamp: 记录上一次生成 ID 的时间戳（毫秒）。
//   - workerId: 当前节点的 ID。
//   - number: 当前毫秒已生成的 ID 序号（从 0 开始累加）。
type Worker struct {
	mu        sync.Mutex // 添加互斥锁 确保并发安全
	timestamp int64      // 记录上一次生成id的时间戳
	workerId  int64      // 该节点的ID
	number    int64      // 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
}

// NewWorkerDefault
/**
 * @description: 生成一个新的节点
 * @param {int64} workerId
 * @return {*}
 */
// NewWorkerDefault 创建一个新的 Worker 节点。
//
// 参数：
//   - workerId: 节点 ID，必须在 [0, workerMax] 范围内。
//
// 返回值：
//   - *Worker: 初始化完成的 Worker 实例。
//   - error: 当 workerId 越界时返回错误。
func NewWorkerDefault(workerId int64) (*Worker, error) {
	// 要先检测workerId是否在上面定义的范围内
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

// GetId
/**
 * @description: 获取下一个ID
 * @return {*}
 */
// GetId 生成并返回下一个唯一 ID。
//
// 返回值：基于时间戳、节点 ID 与本节点序号生成的 int64 唯一 ID。
//
// 副作用：内部使用互斥锁，可能在单毫秒内超过 numberMax 上限时空转等待至下一毫秒。
func (w *Worker) GetId() int64 {
	// 获取id最关键的一点 加锁 加锁 加锁
	w.mu.Lock()
	defer w.mu.Unlock() // 生成完成后记得 解锁 解锁 解锁

	// 获取生成时的时间戳
	now := time.Now().UnixNano() / 1e6 // 纳秒转毫秒
	if w.timestamp == now {
		w.number++

		// 这里要判断，当前工作节点是否在1毫秒内已经生成numberMax个ID
		if w.number > numberMax {
			// 如果当前工作节点在1毫秒内生成的ID已经超过上限 需要等待1毫秒再继续生成
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 如果当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		w.number = 0
		// 下面这段代码看到很多前辈都写在if外面，无论节点上次生成id的时间戳与当前时间是否相同 都重新赋值  这样会增加一丢丢的额外开销 所以我这里是选择放在else里面
		w.timestamp = now // 将机器上一次生成ID的时间更新为当前时间
	}

	ID := int64((now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}
