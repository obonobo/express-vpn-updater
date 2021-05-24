package util

type Maybe struct {
	cached          *Response
	possibleSources [](func() *Response)
}

func MaybeOf(factory func() *Response) *Maybe {
	return &Maybe{
		possibleSources: [](func() *Response){factory},
	}
}

func MaybeJust(response *Response) *Maybe {
	return &Maybe{
		cached: response,
	}
}

func (m *Maybe) Value() *Response {
	if m.cached == nil {
		for _, source := range m.possibleSources {
			if m.cached = source(); m.cached != nil {
				break
			}
		}
	}
	return m.cached
}

func (m *Maybe) OrElse(backup *Response) *Response {
	value := m.Value()
	if value == nil {
		value = backup
	}
	return value
}

func (m *Maybe) OrElseUse(backup *Response) *Maybe {
	if m.Value() == nil {
		m.cached = backup
	}
	return m
}

func (m *Maybe) OrElseGet(factory func() *Response) *Response {
	value := m.Value()
	if value == nil {
		value = factory()
	}
	return value
}

func (m *Maybe) OrElseTry(factory func() *Response) *Maybe {
	m.possibleSources = append(m.possibleSources, factory)
	return m
}
