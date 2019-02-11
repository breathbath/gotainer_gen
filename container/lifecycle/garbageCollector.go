package lifecycle

type GarbageCollector interface {
	Destroy() error
}
