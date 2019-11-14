package hw3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTop10(t *testing.T) {
	text := ""
	top := make([]string, 0, 10)
	expected := make([]string, 0, 10)
	actual := make([]string, 0, 10)
	message := ""

	text = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla magna orci, auctor ut lacus sed, venenatis commodo neque. Integer justo mi, euismod nec euismod sit amet, ultricies sed nisi. Ut quis risus risus. Donec at ultricies arcu, nec blandit ligula. Aenean ut diam ut quam euismod dignissim cursus eget nibh. Mauris porttitor efficitur maximus. Morbi eget urna eu mi varius pharetra.
Nunc efficitur arcu id dapibus ultricies. Morbi nisi elit, eleifend ut ipsum eget, commodo consequat magna. Pellentesque fringilla eros et mauris imperdiet, ut efficitur sapien cursus. Duis fringilla tellus eu purus feugiat scelerisque. Phasellus laoreet urna a odio viverra, at pretium sapien malesuada. Sed scelerisque eget sapien non luctus. Integer eu luctus est. Proin ut lectus aliquet, facilisis nunc ut, accumsan eros.
Suspendisse viverra eu nibh vitae porttitor. Pellentesque molestie tempus leo, condimentum blandit magna molestie eu. Pellentesque suscipit vehicula fermentum. In laoreet lacus nulla, in volutpat elit ultrices mollis. Sed nisl tellus, lacinia vitae lacus vehicula, sodales convallis lorem. Curabitur accumsan tincidunt justo at mattis. Proin felis nisi, volutpat non ipsum vitae, fringilla euismod leo. Nullam tellus est, gravida a bibendum sit amet, viverra id nunc.
Aliquam ut imperdiet quam. Donec congue consectetur mattis. Aenean elementum elit est, egestas tincidunt erat tempor et. Praesent mollis, odio ut dignissim scelerisque, nibh sapien sagittis eros, sed rhoncus odio justo ac orci. Curabitur fringilla vitae mi ac eleifend. Etiam consequat malesuada facilisis. Aliquam quis nisi hendrerit, efficitur libero vel, dapibus nisl. Maecenas tristique ex magna, sed sagittis ante cursus a.
Proin non auctor mi, vel gravida nisl. Nam vehicula nibh eget suscipit porttitor. Vivamus velit nibh, lobortis egestas consectetur ut, accumsan at arcu. Donec imperdiet nibh in tincidunt dictum. Sed consequat ipsum quis velit cursus iaculis. Sed pulvinar dui quis purus dignissim semper. Maecenas dui nunc, sollicitudin vitae ultrices ac, malesuada vitae dolor. Fusce sit amet condimentum lorem. Maecenas ex neque, rutrum eget hendrerit facilisis, pulvinar in metus. Pellentesque eget quam vel tortor hendrerit porta cursus sed mauris.
`
	top = Top10(text)
	expected = []string{"ut", "sed", "eget", "nibh", "vitae", "cursus", "eu"}
	actual = top[:len(expected)]
	message = "Analyze a string: " + text
	require.Len(t, top, 10, message)
	require.ElementsMatch(t, expected, actual, message)

	text = "we have more than ten different words in this long test text"
	top = Top10(text)
	expected = []string{}
	actual = top[:len(expected)]
	message = "Analyze a string: " + text
	require.Len(t, top, 10, message)
	require.ElementsMatch(t, expected, actual, message)

	text = "we have one repeat in this text: word word"
	top = Top10(text)
	expected = []string{"word"}
	actual = top[:len(expected)]
	message = "Analyze a string: " + text
	require.Len(t, top, 8, message)
	require.ElementsMatch(t, expected, actual, message)
}
