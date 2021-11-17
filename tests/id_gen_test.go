package tests

import (
	"testing"

	"github.com/galaxy-future/BridgX/pkg/id_generator"
)

func TestIdGenerator(t *testing.T) {
	id := id_generator.GetNextId()
	t.Logf("%v", id)
}

func BenchmarkIdGenerator(b *testing.B) {
	//	cpu: Intel(R) Xeon(R) Platinum 8163 CPU @ 2.50GHz
	//	BenchmarkIdGenerator
	//	BenchmarkIdGenerator-4   	1000000000	         0.0000017 ns/op
	_ = id_generator.GetNextId()
}
