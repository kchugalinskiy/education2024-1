package wg

type WaitGroup struct {
	exitCh chan interface{}
}

func New() *WaitGroup {
	return &WaitGroup{
		exitCh: make(chan interface{}),
	}
}

func (wg *WaitGroup) Wait() {
	<-wg.exitCh
}

func (wg *WaitGroup) Done() {
	wg.exitCh <- struct{}{}
}
