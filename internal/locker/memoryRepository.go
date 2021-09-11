package locker

var memoryRepository map[LockerId]Locker

func init() {
	memoryRepository = make(map[LockerId]Locker)
}
