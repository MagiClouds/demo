package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	var wg sync.WaitGroup

	wg.Add(2)
	lockKey := "/lock"

	go func() {
		defer wg.Done()
		session, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		m := concurrency.NewMutex(session, lockKey)
		if err := m.Lock(context.TODO()); err != nil {
			log.Fatal("go1 get mutex failed " + err.Error())
		}
		fmt.Printf("go1 get mutex sucess\n")
		fmt.Println(m)
		time.Sleep(time.Duration(10) * time.Second)
		m.Unlock(context.TODO())
		fmt.Printf("go1 release lock\n")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(2) * time.Second)
		session, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		m := concurrency.NewMutex(session, lockKey)
		if err := m.Lock(context.TODO()); err != nil {
			log.Fatal("go2 get mutex failed " + err.Error())
		}
		fmt.Printf("go2 get mutex sucess\n")
		fmt.Println(m)
		time.Sleep(time.Duration(2) * time.Second)
		m.Unlock(context.TODO())
		fmt.Printf("go2 release lock\n")
	}()

	wg.Wait()

	fmt.Println(<-c)
}
