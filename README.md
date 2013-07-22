# Pokemon Universe Game Server
Pokemon Universe game server

## Dependencies

- [NoNameLib](github.com/PokemonUniverse/nonamelib) (github.com/PokemonUniverse/nonamelib)
- [Hood ORM](github.com/eaigner/hood) (github.com/eaigner/hood)
- [MyMySQL](github.com/ziutek/mymysql/godrv) (github.com/ziutek/mymysql/godrv)

# How to install
Clone the `gameserver` repository in the root of your `$GOPATH/src` folder

	git clone https://github.com/PokemonUniverse/gameserver $GOPATH/src

When cloning is done, build the gameserver
	
    cd $GOPATH\bin
    go build gameserver

# TODO List
Note: List items are not sorted in any particular order

- Load all Pokemon (including moves, forms and evolutions)
- Items
- NPCs (loading, moving, player interaction)
- Chat
- Battle interface between PU gameserver and Pokemon-Online battle simulator
- Tile scripts. Ability to link a script directly to a TilePointLayer
- Write tests for all packages (I know.. this should be done before writing actual code) 
- Everything else