package kata2

type Fixture struct {
    input    string
    expected string
    result   *string
}

func (self *Fixture) Assert() bool {
    self.makeKata()
    return self.Result() == self.expected
}

func (self *Fixture) makeKata() {
    result := MakeKata(self.input)
    self.result = &result
}

func (self *Fixture) Result() string {
    return *self.result
}
