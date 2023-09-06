package stores

// import redis
import (
	"context"
)

type Rivia struct {
	redis *Redis
	ctx   context.Context
	db    int
}

func NewRivia(redis *Redis, db int) *Rivia {
	ctx := context.Background()
	return &Rivia{
		redis: redis,
		ctx:   ctx,
		db:    db,
	}
}

func (r *Rivia) Get(id string) string {
	data := r.redis.client.Get(r.ctx, id)

	return data.String()
}

func (r *Rivia) Set(id string, expires int64, data interface{}) error {
	// change db to r.db
	r.redis.client.Conn().Select(r.ctx, r.db)

	err := r.redis.client.Set(r.ctx, id, data, 0).Err()

	if nil != err {
		return err
	}

	return nil
}

func (r *Rivia) Remove(id string) error {

	return nil
}
