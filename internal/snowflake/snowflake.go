package snowflake

import "strconv"

type Snowflake struct {
	ID        uint64
	Timestamp uint64
	WorkerID  uint8
	ProcessID uint8
	Increment uint16
}

func ParseSnowflake(snowflakeID uint64) Snowflake {
	return Snowflake{
		ID:        snowflakeID,
		Timestamp: (snowflakeID >> 22) + 1420070400000,
		WorkerID:  uint8((snowflakeID >> 17) & 0x1F),
		ProcessID: uint8((snowflakeID >> 12) & 0x1F),
		Increment: uint16(snowflakeID & 0xFFF),
	}
}

func ParseSnowflakeString(snowflakeID string) Snowflake {
	snowflakeIDUint64, err := strconv.ParseUint(snowflakeID, 10, 64)

	if err != nil {
		return Snowflake{}
	}

	return Snowflake{
		ID:        snowflakeIDUint64,
		Timestamp: (snowflakeIDUint64 >> 22) + 1420070400000,
		WorkerID:  uint8((snowflakeIDUint64 >> 17) & 0x1F),
		ProcessID: uint8((snowflakeIDUint64 >> 12) & 0x1F),
		Increment: uint16(snowflakeIDUint64 & 0xFFF),
	}
}
