package manager

import (
	"cs425_mp1/internal/config"
	"net"
	"sync"
)

type Manager struct {
    self      config.NodeInfo
    peers     map[string]net.Conn   // nodeID -> TCP connection
    mu        sync.Mutex            // protects peers map
    inbox     chan Message           // incoming messages for the node to consume
    failures  chan string            // IDs of peers that died
    listener  net.Listener
}

func NewManager(self config.NodeInfo, inboxSize int) *Manager {
    return &Manager{
        self:     self,
        peers:    make(map[string]net.Conn),
        inbox:    make(chan Message, inboxSize),
        failures: make(chan string, inboxSize),
    }
}

func (m *Manager) Listen() error { ... }
func (m *Manager) ConnectToPeers(nodes []config.NodeInfo) error { ... }
func (m *Manager) Broadcast(msg Message) { ... }
func (m *Manager) Send(nodeID string, msg Message) error { ... }
func (m *Manager) Inbox() <-chan Message { ... }
func (m *Manager) Failures() <-chan string { ... }
func (m *Manager) Close() { ... }