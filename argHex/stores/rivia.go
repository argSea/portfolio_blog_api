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
	// change db to r.db
	r.redis.client.Conn().Select(r.ctx, r.db)

	data := r.redis.client.Get(r.ctx, id)

	return data.String()
}

func (r *Rivia) Set(id string, data interface{}) error {

	return nil
}

func (r *Rivia) Remove(id string) error {

	return nil
}
