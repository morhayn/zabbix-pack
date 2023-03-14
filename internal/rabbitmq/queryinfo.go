package rabitmq

// QueueInfo represents a queue, its properties and key metrics.
type QueueInfo struct {
	// Queue name
	Name string `json:"name"`
	// Queue type
	Type string `json:"type,omitempty"`
	// Virtual host this queue belongs to
	Vhost string `json:"vhost,omitempty"`
	// Is this queue durable?
	Durable bool `json:"durable"`
	// Is this queue auto-deleted?
	AutoDelete bool `json:"auto_delete"`
	// Is this queue exclusive?
	Exclusive bool `json:"exclusive,omitempty"`
	// Extra queue arguments
	Arguments map[string]interface{} `json:"arguments"`

	// RabbitMQ node that hosts master for this queue
	Node string `json:"node,omitempty"`
	// Queue status
	Status string `json:"state,omitempty"`
	// Queue leader when it is quorum queue
	Leader string `json:"leader,omitempty"`
	// Queue members when it is quorum queue
	Members []string `json:"members,omitempty"`
	// Queue online members when it is quorum queue
	Online []string `json:"online,omitempty"`

	// Total amount of RAM used by this queue
	Memory int64 `json:"memory,omitempty"`
	// How many consumers this queue has
	Consumers int `json:"consumers,omitempty"`
	// Detail information of consumers
	ConsumerDetails *[]ConsumerDetail `json:"consumer_details,omitempty"`
	// Utilisation of all the consumers
	ConsumerUtilisation float64 `json:"consumer_utilisation,omitempty"`
	// If there is an exclusive consumer, its consumer tag
	ExclusiveConsumerTag string `json:"exclusive_consumer_tag,omitempty"`

	// GarbageCollection metrics
	GarbageCollection *GarbageCollectionDetails `json:"garbage_collection,omitempty"`

	// Policy applied to this queue, if any
	Policy string `json:"policy,omitempty"`

	// Total bytes of messages in this queues
	MessagesBytes               int64 `json:"message_bytes,omitempty"`
	MessagesBytesPersistent     int64 `json:"message_bytes_persistent,omitempty"`
	MessagesBytesRAM            int64 `json:"message_bytes_ram,omitempty"`
	MessagesBytesReady          int64 `json:"message_bytes_ready,omitempty"`
	MessagesBytesUnacknowledged int64 `json:"message_bytes_unacknowledged,omitempty"`

	// Total number of messages in this queue
	Messages           int          `json:"messages,omitempty"`
	MessagesDetails    *RateDetails `json:"messages_details,omitempty"`
	MessagesPersistent int          `json:"messages_persistent,omitempty"`
	MessagesRAM        int          `json:"messages_ram,omitempty"`

	// Number of messages ready to be delivered
	MessagesReady        int          `json:"messages_ready,omitempty"`
	MessagesReadyDetails *RateDetails `json:"messages_ready_details,omitempty"`

	// Number of messages delivered and pending acknowledgements from consumers
	MessagesUnacknowledged        int          `json:"messages_unacknowledged,omitempty"`
	MessagesUnacknowledgedDetails *RateDetails `json:"messages_unacknowledged_details,omitempty"`

	MessageStats *MessageStats `json:"message_stats,omitempty"`

	OwnerPidDetails *OwnerPidDetails `json:"owner_pid_details,omitempty"`

	BackingQueueStatus *BackingQueueStatus `json:"backing_queue_status,omitempty"`

	ActiveConsumers int64 `json:"active_consumers,omitempty"`
}

// MessageStats fields repsent a number of metrics related to published messages
type MessageStats struct {
	Publish                 int64       `json:"publish"`
	PublishDetails          RateDetails `json:"publish_details"`
	Deliver                 int64       `json:"deliver"`
	DeliverDetails          RateDetails `json:"deliver_details"`
	DeliverNoAck            int64       `json:"deliver_noack"`
	DeliverNoAckDetails     RateDetails `json:"deliver_noack_details"`
	DeliverGet              int64       `json:"deliver_get"`
	DeliverGetDetails       RateDetails `json:"deliver_get_details"`
	Redeliver               int64       `json:"redeliver"`
	RedeliverDetails        RateDetails `json:"redeliver_details"`
	Get                     int64       `json:"get"`
	GetDetails              RateDetails `json:"get_details"`
	GetNoAck                int64       `json:"get_no_ack"`
	GetNoAckDetails         RateDetails `json:"get_no_ack_details"`
	Ack                     int64       `json:"ack"`
	AckDetails              RateDetails `json:"ack_details"`
	ReturnUnroutable        int64       `json:"return_unroutable"`
	ReturnUnroutableDetails RateDetails `json:"return_unroutable_details"`
	DropUnroutable          int64       `json:"drop_unroutable"`
	DropUnroutableDetails   RateDetails `json:"drop_unroutable_details"`
}

