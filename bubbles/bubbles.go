package bubbles

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
	"github.com/muesli/reflow/wrap"
)

type model struct {
	spec         string
	cursor       int
	config       map[string]int
	configCursor int
	viewport     viewport.Model
	ready        bool
}

func New() model {
	return model{
		config: make(map[string]int),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "up", "k":
			if m.cursor == 1 {
				var (
					cmd  tea.Cmd
					cmds []tea.Cmd
				)
				m.viewport, cmd = m.viewport.Update(msg)
				cmds = append(cmds, cmd)

				return m, tea.Batch(cmds...)
			}

		case "down", "j":
			if m.cursor == 1 {
				var (
					cmd  tea.Cmd
					cmds []tea.Cmd
				)
				m.viewport, cmd = m.viewport.Update(msg)
				cmds = append(cmds, cmd)

				return m, tea.Batch(cmds...)
			}

		case "enter", " ":
			// _, ok := m.selected[m.cursor]
			// if ok {
			// 	delete(m.selected, m.cursor)
			// } else {
			// 	m.selected[m.cursor] = struct{}{}
			// }
		}
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width - columnWidth
		m.viewport.Height = msg.Height - 10
		str := "I'm baby retro chia 3 wolf moon before they sold out wayfarers cliche. Humblebrag knausgaard tilde bushwick thundercats edison bulb. Intelligentsia gluten-free banjo green juice. Whatever woke tilde tbh neutra quinoa locavore vaporware sartorial taxidermy semiotics heirloom. Man braid pitchfork hot chicken fanny pack taiyaki. Cronut tumblr pok pok franzen copper mug bitters slow-carb roof party PBR&B. Food truck single-origin coffee hammock sustainable lo-fi scenester +1 meggings VHS. Gluten-free forage kombucha, thundercats readymade cloud bread beard gochujang taxidermy cronut blue bottle. Mixtape hashtag ugh fanny pack pickled, iPhone jean shorts banh mi tbh farm-to-table art party swag vape letterpress man braid. Synth microdosing flexitarian keytar. Salvia sriracha kogi kale chips copper mug raclette post-ironic hoodie helvetica typewriter lyft wolf. Leggings poke sartorial gochujang disrupt. Bushwick aesthetic man braid brunch, bitters fanny pack vaporware. Bushwick VHS intelligentsia artisan. Kitsch edison bulb art party vinyl umami vexillologist actually scenester tumblr. Chambray tilde pug sriracha snackwave keytar blue bottle offal gastropub brunch. Next level chillwave irony before they sold out biodiesel shoreditch adaptogen church-key brunch cliche intelligentsia air plant bespoke lumbersexual thundercats. Unicorn truffaut viral migas man braid gastropub occupy neutra dreamcatcher mumblecore tumeric normcore, cliche vinyl. Literally pitchfork blue bottle cardigan chartreuse snackwave occupy. Gentrify seitan polaroid squid tbh deep v scenester banjo fanny pack unicorn. Kickstarter quinoa before they sold out hella everyday carry franzen echo park. Forage pitchfork yuccie yr craft beer YOLO man braid plaid kinfolk locavore try-hard direct trade raw denim XOXO wayfarers. Whatever selvage fanny pack, irony quinoa meh post-ironic portland ethical kitsch godard flexitarian sriracha salvia. Offal typewriter shoreditch live-edge selvage stumptown cold-pressed. Shaman VHS flexitarian venmo hashtag, raclette kale chips gentrify slow-carb trust fund jianbing meditation four dollar toast ennui cray. Tacos fanny pack kale chips kickstarter umami. Banh mi vegan neutra truffaut gluten-free ennui. Four dollar toast vegan kickstarter synth godard thundercats wayfarers gentrify woke fixie mustache hashtag slow-carb tumeric. Flannel biodiesel art party, raclette fashion axe sriracha microdosing austin ugh green juice tote bag hell of skateboard kickstarter you probably haven't heard of them. Man bun helvetica hot chicken coloring book wayfarers polaroid. You probably haven't heard of them green juice aesthetic tattooed flexitarian street art."
		wrapAt := m.viewport.Width - m.viewport.Style.GetVerticalPadding()
		str = wordwrap.String(str, wrapAt)
		str = wrap.String(str, wrapAt) // force-wrap long strings
		m.viewport.SetContent(str)
		m.viewport.Init()
		m.ready = true
	}

	return m, nil
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	} else {
		// The header
		s := m.renderLayout()

		// Iterate over our choices
		// for i, choice := range m.choices {

		// 	// Is the cursor pointing at this choice?
		// 	cursor := " " // no cursor
		// 	if m.cursor == i {
		// 		cursor = ">" // cursor!
		// 	}

		// 	// Is this choice selected?
		// 	checked := " " // not selected
		// 	if _, ok := m.selected[i]; ok {
		// 		checked = "x" // selected!
		// 	}

		// 	// Render the row
		// 	s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		// }

		// The footer
		s += "\nPress q to quit.\n"

		// Send the UI for rendering
		return s
	}
}
