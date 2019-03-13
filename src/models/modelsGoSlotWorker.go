package models

// 存放 goSlotWorker.go 使用到的 struct

type ConfWorker struct {
	AmqpHost string		`json:"amqpHost"`
	RedisHost string	`json:"redisHost"`
	RedisPlatformPrefix string `json:"redisPlatformPrefix"`
	RedisSlotPrefix string `json:"redisSlotPrefix"`
	RedisLoginPrefix string `json:"redisLoginPrefix"`
	MgoHost string		`json:"mgoHost"`
	NameOfWorkerQueue []string	`json:"nameOfWorkerQueue"`
	NumberOfWorkerQueue []int		`json:"numberOfWorkerQueue"`
}
