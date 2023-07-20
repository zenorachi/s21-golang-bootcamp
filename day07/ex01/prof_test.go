package cpu

import (
	mincoins "day07/ex02"
	"fmt"
	"github.com/pkg/profile"
	"log"
	"os"
	"os/exec"
	"testing"
)

const USER = "zatannal"

func BenchmarkMinCoins(b *testing.B) {
	prof := profile.Start(profile.ProfilePath("."))

	var (
		coins []int
		val   int
	)
	coins = append([]int{1}, getRandomData(1000)...)
	val = 123451

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = mincoins.MinCoins(val, coins)
	}
	prof.Stop()

	cmd := exec.Command(fmt.Sprintf("/Users/%s//homebrew/Cellar/go/1.20.5/libexec/pkg/tool/darwin_amd64/pprof", USER), "-top", "-unit", "10", "./cpu.pprof")
	cmd.Stdin = os.Stdin

	file, err := os.Create("top10.txt")
	if err != nil {
		log.Fatalln(err)
	}

	cmd.Stdout = file
	err = cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
