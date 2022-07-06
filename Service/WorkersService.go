package Service

import (
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type Worker struct {
	Wg    *sync.WaitGroup
	Jobs  chan map[int]interface{}
	Mongo chan *mongo.Client
}

func (w *Worker) InitWorker(worker_number int) {
	for i := 0; i < worker_number; i++ {
		go func(i int, w *Worker) {

		}(i, w)
	}
}
