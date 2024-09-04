package packet_translate_pool

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	standardPacket "Eulogist/core/standard/protocol/packet"
)

// ...
var StandardPacketIDToNetEasePacketID = map[uint32]uint32{
	standardPacket.IDLogin:                      neteasePacket.IDLogin,
	standardPacket.IDPlayStatus:                 neteasePacket.IDPlayStatus,
	standardPacket.IDServerToClientHandshake:    neteasePacket.IDServerToClientHandshake,
	standardPacket.IDClientToServerHandshake:    neteasePacket.IDClientToServerHandshake,
	standardPacket.IDDisconnect:                 neteasePacket.IDDisconnect,
	standardPacket.IDResourcePacksInfo:          neteasePacket.IDResourcePacksInfo,
	standardPacket.IDResourcePackStack:          neteasePacket.IDResourcePackStack,
	standardPacket.IDResourcePackClientResponse: neteasePacket.IDResourcePackClientResponse,
	standardPacket.IDText:                       neteasePacket.IDText,
	standardPacket.IDSetTime:                    neteasePacket.IDSetTime,
	standardPacket.IDStartGame:                  neteasePacket.IDStartGame,
	standardPacket.IDAddPlayer:                  neteasePacket.IDAddPlayer,
	standardPacket.IDAddActor:                   neteasePacket.IDAddActor,
	standardPacket.IDRemoveActor:                neteasePacket.IDRemoveActor,
	standardPacket.IDAddItemActor:               neteasePacket.IDAddItemActor,
	// ---
	standardPacket.IDTakeItemActor:     neteasePacket.IDTakeItemActor,
	standardPacket.IDMoveActorAbsolute: neteasePacket.IDMoveActorAbsolute,
	standardPacket.IDMovePlayer:        neteasePacket.IDMovePlayer,
	standardPacket.IDPassengerJump:     neteasePacket.IDPassengerJump,
	standardPacket.IDUpdateBlock:       neteasePacket.IDUpdateBlock,
	standardPacket.IDAddPainting:       neteasePacket.IDAddPainting,
	standardPacket.IDTickSync:          neteasePacket.IDTickSync,
	// ---
	standardPacket.IDLevelEvent:           neteasePacket.IDLevelEvent,
	standardPacket.IDMobEffect:            neteasePacket.IDMobEffect,
	standardPacket.IDUpdateAttributes:     neteasePacket.IDUpdateAttributes,
	standardPacket.IDInventoryTransaction: neteasePacket.IDInventoryTransaction,
	standardPacket.IDMobEquipment:         neteasePacket.IDMobEquipment,
	standardPacket.IDMobArmourEquipment:   neteasePacket.IDMobArmourEquipment,
	standardPacket.IDInteract:             neteasePacket.IDInteract,
	standardPacket.IDBlockPickRequest:     neteasePacket.IDBlockPickRequest,
	standardPacket.IDActorPickRequest:     neteasePacket.IDActorPickRequest,
	standardPacket.IDPlayerAction:         neteasePacket.IDPlayerAction,
	// ---
	standardPacket.IDHurtArmour:                  neteasePacket.IDHurtArmour,
	standardPacket.IDSetActorData:                neteasePacket.IDSetActorData,
	standardPacket.IDSetActorMotion:              neteasePacket.IDSetActorMotion,
	standardPacket.IDSetActorLink:                neteasePacket.IDSetActorLink,
	standardPacket.IDSetHealth:                   neteasePacket.IDSetHealth,
	standardPacket.IDSetSpawnPosition:            neteasePacket.IDSetSpawnPosition,
	standardPacket.IDAnimate:                     neteasePacket.IDAnimate,
	standardPacket.IDRespawn:                     neteasePacket.IDRespawn,
	standardPacket.IDContainerOpen:               neteasePacket.IDContainerOpen,
	standardPacket.IDContainerClose:              neteasePacket.IDContainerClose,
	standardPacket.IDPlayerHotBar:                neteasePacket.IDPlayerHotBar,
	standardPacket.IDInventoryContent:            neteasePacket.IDInventoryContent,
	standardPacket.IDInventorySlot:               neteasePacket.IDInventorySlot,
	standardPacket.IDCraftingData:                neteasePacket.IDCraftingData,
	standardPacket.IDCraftingEvent:               neteasePacket.IDCraftingEvent,
	standardPacket.IDGUIDataPickItem:             neteasePacket.IDGUIDataPickItem,
	standardPacket.IDBlockActorData:              neteasePacket.IDBlockActorData,
	standardPacket.IDPlayerInput:                 neteasePacket.IDPlayerInput,
	standardPacket.IDLevelChunk:                  neteasePacket.IDLevelChunk,
	standardPacket.IDSetCommandsEnabled:          neteasePacket.IDSetCommandsEnabled,
	standardPacket.IDSetDifficulty:               neteasePacket.IDSetDifficulty,
	standardPacket.IDChangeDimension:             neteasePacket.IDChangeDimension,
	standardPacket.IDSetPlayerGameType:           neteasePacket.IDSetPlayerGameType,
	standardPacket.IDPlayerList:                  neteasePacket.IDPlayerList,
	standardPacket.IDSimpleEvent:                 neteasePacket.IDSimpleEvent,
	standardPacket.IDEvent:                       neteasePacket.IDEvent,
	standardPacket.IDClientBoundMapItemData:      neteasePacket.IDClientBoundMapItemData,
	standardPacket.IDMapInfoRequest:              neteasePacket.IDMapInfoRequest,
	standardPacket.IDRequestChunkRadius:          neteasePacket.IDRequestChunkRadius,
	standardPacket.IDChunkRadiusUpdated:          neteasePacket.IDChunkRadiusUpdated,
	standardPacket.IDItemFrameDropItem:           neteasePacket.IDItemFrameDropItem,
	standardPacket.IDGameRulesChanged:            neteasePacket.IDGameRulesChanged,
	standardPacket.IDCamera:                      neteasePacket.IDCamera,
	standardPacket.IDBossEvent:                   neteasePacket.IDBossEvent,
	standardPacket.IDShowCredits:                 neteasePacket.IDShowCredits,
	standardPacket.IDAvailableCommands:           neteasePacket.IDAvailableCommands,
	standardPacket.IDCommandRequest:              neteasePacket.IDCommandRequest,
	standardPacket.IDCommandBlockUpdate:          neteasePacket.IDCommandBlockUpdate,
	standardPacket.IDCommandOutput:               neteasePacket.IDCommandOutput,
	standardPacket.IDUpdateTrade:                 neteasePacket.IDUpdateTrade,
	standardPacket.IDUpdateEquip:                 neteasePacket.IDUpdateEquip,
	standardPacket.IDResourcePackDataInfo:        neteasePacket.IDResourcePackDataInfo,
	standardPacket.IDResourcePackChunkData:       neteasePacket.IDResourcePackChunkData,
	standardPacket.IDResourcePackChunkRequest:    neteasePacket.IDResourcePackChunkRequest,
	standardPacket.IDTransfer:                    neteasePacket.IDTransfer,
	standardPacket.IDPlaySound:                   neteasePacket.IDPlaySound,
	standardPacket.IDStopSound:                   neteasePacket.IDStopSound,
	standardPacket.IDSetTitle:                    neteasePacket.IDSetTitle,
	standardPacket.IDAddBehaviourTree:            neteasePacket.IDAddBehaviourTree,
	standardPacket.IDStructureBlockUpdate:        neteasePacket.IDStructureBlockUpdate,
	standardPacket.IDShowStoreOffer:              neteasePacket.IDShowStoreOffer,
	standardPacket.IDPurchaseReceipt:             neteasePacket.IDPurchaseReceipt,
	standardPacket.IDPlayerSkin:                  neteasePacket.IDPlayerSkin,
	standardPacket.IDSubClientLogin:              neteasePacket.IDSubClientLogin,
	standardPacket.IDAutomationClientConnect:     neteasePacket.IDAutomationClientConnect,
	standardPacket.IDSetLastHurtBy:               neteasePacket.IDSetLastHurtBy,
	standardPacket.IDBookEdit:                    neteasePacket.IDBookEdit,
	standardPacket.IDNPCRequest:                  neteasePacket.IDNPCRequest,
	standardPacket.IDPhotoTransfer:               neteasePacket.IDPhotoTransfer,
	standardPacket.IDModalFormRequest:            neteasePacket.IDModalFormRequest,
	standardPacket.IDModalFormResponse:           neteasePacket.IDModalFormResponse,
	standardPacket.IDServerSettingsRequest:       neteasePacket.IDServerSettingsRequest,
	standardPacket.IDServerSettingsResponse:      neteasePacket.IDServerSettingsResponse,
	standardPacket.IDShowProfile:                 neteasePacket.IDShowProfile,
	standardPacket.IDSetDefaultGameType:          neteasePacket.IDSetDefaultGameType,
	standardPacket.IDRemoveObjective:             neteasePacket.IDRemoveObjective,
	standardPacket.IDSetDisplayObjective:         neteasePacket.IDSetDisplayObjective,
	standardPacket.IDSetScore:                    neteasePacket.IDSetScore,
	standardPacket.IDLabTable:                    neteasePacket.IDLabTable,
	standardPacket.IDUpdateBlockSynced:           neteasePacket.IDUpdateBlockSynced,
	standardPacket.IDMoveActorDelta:              neteasePacket.IDMoveActorDelta,
	standardPacket.IDSetScoreboardIdentity:       neteasePacket.IDSetScoreboardIdentity,
	standardPacket.IDSetLocalPlayerAsInitialised: neteasePacket.IDSetLocalPlayerAsInitialised,
	standardPacket.IDUpdateSoftEnum:              neteasePacket.IDUpdateSoftEnum,
	standardPacket.IDNetworkStackLatency:         neteasePacket.IDNetworkStackLatency,
	// ---
	standardPacket.IDSpawnParticleEffect:         neteasePacket.IDSpawnParticleEffect,
	standardPacket.IDAvailableActorIdentifiers:   neteasePacket.IDAvailableActorIdentifiers,
	standardPacket.IDNetworkChunkPublisherUpdate: neteasePacket.IDNetworkChunkPublisherUpdate,
	standardPacket.IDBiomeDefinitionList:         neteasePacket.IDBiomeDefinitionList,
	standardPacket.IDLevelSoundEvent:             neteasePacket.IDLevelSoundEvent,
	standardPacket.IDLevelEventGeneric:           neteasePacket.IDLevelEventGeneric,
	standardPacket.IDLecternUpdate:               neteasePacket.IDLecternUpdate,
	// ---
	standardPacket.IDAddEntity:                     neteasePacket.IDAddEntity,
	standardPacket.IDRemoveEntity:                  neteasePacket.IDRemoveEntity,
	standardPacket.IDClientCacheStatus:             neteasePacket.IDClientCacheStatus,
	standardPacket.IDOnScreenTextureAnimation:      neteasePacket.IDOnScreenTextureAnimation,
	standardPacket.IDMapCreateLockedCopy:           neteasePacket.IDMapCreateLockedCopy,
	standardPacket.IDStructureTemplateDataRequest:  neteasePacket.IDStructureTemplateDataRequest,
	standardPacket.IDStructureTemplateDataResponse: neteasePacket.IDStructureTemplateDataResponse,
	// ---
	standardPacket.IDClientCacheBlobStatus:   neteasePacket.IDClientCacheBlobStatus,
	standardPacket.IDClientCacheMissResponse: neteasePacket.IDClientCacheMissResponse,
	// standardPacket.IDEducationSettings:                  neteasePacket.IDEducationSettings, // TODO
	standardPacket.IDEmote:                             neteasePacket.IDEmote,
	standardPacket.IDMultiPlayerSettings:               neteasePacket.IDMultiPlayerSettings,
	standardPacket.IDSettingsCommand:                   neteasePacket.IDSettingsCommand,
	standardPacket.IDAnvilDamage:                       neteasePacket.IDAnvilDamage,
	standardPacket.IDCompletedUsingItem:                neteasePacket.IDCompletedUsingItem,
	standardPacket.IDNetworkSettings:                   neteasePacket.IDNetworkSettings,
	standardPacket.IDPlayerAuthInput:                   neteasePacket.IDPlayerAuthInput,
	standardPacket.IDCreativeContent:                   neteasePacket.IDCreativeContent,
	standardPacket.IDPlayerEnchantOptions:              neteasePacket.IDPlayerEnchantOptions,
	standardPacket.IDItemStackRequest:                  neteasePacket.IDItemStackRequest,
	standardPacket.IDItemStackResponse:                 neteasePacket.IDItemStackResponse,
	standardPacket.IDPlayerArmourDamage:                neteasePacket.IDPlayerArmourDamage,
	standardPacket.IDCodeBuilder:                       neteasePacket.IDCodeBuilder,
	standardPacket.IDUpdatePlayerGameType:              neteasePacket.IDUpdatePlayerGameType,
	standardPacket.IDEmoteList:                         neteasePacket.IDEmoteList,
	standardPacket.IDPositionTrackingDBServerBroadcast: neteasePacket.IDPositionTrackingDBServerBroadcast,
	standardPacket.IDPositionTrackingDBClientRequest:   neteasePacket.IDPositionTrackingDBClientRequest,
	standardPacket.IDDebugInfo:                         neteasePacket.IDDebugInfo,
	standardPacket.IDPacketViolationWarning:            neteasePacket.IDPacketViolationWarning,
	standardPacket.IDMotionPredictionHints:             neteasePacket.IDMotionPredictionHints,
	standardPacket.IDAnimateEntity:                     neteasePacket.IDAnimateEntity,
	standardPacket.IDCameraShake:                       neteasePacket.IDCameraShake,
	standardPacket.IDPlayerFog:                         neteasePacket.IDPlayerFog,
	standardPacket.IDCorrectPlayerMovePrediction:       neteasePacket.IDCorrectPlayerMovePrediction,
	standardPacket.IDItemComponent:                     neteasePacket.IDItemComponent,
	standardPacket.IDFilterText:                        neteasePacket.IDFilterText,
	standardPacket.IDClientBoundDebugRenderer:          neteasePacket.IDClientBoundDebugRenderer,
	standardPacket.IDSyncActorProperty:                 neteasePacket.IDSyncActorProperty,
	standardPacket.IDAddVolumeEntity:                   neteasePacket.IDAddVolumeEntity,
	standardPacket.IDRemoveVolumeEntity:                neteasePacket.IDRemoveVolumeEntity,
	standardPacket.IDSimulationType:                    neteasePacket.IDSimulationType,
	standardPacket.IDNPCDialogue:                       neteasePacket.IDNPCDialogue,
	standardPacket.IDEducationResourceURI:              neteasePacket.IDEducationResourceURI,
	standardPacket.IDCreatePhoto:                       neteasePacket.IDCreatePhoto,
	standardPacket.IDUpdateSubChunkBlocks:              neteasePacket.IDUpdateSubChunkBlocks,
	standardPacket.IDSubChunk:                          neteasePacket.IDSubChunk,
	standardPacket.IDSubChunkRequest:                   neteasePacket.IDSubChunkRequest,
	standardPacket.IDClientStartItemCooldown:           neteasePacket.IDClientStartItemCooldown,
	standardPacket.IDScriptMessage:                     neteasePacket.IDScriptMessage,
	standardPacket.IDCodeBuilderSource:                 neteasePacket.IDCodeBuilderSource,
	standardPacket.IDTickingAreasLoadStatus:            neteasePacket.IDTickingAreasLoadStatus,
	standardPacket.IDDimensionData:                     neteasePacket.IDDimensionData,
	// standardPacket.IDAgentAction:                      neteasePacket.IDAgentAction, // TODO
	standardPacket.IDChangeMobProperty: neteasePacket.IDChangeMobProperty,
	// standardPacket.IDLessonProgress:    neteasePacket.IDLessonProgress, // TODO
	standardPacket.IDRequestAbility:                neteasePacket.IDRequestAbility,
	standardPacket.IDRequestPermissions:            neteasePacket.IDRequestPermissions,
	standardPacket.IDToastRequest:                  neteasePacket.IDToastRequest,
	standardPacket.IDUpdateAbilities:               neteasePacket.IDUpdateAbilities,
	standardPacket.IDUpdateAdventureSettings:       neteasePacket.IDUpdateAdventureSettings,
	standardPacket.IDDeathInfo:                     neteasePacket.IDDeathInfo,
	standardPacket.IDEditorNetwork:                 neteasePacket.IDEditorNetwork,
	standardPacket.IDFeatureRegistry:               neteasePacket.IDFeatureRegistry,
	standardPacket.IDServerStats:                   neteasePacket.IDServerStats,
	standardPacket.IDRequestNetworkSettings:        neteasePacket.IDRequestNetworkSettings,
	standardPacket.IDGameTestRequest:               neteasePacket.IDGameTestRequest,
	standardPacket.IDGameTestResults:               neteasePacket.IDGameTestResults,
	standardPacket.IDUpdateClientInputLocks:        neteasePacket.IDUpdateClientInputLocks,
	standardPacket.IDCameraPresets:                 neteasePacket.IDCameraPresets,
	standardPacket.IDUnlockedRecipes:               neteasePacket.IDUnlockedRecipes,
	standardPacket.IDCameraInstruction:             neteasePacket.IDCameraInstruction,
	standardPacket.IDCompressedBiomeDefinitionList: neteasePacket.IDCompressedBiomeDefinitionList,
	standardPacket.IDTrimData:                      neteasePacket.IDTrimData,
	standardPacket.IDOpenSign:                      neteasePacket.IDOpenSign,
	standardPacket.IDAgentAnimation:                neteasePacket.IDAgentAnimation,
}

// ...
var NetEasePacketIDToStandardPacketID map[uint32]uint32
