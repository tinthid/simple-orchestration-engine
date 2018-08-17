package redis

func (r *Redis) GetByKey(key string) (string, error) {
	val, err := r.RedisClient.Get(key).Result()
	return val, err
}

func (r *Redis) SetByKey(key string, val string) (error) {
	err := r.RedisClient.Set(key, val, 0).Err()
	return err
}
