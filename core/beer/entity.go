package beer

//as tags json definem como os dados vão ser transformados em JSON
type Beer struct {
	ID    int64     `json:"id"`
	Name  string    `json:"name"`
	Type  BeerType  `json:"type"`
	Style BeerStyle `json:"style"`
}

/*
o comando usado para criar o banco de dados foi:

CREATE TABLE beer (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name text NOT NULL,
   type integer NOT NULL,
   style integer not null
);

Considerei que a criação do banco estava fora do escopo do aplicativo
sendo criado por um(a) analista/DBA
*/

//https://www.thebeerstore.ca/beer-101/beer-types/
type BeerType int

const (
	TypeAle   = 1
	TypeLager = 2
	TypeMalt  = 3
	TypeStout = 4
)

//desta forma a função String pertence ao tipo e pode ser usada da seguinte forma:
// var x TypeAle
// fmt.Println(x.String())
func (t BeerType) String() string {
	switch t {
	case TypeAle:
		return "Ale"
	case TypeLager:
		return "Lager"
	case TypeMalt:
		return "Malt"
	case TypeStout:
		return "Stout"
	}
	return "Unknown"
}

type BeerStyle int

//usando desta forma o compilador vai automaticamente definir os ids sequencialmente
const (
	StyleAmber = iota + 1
	StyleBlonde
	StyleBrown
	StyleCream
	StyleDark
	StylePale
	StyleStrong
	StyleWheat
	StyleRed
	StyleIPA
	StyleLime
	StylePilsner
	StyleGolden
	StyleFruit
	StyleHoney
)

func (t BeerStyle) String() string {
	switch t {
	case StyleAmber:
		return "Amber"
	case StyleBlonde:
		return "Blonde"
	case StyleBrown:
		return "Brown"
	case StyleCream:
		return "Cream"
	case StyleDark:
		return "Dark"
	case StylePale:
		return "Pale"
	case StyleStrong:
		return "Strong"
	case StyleWheat:
		return "Wheat"
	case StyleRed:
		return "Red"
	case StyleIPA:
		return "India Pale Ale"
	case StyleLime:
		return "Lime"
	case StylePilsner:
		return "Pilsner"
	case StyleGolden:
		return "Golden"
	case StyleFruit:
		return "Fruit"
	case StyleHoney:
		return "Honey"
	}
	return "Unknown"
}
