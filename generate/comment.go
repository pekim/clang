package generate

import (
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

func commentText(comment clang.Comment) string {
	var builder strings.Builder
	commentChildrenText(comment, &builder)
	return builder.String()
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
