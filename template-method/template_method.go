package template_method

type printer interface {
	open() string
	print() string
	close() string
}


type NormalDisplay struct {
	str string
}

func (n *NormalDisplay) open() string {
	return "======\n"
}
func (n *NormalDisplay) print() string {
	return "|" + string(n.str) + "|\n"
}
func (n *NormalDisplay) close() string {
	return "======"
}

type GorgeousDisplay struct {
	Str string
}

func (g *GorgeousDisplay) open() string {
	return "＊*＊*＊*＊*＊\n"
}
func (g *GorgeousDisplay) print() string {
	return "＊" + g.Str + "＊\n"
}
func (g *GorgeousDisplay) close() string {
	return "＊*＊*＊*＊*＊"
}

/*
interfaceを埋め込む方法
strategyパターンでも同様の方法が採用されている。
 */
type Display struct {
	printer printer
}

func (g *Display) Display() string {
	result := g.printer.open()
	for i := 0; i < 5; i++ {
		result += g.printer.print()
	}
	result += g.printer.close()
	return result
}