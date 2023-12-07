package mapsutils

func Map2[K, L comparable, V, S interface{}](data map[K]V, newKey func(k K) L, convert func(v V, k K) S) map[L]S {
	result := make(map[L]S)
	for k, v := range data {
		result[newKey(k)] = convert(v, k)
	}
	return result
}

func Map[K comparable, V, S interface{}](data map[K]V, convert func(v V, k K) S) map[K]S {
	result := make(map[K]S)
	for k, v := range data {
		result[k] = convert(v, k)
	}
	return result
}

func ValuesOfMap[K comparable, V interface{}](data map[K]V) []V {
	var values []V
	for _, v := range data {
		values = append(values, v)
	}
	return values
}

func KeysOfMap[K comparable, V interface{}](data map[K]V) []K {
	var keys []K
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}

func ArrayToMap[T interface{}](array []T) map[int]T {
	result := make(map[int]T)
	for i, v := range array {
		result[i] = v
	}
	return result
}
