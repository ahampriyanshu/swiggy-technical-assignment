package game

type Player struct {
	id    int
	cards []Card
}

func (p *Player) AddCard(card Card) {
	p.cards = append(p.cards, card)
}

func (p *Player) RemoveCard(card Card) {
	for i, c := range p.cards {
		if c == card {
			p.cards = append(p.cards[:i], p.cards[i+1:]...)
			break
		}
	}
}

func (p *Player) GiveID() int {
	return p.id
}

func (p *Player) GiveCards() []Card {
	return p.cards
}
