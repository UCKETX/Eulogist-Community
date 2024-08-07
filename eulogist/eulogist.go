package Eulogist

import (
	RaknetConnection "Eulogist/core/raknet"
	Client "Eulogist/proxy/mc_client"
	Server "Eulogist/proxy/mc_server"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/pterm/pterm"
)

// Eulogist 函数是整个“赞颂者”程序的入口点
func Eulogist() error {
	var err error
	var config *EulogistConfig
	var neteaseConfigPath string
	var neteaseSkinFileName string
	var waitGroup sync.WaitGroup
	var client *Client.MinecraftClient
	var server *Server.MinecraftServer
	var ClientWasConnected chan struct{}

	// 读取配置文件
	{
		config, err = ReadEulogistConfig()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
	}

	// 使 赞颂者 连接到 网易租赁服
	{
		pterm.Info.Println("Now we try to communicate with Auth Server.")

		server, err = Server.ConnectToServer(
			Server.BasicConfig{
				ServerCode:     config.RentalServerCode,
				ServerPassword: config.RentalServerPassword,
				Token:          config.FBToken,
				AuthServer:     LookUpAuthServerAddress(config.FBToken),
			},
		)
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}

		pterm.Success.Println("Success to create connection with NetEase Minecraft Bedrock Rental Server, now we try to create handshake with it.")

		err = server.WaitClientHandshakeDown()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}

		pterm.Success.Println("Success to create handshake with NetEase Minecraft Bedrock Rental Server.")
	}

	// 根据配置文件的启动类型决定启动方式
	if config.LaunchType == LaunchTypeNormal {
		// 检查 Minecraft 客户端是否存在
		if !FileExist(config.NEMCPath) {
			return fmt.Errorf("Eulogist: Client not found, maybe you did not download or the the path is incorrect")
		}
		// 生成皮肤文件
		if playerSkin := server.GetPlayerSkin(); !FileExist(config.SkinPath) && playerSkin != nil {
			if RaknetConnection.IsZIPFile(playerSkin.FullSkinData) {
				neteaseSkinFileName = "skin.zip"
			} else {
				neteaseSkinFileName = "skin.png"
			}
			err = os.WriteFile(neteaseSkinFileName, playerSkin.FullSkinData, 0600)
			if err != nil {
				return fmt.Errorf("Eulogist: %v", err)
			}
			currentPath, _ := os.Getwd()
			config.SkinPath = fmt.Sprintf(`%s\%s`, currentPath, neteaseSkinFileName)
		}
		// 启动 Eulogist 服务器
		client, ClientWasConnected, err = Client.RunServer()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		client.SetPlayerSkin(server.GetPlayerSkin())
		// 生成网易配置文件
		neteaseConfigPath, err = GenerateNetEaseConfig(config, client.GetServerIP(), client.GetServerPort())
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		// 启动 Minecraft 客户端
		command := exec.Command(config.NEMCPath, fmt.Sprintf("config=%s", neteaseConfigPath))
		go command.Run()
		// 打印准备完成的信息
		pterm.Success.Println("Eulogist is ready! Now we are going to start Minecraft Client.\nThen, the Minecraft Client will connect to Eulogist automatically.")
	} else {
		// 启动 Eulogist 服务器
		client, ClientWasConnected, err = Client.RunServer()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		client.SetPlayerSkin(server.GetPlayerSkin())
		// 打印 赞颂者 准备完成的信息
		pterm.Success.Printf(
			"Eulogist is ready! Please connect to Eulogist manually.\nEulogist server address: %s:%d\n",
			client.GetServerIP(), client.GetServerPort(),
		)
	}

	// 等待 Minecraft 客户端与 赞颂者 完成基本数据包交换
	{
		// 等待 Minecraft 客户端连接
		if config.LaunchType == LaunchTypeNormal {
			timer := time.NewTimer(time.Second * 120)
			defer timer.Stop()
			select {
			case <-timer.C:
				return fmt.Errorf("Eulogist: Failed to create connection with Minecraft")
			case <-ClientWasConnected:
				close(ClientWasConnected)
			}
		} else {
			<-ClientWasConnected
			close(ClientWasConnected)
		}
		pterm.Success.Println("Success to create connection with Minecraft Client, now we try to create handshake with it.")
		// 等待 Minecraft 客户端完成握手
		err = client.WaitClientHandshakeDown()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		pterm.Success.Println("Success to create handshake with Minecraft Client, and then you will login to NetEase Minecraft Bedrock Rental Server.")
	}

	// 设置等待队列
	waitGroup.Add(2)

	// 处理网易租赁服到赞颂者的数据包
	go func() {
		defer func() {
			server.CloseConnection()
			client.CloseConnection()
			waitGroup.Add(-1)
		}()
		for {
			// 初始化一个函数，
			// 用于同步数据到 Minecraft 客户端
			syncFunc := func() error {
				if shieldID := server.GetShieldID(); shieldID != 0 {
					client.SetShieldID(shieldID)
				}
				if entityUniqueID := server.GetEntityUniqueID(); entityUniqueID != 0 {
					client.SetEntityUniqueID(entityUniqueID)
				}
				return nil
			}
			// 读取、过滤数据包，
			// 然后抄送其到 Minecraft 客户端
			errList := server.FiltePacketsAndSendCopy(server.ReadPackets(), client.WritePackets, syncFunc)
			for _, err = range errList {
				if err != nil {
					pterm.Warning.Printf("Eulogist: %v\n", err)
				}
			}
			// 检查连接状态
			select {
			case <-server.GetContext().Done():
				return
			case <-client.GetContext().Done():
				return
			default:
			}
		}
	}()

	// 处理 Minecraft 客户端到赞颂者的数据包
	go func() {
		defer func() {
			client.CloseConnection()
			server.CloseConnection()
			waitGroup.Add(-1)
		}()
		for {
			server.WritePackets(client.ReadPackets())
			select {
			case <-client.GetContext().Done():
				return
			case <-server.GetContext().Done():
				return
			default:
			}
		}
	}()

	// 等待所有 goroutine 完成
	waitGroup.Wait()
	pterm.Info.Println("Server Down. Now all connection was closed.")
	return nil
}
