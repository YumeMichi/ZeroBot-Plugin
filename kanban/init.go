// Package kanban 打印版本信息
package kanban

import (
	"fmt"

	"github.com/FloatTech/ZeroBot-Plugin/kanban/banner"
)

//go:generate go run github.com/FloatTech/ZeroBot-Plugin/kanban/gen

func init() {
	PrintBanner()
}

// PrintBanner ...
func PrintBanner() {
	fmt.Print(
		"\n======================[ZeroBot-Plugin]======================",
		"\n", banner.Banner, "\n",
		"============================================================\n\n",
	)
}
