package discord

import (
	"math/rand"
	"time"
)

// ランダムなint型の値を返します
//
// `arg`で上限値を指定できます。
func RandInt(arg ...interface{}) int {
	rand.Seed(time.Now().UnixNano())

	min := 0
	max := 0

	var ok bool
	switch len(arg) {
	case 0:
		return rand.Int()
	case 1:
		min, ok = arg[0].(int)
		if min == 0 {
			panic("最小値は0以上を指定してください")
		}

		if !ok {
			panic("引数をintに変換できません")
		}

		return rand.Intn(min)
	case 2:
		min, ok = arg[0].(int)
		if min == 0 {
			panic("最小値は0以上を指定してください")
		}
		if !ok {
			panic("引数をintに変換できません")
		}

		max, ok = arg[1].(int)
		if !ok {
			panic("引数をintに変換できません")
		}
		if min >= max {
			panic("最大値は最小値より大きい値を指定してください")
		}
		if !ok {
			panic("引数をintに変換できません")
		}
	}

	return rand.Intn(max-min) + min
}