type BackingQueueStatus struct {
	Q1 int `json:"q1,omitempty"`
	Q2 int `json:"q2,omitempty"`
	Q3 int `json:"q3,omitempty"`
	Q4 int `json:"q4,omitempty"`
	// Total queue length
	Length int64 `json:"len,omitempty"`
	// Number of pending acks from consumers
	PendingAcks int64 `json:"pending_acks,omitempty"`
	// Number of messages held in RAM
	RAMMessageCount int64 `json:"ram_msg_count,omitempty"`
	// Number of outstanding acks held in RAM
	RAMAckCount int64 `json:"ram_ack_count,omitempty"`
	// Number of persistent messages in the store
	PersistentCount int64 `json:"persistent_count,omitempty"`
	// Average ingress (inbound) rate, not including messages
	// that straight through to auto-acking consumers.
	AverageIngressRate float64 `json:"avg_ingress_rate,omitempty"`
	// Average egress (outbound) rate, not including messages
	// that straight through to auto-acking consumers.
	AverageEgressRate float64 `json:"avg_egress_rate,omitempty"`
	// rate at which unacknowledged message records enter RAM,
	// e.g. because messages are delivered requiring acknowledgement
	AverageAckIngressRate float32 `json:"avg_ack_ingress_rate,omitempty"`
	// rate at which unacknowledged message records leave RAM,
	// e.g. because acks arrive or unacked messages are paged out
	AverageAckEgressRate float32 `json:"avg_ack_egress_rate,omitempty"`
}

// OwnerPidDetails describes an exclusive queue owner (connection).
type OwnerPidDetails struct {
	Name     string `json:"name,omitempty"`
	PeerPort int    `json:"peer_port,omitempty"`
	PeerHost string `json:"peer_host,omitempty"`
}

// ConsumerDetail describe consumer information with a queue
type ConsumerDetail struct {
	Arguments      map[string]interface{} `json:"arguments"`
	ChannelDetails ChannelDetails         `json:"channel_details"`
	AckRequired    bool                   `json:"ack_required"`
	Active         bool                   `json:"active"`
	ActiveStatus   string                 `json:"active_status"`
	ConsumerTag    string                 `json:"consumer_tag"`
	Exclusive      bool                   `json:"exclusive,omitempty"`
	PrefetchCount  uint                   `json:"prefetch_count"`
	Queue          QueueDetail            `json:"queue"`
}

// ChannelDetails describe channel information with a consumer
type ChannelDetails struct {
	ConnectionName string `json:"connection_name"`
	Name           string `json:"name"`
	Node           string `json:"node"`
	Number         uint   `json:"number"`
	PeerHost       string `json:"peer_host"`
	PeerPort       uint   `json:"peer_port"`
	User           string `json:"user"`
}
type GarbageCollectionDetails struct {
	FullSweepAfter  int `json:"fullsweep_after"`
	MaxHeapSize     int `json:"max_heap_size"`
	MinBinVheapSize int `json:"min_bin_vheap_size"`
	MinHeapSize     int `json:"min_heap_size"`
	MinorGCs        int `json:"minor_gcs"`
}

// QueueDetail describe queue information with a consumer
type QueueDetail struct {
	Name  string `json:"name"`
	Vhost string `json:"vhost,omitempty"`
}

// RateDetailSample single touple
type RateDetailSample struct {
	Sample    int64 `json:"sample"`
	Timestamp int64 `json:"timestamp"`
}
type RateDetails struct {
	Rate    float32            `json:"rate"`
	Samples []RateDetailSample `json:"samples"`
}
