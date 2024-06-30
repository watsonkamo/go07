package piscine

import "ft"

func PrintNbr(nb int) {
	if nb == 0 {
		ft.PrintRune('0')
		return
	}
	if nb < 0 {
		ft.PrintRune('-')
		nb = -nb
	}
	if nb < 10 {
		ft.PrintRune(rune(nb) + '0')
		return
	}
	PrintNbr(nb / 10)
	PrintNbr(nb % 10)
}

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
	ft.PrintRune('\n')
}
