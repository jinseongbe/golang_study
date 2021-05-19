package dataStruct

func Hash(s string) int {
	h := 0
	a := 256  // A값은 S문자 값의 범위를 사용하는걸 권장(0~255)
	b := 3571 // B값은 소수를 사용하는게 좋다

	for i := 0; i < len(s); i++ {
		h = (h*a + int(s[i])) % b
	}

	return h
}

type keyValue struct {
	key   string
	value string
}

type Map struct {
	keyArray [3571][]keyValue
}

func CreatMap() *Map {
	return &Map{}
}

func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.keyArray[h] = append(m.keyArray[h], keyValue{key, value})
}

func (m *Map) Get(key string) string {
	h := Hash(key)
	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].key == key {
			return m.keyArray[h][i].value
		}
	}

	return ""
}
