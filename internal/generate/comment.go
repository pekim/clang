package generate

import (
	"strings"

	"github.com/pekim/clang"
)

func commentText(comment clang.Comment) string {
	var builder strings.Builder
	commentChildrenText(comment, &builder)

	// trim leading and trailing white space from lines
	lines := strings.Split(builder.String(), "\n")
	builder.Reset()
	for _, line := range lines {
		builder.WriteString(strings.TrimSpace(line))
		builder.WriteRune('\n')
	}

	return strings.TrimSpace(builder.String())
}

func commentChildrenText(comment clang.Comment, builder *strings.Builder) {
	for i := range comment.NumChildren() {
		child := comment.Child(i)
		switch child.Kind() {

		case clang.Comment_Paragraph:
			builder.WriteString(child.TextComment_getText())
			commentChildrenText(child, builder)
			builder.WriteString("\n\n")

		case clang.Comment_Text:
			builder.WriteString(child.TextComment_getText())

		case clang.Comment_VerbatimLine:
			child.TextComment_getText()

		case clang.Comment_InlineCommand:
			for i := range child.InlineCommandComment_getNumArgs() {
				builder.WriteString(child.InlineCommandComment_getArgText(i))
			}
		}
	}
}
