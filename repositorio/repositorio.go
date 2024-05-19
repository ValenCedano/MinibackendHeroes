package repositorio

import "github.com/ValenCedano/MinibackendHeroes/models"

type BaseDatos struct {
	Superheroes map[string]models.Superhero
}

func NewBaseDatos() *BaseDatos {
	return &BaseDatos{
		Superheroes: map[string]models.Superhero{
			"Spider-Man": {
				Name: "Spider-Man",
				Powerstats: models.Powerstats{
					Intelligence: 90,
					Strength:     55,
					Speed:        67,
					Durability:   75,
					Power:        74,
					Combat:       85,
				},
				Biography: models.Biography{
					FullName: "Peter Parker",
				},
				Images: models.Images{
					XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg",
					SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg",
					MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg",
					LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg",
				},
			},
			"Iron Man": {
				Name: "Iron Man",
				Powerstats: models.Powerstats{
					Intelligence: 100,
					Strength:     85,
					Speed:        58,
					Durability:   85,
					Power:        100,
					Combat:       64,
				},
				Biography: models.Biography{
					FullName: "Tony Stark",
				},
				Images: models.Images{
					XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg",
					SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg",
					MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg",
					LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg",
				},
			},
			"Black Widow": {
				Name: "Black Widow",
				Powerstats: models.Powerstats{
					Intelligence: 75,
					Strength:     13,
					Speed:        33,
					Durability:   30,
					Power:        36,
					Combat:       100,
				},
				Biography: models.Biography{
					FullName: "Natasha Romanoff",
				},
				Images: models.Images{
					XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg",
					SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg",
					MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg",
					LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg",
				},
			},
			"Thor": {
				Name: "Thor",
				Powerstats: models.Powerstats{
					Intelligence: 69,
					Strength:     100,
					Speed:        83,
					Durability:   100,
					Power:        100,
					Combat:       100,
				},
				Biography: models.Biography{
					FullName: "Thor Odinson",
				},
				Images: models.Images{
					XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg",
					SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg",
					MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg",
					LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg",
				},
			},
			"Wolverine": {
				Name: "Wolverine",
				Powerstats: models.Powerstats{
					Intelligence: 63,
					Strength:     32,
					Speed:        50,
					Durability:   100,
					Power:        89,
					Combat:       100,
				},
				Biography: models.Biography{
					FullName: "John Logan",
				},
				Images: models.Images{
					XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
					SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
					MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
					LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg",
				},
			},
		},
	}
}
func (bd *BaseDatos) GetSuperheroByName(name string) (models.Superhero, bool) {
	hero, found := bd.Superheroes[name]
	return hero, found
}
