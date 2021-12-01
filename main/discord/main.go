package discord

import (
	"log"
	"math/rand"
	"runtime"
	"rwby-adventures/config"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	CommandRouter = router{
		Prefix:         "r!",
		ListenerPrefix: "<",
		RateLimit:      2000,
	}

	manager, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(err)
	}
	Session = manager

	// recommended, err := manager.GetRecommendedCount()
	// if err != nil {
	// 	log.Fatal("Failed getting recommended shard count")
	// }
	// //if recommended < 2 {
	// manager.SetNumShards(recommended)
	// //}

	manager.AddHandler(routeMessages)
	manager.AddHandler(routeInteraction)
	manager.AddHandler(routeComponents)
	manager.AddHandler(botReady)

	log.Println("Starting the shard manager")
	err = manager.Open()
	if err != nil {
		log.Fatal("Failed to start: ", err)
	}

	time.Sleep(1 * time.Second)

	DefaultFooter = &discordgo.MessageEmbedFooter{
		IconURL: manager.State.User.AvatarURL("256"),
		Text:    "Made by Yewolf - Support : " + config.SupportLink,
	}
}

func botReady(s *discordgo.Session, evt *discordgo.Ready) {
	s.UpdateGameStatus(0, ">help for help (Shard : "+strconv.Itoa(s.ShardID+1)+")")
}
