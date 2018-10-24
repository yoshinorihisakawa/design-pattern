package template_method_2

type printer interface {
	open() string
	print() string
	close() string
}

type AbstractDisplay struct {
}

/*
client-specified self patternを採用
子のメソッドを呼ぶ事ができないので、引数に子構造体を受け取り、
親構造体のメソッド内で子構造体を呼ぶようにしている。
 */
func (g *AbstractDisplay) Display(printer printer) string {
	result := printer.open()
	for i := 0; i < 5; i++ {
		result += printer.print()
	}
	result += printer.close()
	return result
}

/*
AbstractDisplayを埋め込む事で、継承関係のような状態を実現している。
 */
type NormalDisplay struct {
	*AbstractDisplay
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
	*AbstractDisplay
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

