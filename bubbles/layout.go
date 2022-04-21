package bubbles

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/term"
)

const (
	width = 96

	columnWidth = 30
)

var (

	// General.

	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	url = lipgloss.NewStyle().Foreground(special).Render

	// Title.

	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).
			SetString("Lip Gloss")

	descStyle = lipgloss.NewStyle().MarginTop(1)

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)

	// Dialog.

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginRight(2).
				Underline(true)

	// List.

	list = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(subtle).
		MarginRight(2).
		Height(8).
		Width(columnWidth + 1)

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(subtle).
			MarginRight(2).
			Render

	listItem = lipgloss.NewStyle().PaddingLeft(2).Render

	checkMark = lipgloss.NewStyle().SetString("✓").
			Foreground(special).
			PaddingRight(1).
			String()

	listDone = func(s string) string {
		return checkMark + lipgloss.NewStyle().
			Strikethrough(true).
			Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
			Render(s)
	}

	// Paragraphs/History.

	viewpointStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(highlight).
			Margin(1, 3, 0, 0).
			Padding(1, 2).
			Height(19).
			Width(columnWidth).
			Inline(false)

	// Status Bar.

	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.Copy().
			Background(lipgloss.Color("#A550DF")).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

	fishCakeStyle = statusNugget.Copy().Background(lipgloss.Color("#6124DF"))

	// Page.

	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)

func renderViewport() viewport.Model {
	physicalWidth, physicalHeight, _ := term.GetSize(int(os.Stdout.Fd()))

	output := viewport.New(physicalWidth-columnWidth-10, physicalHeight)
	output.Style = viewpointStyle
	str := "I'm baby retro chia 3 wolf moon before they sold out wayfarers cliche. Humblebrag knausgaard tilde bushwick thundercats edison bulb. Intelligentsia gluten-free banjo green juice. Whatever woke tilde tbh neutra quinoa locavore vaporware sartorial taxidermy semiotics heirloom. Man braid pitchfork hot chicken fanny pack taiyaki. Cronut tumblr pok pok franzen copper mug bitters slow-carb roof party PBR&B. Food truck single-origin coffee hammock sustainable lo-fi scenester +1 meggings VHS. Gluten-free forage kombucha, thundercats readymade cloud bread beard gochujang taxidermy cronut blue bottle. Mixtape hashtag ugh fanny pack pickled, iPhone jean shorts banh mi tbh farm-to-table art party swag vape letterpress man braid. Synth microdosing flexitarian keytar. Salvia sriracha kogi kale chips copper mug raclette post-ironic hoodie helvetica typewriter lyft wolf. Leggings poke sartorial gochujang disrupt. Bushwick aesthetic man braid brunch, bitters fanny pack vaporware. Bushwick VHS intelligentsia artisan. Kitsch edison bulb art party vinyl umami vexillologist actually scenester tumblr. Chambray tilde pug sriracha snackwave keytar blue bottle offal gastropub brunch. Next level chillwave irony before they sold out biodiesel shoreditch adaptogen church-key brunch cliche intelligentsia air plant bespoke lumbersexual thundercats. Unicorn truffaut viral migas man braid gastropub occupy neutra dreamcatcher mumblecore tumeric normcore, cliche vinyl. Literally pitchfork blue bottle cardigan chartreuse snackwave occupy. Gentrify seitan polaroid squid tbh deep v scenester banjo fanny pack unicorn. Kickstarter quinoa before they sold out hella everyday carry franzen echo park. Forage pitchfork yuccie yr craft beer YOLO man braid plaid kinfolk locavore try-hard direct trade raw denim XOXO wayfarers. Whatever selvage fanny pack, irony quinoa meh post-ironic portland ethical kitsch godard flexitarian sriracha salvia. Offal typewriter shoreditch live-edge selvage stumptown cold-pressed. Shaman VHS flexitarian venmo hashtag, raclette kale chips gentrify slow-carb trust fund jianbing meditation four dollar toast ennui cray. Tacos fanny pack kale chips kickstarter umami. Banh mi vegan neutra truffaut gluten-free ennui. Four dollar toast vegan kickstarter synth godard thundercats wayfarers gentrify woke fixie mustache hashtag slow-carb tumeric. Flannel biodiesel art party, raclette fashion axe sriracha microdosing austin ugh green juice tote bag hell of skateboard kickstarter you probably haven't heard of them. Man bun helvetica hot chicken coloring book wayfarers polaroid. You probably haven't heard of them green juice aesthetic tattooed flexitarian street art."
	wrapAt := output.Width - output.Style.GetVerticalPadding()
	str = wordwrap.String(str, wrapAt)
	output.SetContent(str)
	return output
}

func renderConfigTab() string {
	return lipgloss.JoinVertical(lipgloss.Top,
		list.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("Input Mode"),
				listDone("fspec"),
				listItem("llvm ir"),
				listItem("smt2"),
			),
		),
		list.Copy().Width(columnWidth).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("Output Mode"),
				listItem("ast"),
				listItem("llvm ir"),
				listDone("smt2"),
			),
		),
	)
}

func renderHeader() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render("Fault"),
		descStyle.Render("A Model Checker for System Dynamic Models"),
		infoStyle.Render("Created By Marianne Bellotti"+divider+url("https://github.com/Fault-lang/Fault")),
	)

}

func (m model) renderLayout() string {
	doc := strings.Builder{}

	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	m.viewport = renderViewport()
	doc.WriteString(lipgloss.JoinVertical(lipgloss.Top, renderHeader(), lipgloss.JoinHorizontal(lipgloss.Top, renderConfigTab(), m.viewport.View())))
	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

	// Okay, let's print it
	return fmt.Sprint(docStyle.Render(doc.String()))
}
