package process

type Info struct {
	// reader name, used to find reader from otkafka.ReaderMaker
	Name string
	// reader workers count
	ReadWorker int
	// batch workers count
	BatchWorker int
	// data size for batch processing
	BatchSize int
	// handler workers count
	HandleWorker int
	// the size of the data channel
	ChanSize int
}

func (i Info) name() string {
	if i.Name == "" {
		return "default"
	}
	return i.Name
}

func (i *Info) readWorker() int {
	if i.ReadWorker <= 0 {
		return 1
	}
	return i.ReadWorker
}

func (i *Info) batchWorker() int {
	if i.BatchWorker <= 0 {
		return 1
	}
	return i.BatchWorker
}

func (i *Info) batchSize() int {
	if i.BatchSize <= 0 {
		return 1
	}
	return i.BatchSize
}

func (i *Info) handleWorker() int {
	if i.HandleWorker <= 0 {
		return 1
	}
	return i.HandleWorker
}

func (i *Info) chanSize() int {
	if i.ChanSize <= 0 {
		return 100
	}
	return i.ChanSize
}
