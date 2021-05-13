package main

import "testing"

func BenchmarkLogTest(b *testing.B) {
	b.Run("sample log", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sampleLogTest()
		}
	})

}

func BenchmarkNotSampleLogTest (b *testing.B) {
	b.Run("not sample log", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			notSampleLogTest()
		}
	})
}




