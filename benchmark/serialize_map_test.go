package benchmark

import "testing"

func addToMap(numberOfInclutions int) map[int]string {
	var mapInstance map[int]string = make(map[int]string)
	for i := 0; i <= numberOfInclutions; i++ {
		mapInstance[i] = "fabian"
	}

	return mapInstance
}

func BenchmarkAddToMap(benchCase *testing.B) {
	for i := 0; i <= benchCase.N; i++ {
		addToMap(i)
	}
}
