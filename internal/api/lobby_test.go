package api

import (
	"testing"

	"github.com/katelinlis/BackendMasters/internal/model"
)

func TestCreateLobby(t *testing.T) {
	lobby := model.LobbyList{}
	lobbyNew := model.NewLobby("title", "username", 1)

	lobby.AddNewLobby(lobbyNew)
	find := lobby.GetLobby(lobbyNew.ID)

	find.AddToLobby("user2", 2)

	t.Log(find.ID)
	t.Log(lobbyNew.ID)
	//assert.Equal(t, lobbyNew, find)
}
