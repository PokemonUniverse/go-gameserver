package pokemon

import (
	"github.com/PokemonUniverse/nonamelib/log"
	"github.com/eaigner/hood"

	"gameserver/data/entities"
)

var (
	pokemonStore map[int64]*Pokemon
	movesStore   map[int64]*Move
)

func init() {
	pokemonStore = make(map[int64]*Pokemon)
	movesStore = make(map[int64]*Move)
}

func Load(_hood *hood.Hood) {
	// Load moves
	loadMoves(_hood)

	// Load Pokemon
	loadPokemon(_hood)
}

func loadMoves(_hood *hood.Hood) {
	log.Verbose("Pokemon.Manager", "loadMoves", "Loading pokemon moves from database")

	var moves []entities.Move
	if err := _hood.Find(&moves); err != nil {
		log.Error("Pokemon.Manager", "loadMoves", "Error while loading moves: %s", err.Error())
	} else {
		for _, moveEntity := range moves {
			move := NewMove(&moveEntity)
			movesStore[move.GetMoveId()] = move
		}
	}
}

func loadPokemon(_hood *hood.Hood) {
	log.Verbose("Pokemon.Manager", "loadPokemon", "Loading pokemon from database")

	var pokemon []entities.Pokemon
	if err := _hood.Find(&pokemon); err != nil {
		log.Error("Pokemon.Manager", "loadPokemon", "Error while loading pokemon: %s", err.Error())
	} else {
		for _, pokemonEntity := range pokemon {
			poke := NewPokemon(&pokemonEntity)

			if err := poke.LinkMoves(_hood); err != nil {
				log.Error("Pokemon.Manager", "loadPokemon", "Failed to link moves to Pokemon %d. Error: %s", poke.GetPokemonId(), err.Error())
			} else {
				pokemonStore[poke.GetPokemonId()] = poke
			}
		}
	}
}

func GetPokemonById(_id int64) (*Pokemon, bool) {
	pokemon, found := pokemonStore[_id]
	return pokemon, found
}

func GetMoveById(_id int64) (*Move, bool) {
	move, found := movesStore[_id]
	return move, found
}
