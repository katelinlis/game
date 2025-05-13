import { defineStore } from 'pinia'
import {addr, domain, headerAuth, useGameStore, wsAddr} from "~/store/game";
import {io, Socket} from "socket.io-client";
import type {DefaultEventsMap} from "@socket.io/component-emitter";

type Player = {
    id: number
    name: string
    ready: boolean
}


var socket : Socket<DefaultEventsMap, DefaultEventsMap>

export const useLobbyStore = defineStore('lobby', {
    state: () => ({
        userID: "",
        scriptModal: false,
        name: '',
        lobbyList: [] ,
        lobby: {
            name: '',
            playerList: [] as Player[]
        }
    }),
    actions: {

        initWSListen(token: string,game: string){

            const gameStore = useGameStore()


            socket = io(`${wsAddr}`, {
                transports: ["websocket"],
                autoConnect: true,                 // Включение авто-подключения
                reconnection: true,                // Включение авто-реконнекта
                reconnectionAttempts: Infinity,    // Неограниченное число попыток
                reconnectionDelay: 2000,           // Задержка между попытками подключения (2 секунды)
                timeout: 20000,                    // Таймаут на подключение (20 секунд)

            })

            socket.on("reconnect", () => {
                // Получаем ID пользователя и ID лобби из локального хранилища
                const token = localStorage.getItem("token");
                const lobbyID = localStorage.getItem("game");

                if (token && lobbyID) {
                    console.log("Rejoining lobby after reconnect...");
                    socket.emit("join-lobby", lobbyID, token);
                }
            });



            socket.emit("join-lobby",game,token)

            socket.on("lobby-state", (playerList) => {
                console.log("Lobby state updated:", playerList)
                this.lobby.playerList = playerList
            })


            // Прием уведомлений
            socket.on("notification", (notif) => {
               console.log(notif.Message)
                if(notif.Message == "Start game"){
                    navigateTo("/game")
                }else if(notif.Message=='not ready') {
                    console.log( this.lobby.playerList)
                    this.lobby.playerList.forEach((item: any) => {
                        if (item.id == notif.Who) {
                            item.ready = false
                        }
                    })
                }else if (notif.Message=='dice') {
                    gameStore.hasRolled = false
                    gameStore.diceResult = notif.Number
                    gameStore.dice2Result = notif.Number2
                    gameStore.modalVisibleDice = true
                    let addTime = 0
                    if(notif.Number2 != 0){
                        addTime = 1000
                    }
                    setTimeout(()=>{
                        gameStore.modalVisibleDice = false
                        gameStore.getPlayer()
                    },2000+addTime)

                    //TODO TOAST
                }else if (notif.Message=='Turn') {
                    gameStore.currentPlayer.name = notif.Name
                }else if(notif.Message=='newBalance'){
                    gameStore.coins = notif.Number
                }else if(notif.Message=='ready'){
                    this.lobby.playerList.forEach((item:any)=>{
                        if(item.id == notif.Who){
                            item.ready = true
                        }
                    })
                }else if(notif.Message == 'connectUser'){
                    if(-1==this.lobby.playerList.findIndex((player)=>{
                       return player.id == notif.Who
                    })){
                        this.lobby.playerList.push({
                            id: notif.Who,
                            name: notif.Name,
                            ready: false
                        })
                    }

                }else if(notif.Message == 'disconnectUser'){
                   this.lobby.playerList =    this.lobby.playerList.filter((player)=>{
                       if(player.id==notif.Who){
                           return player
                       }
                   })
                }


            })

            setInterval(() => {
                socket.emit("ping");
            }, 15000);


            // Обработка отключений
            socket.on("disconnect", () => {
                console.log("Disconnected from server")
                navigateTo("/lobby")
            })




        },
        logout(){
            localStorage.removeItem('token')
        },


         getLobby(){
             fetch(`${addr}/lobby`,{
                method: "GET",
                headers: headerAuth()
            }).then((data)=>data.json()).then((res)=>{
                console.log(res)
                this.lobbyList = res
            })

        },

        ready(){
            let lobbyID = localStorage.getItem("game")
            fetch(`${addr}/lobby/${lobbyID}/ready`,{
                method: "POST",
                headers: headerAuth()
            }).then((data)=>data.json()).then((res)=>{
                console.log(res)


            })
        },

        initGame(){
            let lobbyID = localStorage.getItem("game")
            fetch(`${addr}/init/${lobbyID}`,{
                method: "POST",
                headers: headerAuth()
            }).then((data)=>data.json()).then((res)=>{
                console.log(res)


            })
        },

         async disconnect(){
            if(socket)
            socket.disconnect()
             navigateTo('/lobby')
        },

        connectThisLobby(){

            let lobbyID = localStorage.getItem("game")
            let token = localStorage.getItem("token")
            if(token && lobbyID)
          this.initWSListen(token,lobbyID)
        },

        connectLobby(lobbyID: string){
            fetch(`${addr}/lobby/${lobbyID}`,{
                method: "PUT",
                body: JSON.stringify({username: this.name}),
                headers: headerAuth()
            }).then((data)=>{
                if(data.status!=200){
                    throw "err"
                }
                return data.json()
            }).then((res)=>{
                this.lobby = res
                localStorage.setItem("game",res.id)

                this.connectThisLobby()
                navigateTo('/lobbyroom')
            })

        },

        getLobbyThis(){
            let lobbyID = localStorage.getItem("game")
            fetch(`${addr}/lobby/${lobbyID}`,{
                method: "GET",

                headers: headerAuth()
            }).then((data)=>data.json()).then((res)=>{
                this.lobby = res
                //this.connectThisLobby()

            })
        },

        async createLobby(title: string){
            fetch(`${addr}/lobby`,{
                method: "POST",
                body: JSON.stringify({title: title,username: this.name}),
                headers: headerAuth()
            }).then((data)=>{

               return data.json()
            }).then((res)=>{

                this.lobby = res
                localStorage.setItem("game",res.id)
                this.connectThisLobby()
                navigateTo('/lobbyroom')
            })

        }

    }
})