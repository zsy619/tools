package xuuid

import (
	"errors"
	"sync"
	"time"
)

const (
	// 自定义纪元起点——时间偏移量，单位为毫秒
	// 即 2014 年 10 月 25 日 05:06:02.373 UTC，恰好等于 sqrt(2) 的前几位数字
	epoch int64 = 1414213562373

	// 分配给服务器 ID 的位数（最大 1023）
	numWorkerBits = 10
	// 分配给每毫秒计数器的位数
	numSequenceBits = 12

	// workerId 掩码
	maxWorkerId = -1 ^ (-1 << numWorkerBits)
	// 序号掩码
	maxSequence = -1 ^ (-1 << numSequenceBits)
)

// SnowFlake 是保存雪花算法（Snowflake）相关数据的状态结构体。
type SnowFlake struct {
	lastTimestamp uint64
	sequence      uint32
	workerId      uint32
	lock          sync.Mutex
}

// Pack 将各字段按位打包成一个雪花 ID 值。
func (sf *SnowFlake) pack() uint64 {
	return (sf.lastTimestamp << (numWorkerBits + numSequenceBits)) |
		(uint64(sf.workerId) << numSequenceBits) |
		(uint64(sf.sequence))
}

// NewSnowFlake 使用给定的 workerId 初始化雪花 ID 生成器。
func NewSnowFlake(workerId uint32) (*SnowFlake, error) {
	if workerId > maxWorkerId {
		return nil, errors.New("invalid worker Id")
	}
	return &SnowFlake{workerId: workerId}, nil
}

// Next 生成下一个唯一的 ID。
func (sf *SnowFlake) Next() (uint64, error) {
	sf.lock.Lock()
	defer sf.lock.Unlock()

	ts := timestamp()
	if ts == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & maxSequence
		if sf.sequence == 0 {
			ts = sf.waitNextMilli(ts)
		}
	} else {
		sf.sequence = 0
	}

	if ts < sf.lastTimestamp {
		return 0, errors.New("invalid system clock")
	}
	sf.lastTimestamp = ts
	return sf.pack(), nil
}

// 序列号已耗尽，等待下一个毫秒到来。
func (sf *SnowFlake) waitNextMilli(ts uint64) uint64 {
	for ts == sf.lastTimestamp {
		time.Sleep(100 * time.Microsecond)
		ts = timestamp()
	}
	return ts
}

func timestamp() uint64 {
	// 将纳秒转换为毫秒，并扣除自定义纪元起点。
	return uint64(time.Now().UnixNano()/int64(1000000) - epoch)
}
