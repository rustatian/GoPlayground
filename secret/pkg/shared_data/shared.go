package shared_data

type Shared interface {
	Set()
	Get()
}

type SharedData struct {
}

func NewSharedData() Shared {
	return &SharedData{}
}

func (s *SharedData) Set() {

}

func (s *SharedData) Get() {

}
