package main

import "fmt"

type Builder interface {
	makeHeader(header string)
	makeContent(header string)
	makeFooter(header string)
	getResult() *Letter
}

type Director struct {
	builder Builder
}

func (d *Director) Construct(header string, content string, footer string) *Letter {
	d.builder.makeHeader(header)
	d.builder.makeContent(content)
	d.builder.makeFooter(footer)
	return d.builder.getResult()
}

type MemoBuilder struct {
	letter *Letter
}

func NewMemoBuilder() *MemoBuilder {
	return &MemoBuilder{new(Letter)}
}

type Letter struct {
	Header  string
	Content string
	Footer  string
}

func (mb *MemoBuilder) makeHeader(header string) {
	mb.letter.Header = header
}

func (mb *MemoBuilder) makeContent(content string) {
	mb.letter.Content = content
}

func (mb *MemoBuilder) makeFooter(footer string) {
	mb.letter.Footer = footer
}

func (mb MemoBuilder) getResult() *Letter {
	return mb.letter
}

func main() {
	header := "Уважаемый Иван Иванович!\n"
	content := "Рассмотрев Ваше предложение о поставке сообщаю: условия неприемлемы. Прошу скидку 5%!\n"
	footer := "С уважением, Василий Васильевич"
	d := Director{NewMemoBuilder()}
	letter := d.Construct(header, content, footer)
	fmt.Println(letter)
}
