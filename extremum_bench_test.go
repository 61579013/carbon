package carbon

import "testing"

func BenchmarkCarbon_Closest(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Closest(Yesterday(), Yesterday().AddDay())
	}
	for n := 0; n < b.N; n++ {
		now.Closest(Yesterday().AddDay(), Yesterday())
	}
	for n := 0; n < b.N; n++ {
		now.Closest(Parse("xxx"), Parse("2022-08-05"))
	}
	for n := 0; n < b.N; n++ {
		now.Closest(Parse("2020-08-05"), Parse("xxx"))
	}
}

func BenchmarkCarbon_Farthest(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Farthest(now.AddYear(), Yesterday())
	}
	for n := 0; n < b.N; n++ {
		now.Farthest(Yesterday(), now.AddYear())
	}
	for n := 0; n < b.N; n++ {
		now.Farthest(Parse("xxx"), Parse("2022-08-05"))
	}
	for n := 0; n < b.N; n++ {
		now.Farthest(Parse("2020-08-05"), Parse("xxx"))
	}
}
