package groupcache

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/adrian-lin-1-0-0/groupcache/discovery"
	etcd "go.etcd.io/etcd/client/v3"
)

var (
	ErrServiceNameEmpty     = errors.New("service name is empty")
	ErrServiceEndpointEmpty = errors.New("service endpoint is empty")
	ErrGetServicesFailed    = errors.New("get services failed")
)

type Client struct {
	client    *etcd.Client
	self      discovery.Service
	ctx       context.Context
	ttl       time.Duration
	heartbeat time.Duration
	leaseID   etcd.LeaseID
	once      sync.Once
}

const (
	defaultTTL       = 10 * time.Second
	defaultHeartbeat = 3 * time.Second
)

var _ discovery.Client = (*Client)(nil)

func NewClientWithEtcdClient(ctx context.Context, cli *etcd.Client, service discovery.Service) (*Client, error) {

	if service.Name == "" {
		return nil, ErrServiceNameEmpty
	}

	if service.Endpoint == "" {
		return nil, ErrServiceEndpointEmpty
	}

	return &Client{
		client:    cli,
		self:      service,
		ctx:       ctx,
		ttl:       defaultTTL,
		heartbeat: defaultHeartbeat,
		once:      sync.Once{},
	}, nil
}

func NewEtcdClient(ctx context.Context, endpoints []string, service discovery.Service) (*Client, error) {

	if service.Name == "" {
		return nil, ErrServiceNameEmpty
	}

	if service.Endpoint == "" {
		return nil, ErrServiceEndpointEmpty
	}

	cli, err := etcd.New(etcd.Config{
		Endpoints:   endpoints,
		Context:     ctx,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		client:    cli,
		self:      service,
		ctx:       ctx,
		ttl:       defaultTTL,
		heartbeat: defaultHeartbeat,
		once:      sync.Once{},
	}, nil
}

func (c *Client) GetServices(basePath string) ([]discovery.Service, error) {
	resp, err := c.client.Get(c.ctx, basePath, etcd.WithPrefix())
	if err != nil {
		return nil, ErrGetServicesFailed
	}
	var services []discovery.Service
	for _, kv := range resp.Kvs {
		if string(kv.Key) == c.self.Name {
			continue
		}
		services = append(services, discovery.Service{
			Name:     string(kv.Key),
			Endpoint: string(kv.Value),
		})
	}
	return services, nil
}

func (c *Client) WatchPrefix(basePath string) (discovery.EventChan, error) {
	watchChan := c.client.Watch(c.ctx, basePath, etcd.WithPrefix())
	eventChan := make(chan discovery.Event)
	go func() {
		for resp := range watchChan {
			for _, ev := range resp.Events {
				var event discovery.Event
				switch ev.Type {
				case etcd.EventTypePut:
					event.Type = discovery.EventPut
				case etcd.EventTypeDelete:
					event.Type = discovery.EventDelete
				}
				event.Key = string(ev.Kv.Key)
				event.Value = string(ev.Kv.Value)
				eventChan <- event
			}
		}
	}()
	return eventChan, nil
}

func (c *Client) InitLease(ttl int64) error {
	lease := etcd.NewLease(c.client)
	leaseResp, err := lease.Grant(c.ctx, ttl)
	if err != nil {
		return fmt.Errorf("grant lease failed: %v", err)
	}
	c.leaseID = leaseResp.ID
	return nil
}

func (c *Client) Register(service discovery.Service) error {
	c.once.Do(func() {
		err := c.InitLease(int64(c.ttl.Seconds()))
		if err != nil {
			return
		}
		go c.loop()
	})

	if c.leaseID == 0 {
		panic(errors.New("leaseID is nil"))
	}

	_, err := c.client.Put(c.ctx, service.Name, service.Endpoint, etcd.WithLease(c.leaseID))

	if err != nil {
		panic("register service failed :" + err.Error())
	}

	return err
}

func (c *Client) Unregister(service discovery.Service) error {
	_, err := c.client.Delete(c.ctx, service.Name)
	return err
}

func (c *Client) loop() {
	tick := time.NewTicker(c.heartbeat)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			c.keepAlive()
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Client) keepAlive() {
	lease := etcd.NewLease(c.client)
	leaseKeepAliveQueue, err := lease.KeepAlive(c.ctx, c.leaseID)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case _, ok := <-leaseKeepAliveQueue:
			if !ok {
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}
