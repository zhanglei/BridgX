package clients

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/galaxy-future/BridgX/config"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var syncLocker *SyncLocker
var ErrReviewFailed = errors.New("review key already exists")

type SyncLocker struct {
	etcdClient *clientv3.Client
}

func initSyncLocker(config *config.EtcdConfig) error {
	syncLocker = &SyncLocker{}
	if config != nil {
		etcdClient, err := clientv3.New(clientv3.Config{
			Endpoints:   config.Endpoints,
			DialTimeout: config.DailTimeout,
		})
		if err != nil {
			return err
		}

		err = pingEtcdServer(etcdClient)
		if err != nil {
			return fmt.Errorf("failed to ping etcd :%w", err)
		}
		syncLocker.etcdClient = etcdClient
	}
	return nil
}

func SyncRun(TTL int, key string, job func() error) error {
	//if not config etcd client , just run this job
	if syncLocker.etcdClient == nil {
		return job()
	}
	// lock and run this job
	session, err := concurrency.NewSession(syncLocker.etcdClient, concurrency.WithTTL(TTL))
	if err != nil {
		return err
	}
	defer func(session *concurrency.Session) {
		_ = session.Close()
	}(session)

	l := concurrency.NewMutex(session, key)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := l.TryLock(ctx); err != nil {
		return err
	}

	err = job()

	_ = l.Unlock(context.Background())
	return err
}
func pingEtcdServer(etcdClient *clientv3.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := etcdClient.Put(ctx, "ping", "pong")
	if err != nil {
		return err
	}
	return nil
}
