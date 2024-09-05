package block_actors

import general "Eulogist/minecraft/protocol/block_actors/general_actors"

// 箱子
type Chest struct {
	general.ChestBlockActor `mapstructure:",squash"`
}

// ID ...
func (c *Chest) ID() string {
	return IDChest
}