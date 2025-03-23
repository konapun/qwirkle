package game

type Input struct {
  layers map[string]Layer
}

func (i *Input) Read() (string, error) {
  return "", nil
}

type Layer struct {

}
