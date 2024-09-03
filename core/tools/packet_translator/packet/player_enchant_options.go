package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/packet_translator"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type PlayerEnchantOptions struct{}

// 将 from 转换为 neteaseProtocol.ItemEnchantments
func (pk *PlayerEnchantOptions) ToNetEaseItemEnchantments(
	from standardProtocol.ItemEnchantments,
) neteaseProtocol.ItemEnchantments {
	return neteaseProtocol.ItemEnchantments{
		Slot: from.Slot,
		Enchantments: [3][]neteaseProtocol.EnchantmentInstance{
			packet_translator.ConvertSlice(
				from.Enchantments[0],
				func(from standardProtocol.EnchantmentInstance) neteaseProtocol.EnchantmentInstance {
					return neteaseProtocol.EnchantmentInstance{
						Type:       from.Type,
						Level:      from.Level,
						ModEnchant: "",
					}
				},
			),
			packet_translator.ConvertSlice(
				from.Enchantments[1],
				func(from standardProtocol.EnchantmentInstance) neteaseProtocol.EnchantmentInstance {
					return neteaseProtocol.EnchantmentInstance{
						Type:       from.Type,
						Level:      from.Level,
						ModEnchant: "",
					}
				},
			),
			packet_translator.ConvertSlice(
				from.Enchantments[2],
				func(from standardProtocol.EnchantmentInstance) neteaseProtocol.EnchantmentInstance {
					return neteaseProtocol.EnchantmentInstance{
						Type:       from.Type,
						Level:      from.Level,
						ModEnchant: "",
					}
				},
			),
		},
		Unknown: 0,
	}
}

// 将 from 转换为 standardProtocol.ItemEnchantments
func (pk *PlayerEnchantOptions) ToStandardItemEnchantments(
	from neteaseProtocol.ItemEnchantments,
) standardProtocol.ItemEnchantments {
	return standardProtocol.ItemEnchantments{
		Slot: from.Slot,
		Enchantments: [3][]standardProtocol.EnchantmentInstance{
			packet_translator.ConvertSlice(
				from.Enchantments[0],
				func(from neteaseProtocol.EnchantmentInstance) standardProtocol.EnchantmentInstance {
					return standardProtocol.EnchantmentInstance{
						Type:  from.Type,
						Level: from.Level,
					}
				},
			),
			packet_translator.ConvertSlice(
				from.Enchantments[1],
				func(from neteaseProtocol.EnchantmentInstance) standardProtocol.EnchantmentInstance {
					return standardProtocol.EnchantmentInstance{
						Type:  from.Type,
						Level: from.Level,
					}
				},
			),
			packet_translator.ConvertSlice(
				from.Enchantments[2],
				func(from neteaseProtocol.EnchantmentInstance) standardProtocol.EnchantmentInstance {
					return standardProtocol.EnchantmentInstance{
						Type:  from.Type,
						Level: from.Level,
					}
				},
			),
		},
	}
}

func (pk *PlayerEnchantOptions) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.PlayerEnchantOptions{}
	input := standard.(*standardPacket.PlayerEnchantOptions)

	p.Options = packet_translator.ConvertSlice(
		input.Options,
		func(from standardProtocol.EnchantmentOption) neteaseProtocol.EnchantmentOption {
			return neteaseProtocol.EnchantmentOption{
				Cost:            from.Cost,
				Enchantments:    pk.ToNetEaseItemEnchantments(from.Enchantments),
				Name:            from.Name,
				RecipeNetworkID: from.RecipeNetworkID,
			}
		},
	)

	return &p
}

func (pk *PlayerEnchantOptions) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.PlayerEnchantOptions{}
	input := netease.(*neteasePacket.PlayerEnchantOptions)

	p.Options = packet_translator.ConvertSlice(
		input.Options,
		func(from neteaseProtocol.EnchantmentOption) standardProtocol.EnchantmentOption {
			return standardProtocol.EnchantmentOption{
				Cost:            from.Cost,
				Enchantments:    pk.ToStandardItemEnchantments(from.Enchantments),
				Name:            from.Name,
				RecipeNetworkID: from.RecipeNetworkID,
			}
		},
	)

	return &p
}
