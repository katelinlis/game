package model

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Lobby struct {
	sync.Mutex
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	PlayerList  PlayerList `json:"playerList"`
	GameStarted bool       `json:"gameStarted"`
	Created     time.Time  `json:"created"`
}

type LobbyPublic struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	PlayerList  PlayerList `json:"playerList"`
	GameStarted bool       `json:"gameStarted"`
	Created     time.Time  `json:"created"`
}

func (l *Lobby) InitGorutineLobby() {

	ticker := time.NewTicker(10 * time.Second) // Проверяем каждые 10 секунд
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case <-ticker.C:
			func() {
				if !l.GameStarted {

					l.Lock()

					now := time.Now()
					// Удаляем игроков, которые неактивны более 30 секунд
					for _, player := range l.PlayerList {
						if now.Sub(player.Online) > 60*time.Second {
							log.Printf("Player %d is inactive for more than 30 seconds and will be removed.", player.ID)
							l.PlayerList = l.PlayerList.Delete(player.ID)
							// Уведомляем клиента о его удалении

							l.BroadcastSendMessage(Nottif{Name: player.Name, Message: "disconnectUser", Who: player.ID})

						}
					}
					l.Unlock()
				} else {
					ticker.Stop()
					return
				}
			}()
		}
	}

}

type LobbyList struct {
	sync.RWMutex
	List []*Lobby
}

func (l *LobbyList) DeletesLobby() {
	ticker := time.NewTicker(1 * time.Minute) // Проверяем каждые 10 секунд
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			func() {

				for i := 0; i < len(l.List); i++ {
					item := l.List[i]
					if !item.GameStarted && item.Created.After(time.Now().Add(10*time.Minute)) {
						l.List = append(l.List[:i], item)
					}
				}
			}()
		}
	}
}

func (l *LobbyList) AddNewLobby(lobby *Lobby) {
	l.Lock()
	defer l.Unlock()
	(*l).List = append((*l).List, lobby)
	lobby.Created = time.Now()
	go lobby.InitGorutineLobby()
}
func (l *LobbyList) GetLobby(id uuid.UUID) *Lobby {
	l.RLock()
	defer l.RUnlock()
	for i := 0; i < len((*l).List); i++ {
		if (*l).List[i].ID == id {
			return (*l).List[i]
		}
	}
	return nil
}

func NewLobby(name string, user string, userID int64) *Lobby {
	PlayerList := make([]*PlayerListLobby, 0)

	PlayerList = append(PlayerList, &PlayerListLobby{
		ID:      userID,
		Name:    user,
		QNotiff: make(chan Nottif, 10),
	})

	return &Lobby{
		ID:         uuid.New(),
		Created:    time.Now(),
		Name:       strings.ToLower(name),
		PlayerList: PlayerList,
	}
}

func (lobby *Lobby) AddToLobby(user string, userID int64) {
	defer lobby.Unlock()
	lobby.Lock()
	lobby.PlayerList = append(lobby.PlayerList, &PlayerListLobby{
		ID:      userID,
		Name:    user,
		QNotiff: make(chan Nottif, 10),
	})

}

func (lobby *Lobby) PersanalSendMessage(nottif Nottif, player int64) {
	for _, item := range lobby.PlayerList {
		if item.ID == player {
			select {
			case item.QNotiff <- nottif:
			default:
				log.Printf("QNotiff full or unread for player %s, message dropped", item.ID)
			}
		}

	}
}

func (lobby *Lobby) BroadcastSendMessage(nottif Nottif) {

	for _, item := range lobby.PlayerList {
		select {
		case item.QNotiff <- nottif:
		default:
			log.Printf("QNotiff full or unread for player %s, message dropped", item.ID)
		}
	}

}

type PlayerListLobby struct {
	ID      int64       `json:"id"`
	Name    string      `json:"name"`
	Ready   bool        `json:"ready"`
	QNotiff chan Nottif `json:"-"`
	Online  time.Time   `json:"online"`
	WSid    string      `json:"ws_id"`
}
type PlayerList []*PlayerListLobby

func (l PlayerList) GetPlayer(id int64) *PlayerListLobby {
	for i := 0; i < len(l); i++ {
		if (l)[i].ID == id {
			return l[i]
		}
	}
	return nil
}

func (l PlayerList) Delete(player int64) PlayerList {
	local := l
	for i := 0; i < len(l); i++ {
		if l[i].ID == player {
			local = append(l[:i], l[i+1:]...)
			break
		}
	}

	return local
}
