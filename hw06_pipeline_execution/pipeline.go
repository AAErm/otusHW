package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func worker(in In, done In, out Bi) {
	defer close(out)
	for {
		select {
		case <-done:
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			out <- v
		}
	}
}

func ExecutePipeline(in, done In, stages ...Stage) Out {
	for _, stage := range stages {
		stagesCh := make(Bi)
		go worker(in, done, stagesCh)
		in = stage(stagesCh)
	}

	return in
}
