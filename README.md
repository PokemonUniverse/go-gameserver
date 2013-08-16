# Pokemon Universe Game Server
Pokemon Universe game server

## Dependencies

- [NoNameLib](github.com/PokemonUniverse/nonamelib) (github.com/PokemonUniverse/nonamelib)
- [Go Websockets](https://code.google.com/p/go/source/checkout?repo=net) (https://code.google.com/p/go/source/checkout?repo=net)
- [Hood ORM](github.com/eaigner/hood) (github.com/eaigner/hood)
- [MyMySQL](github.com/ziutek/mymysql/godrv) (github.com/ziutek/mymysql/godrv)

For the lazy:

    go get https://code.google.com/p/go.net/
    go get https://github.com/PokemonUniverse/nonamelib
    go get https://github.com/eaigner/hood
    go get https://github.com/ziutek/mymysql/godrv

# How to install
Clone the `gameserver` repository in the root of your `$GOPATH/src` folder

	git clone https://github.com/PokemonUniverse/gameserver $GOPATH/src

When cloning is done, build the gameserver
	
    cd $GOPATH/bin
    go build gameserver

# TODO List
Note: List items are not sorted in any particular order

- ~~Load Pokemon data (including moves, forms and evolutions)~~
- Items
- NPCs (loading, moving, player interaction)
- Chat
- Battle interface between PU gameserver and Pokemon-Online battle simulator
- Tile scripts. Ability to link a script directly to a TilePointLayer
- Pokemon logic (leveling up, learning moves, etc..)
- Headless client for Pokemon Online Battle Simulator (POBS)
- Interface between POBS and our server (for updating Pokemon info)
- Write tests for all packages (I know.. this should be done before writing actual code) 
- Everything else