import { defineStore } from 'pinia'

export type Card = {
    id: number,
    name: string,
    type: string,
    range: string,
    income: number,
    description: string,
    built: boolean,
    cost: number,
}

function convertBuild(item: any,main: boolean):Card{


    let range = ''
    if(item.rules.value){
        range = item.rules.value
    }else{
        range = item.rules.value_range[0] + ' - ' + item.rules.value_range[1]
    }

    let builded = false

    if(main){
        builded = item.builded
    }

    return {
        cost: item.cost,
        type: item.color,
        name: item.name,
        income: item.rules.ValueAdd,
        range:range,
        built:builded
    } as Card
}

 export const headerAuth = ()=>{

    return {
        'Content-Type': 'application/json',
            'auth': localStorage.getItem('token') || ''
    }
}
export var wsAddr = ""
export var domain = "/api"
export var addr = ""+domain
if(import.meta.env.DEV){
       wsAddr = "wss://game.katelinlis.com"
       domain = "game.katelinlis.com/api"
       addr = "https://"+domain
}





export const useGameStore = defineStore('game', {
    state: () => ({
        modalVisibleDice: false,
        diceResult: 0,
        dice2Result: 0,
        playerName: 'Катя',
        coins: 5,
        currentPlayer: {
            name: '',
            id: 0,
        },
        hasRolled: false,
        cards: [] as Card[],
        shop: [] as Card[],
        mainBuild: [] as Card[],
        selectedCard: null
    }),
    actions: {

        logout(){
            localStorage.removeItem('token')
        },
        fillCards(res){
            this.cards = []
            res.builds.forEach((item:any)=>{
                this.cards.push(convertBuild(item,false))
            })
            this.mainBuild = []
            res.main_builds.forEach((item:any)=>{
                this.mainBuild.push(convertBuild(item,true))
            })

        },
        initGame(){

            let data = localStorage.getItem('token')
            if(!data){
                fetch(`${addr}/connect`,{method: 'POST', body: JSON.stringify({name: this.playerName})}).then(res => res.json()).then(res => {
                    this.coins = res.bank

                    this.fillCards(res)

                    localStorage.setItem('token', res.id)
                    //this.initWS(res.id)
                })
            }else{
                //this.initWS(data)
                this.getPlayer()
            }



        },
        getPlayer(){
            let lobbyID = localStorage.getItem("game")
            fetch(`${addr}/getUser/${lobbyID}`,{
                method: "GET",
                headers: headerAuth()
            }).then((data)=>data.json()).then((res)=>{
                this.coins = res.bank
                this.fillCards(res)
            })
            this.deck()
        },
        deck(){
            let lobbyID = localStorage.getItem("game")
            fetch(`${addr}/deck/${lobbyID}`,{
                method: 'GET',
                headers: headerAuth(),
            }).then(data=>data.json()).then((res)=>{

                this.shop = []
                if(res)
                res.forEach((item:any)=>{
                    this.shop.push(convertBuild(item,false))
                })
            })
        },

        async dice(two = false){
            let lobbyID = localStorage.getItem("game")
            return fetch(`${addr}/dice/${lobbyID}${two ? '?two':''}`,{
                method: 'POST',
                headers: headerAuth(),
            }).then((data)=>{
                if(data.status!= 200){throw "error"}
                return data.json()
            }).then((res)=>{
                this.getPlayer()
                this.deck()
                this.hasRolled = true
                return res
            })
        },
        async buildMain(name: string) {
            let lobbyID = localStorage.getItem("game")
            return fetch(`${addr}/build/${lobbyID}?item=${name}`,{
                method: 'POST',
                headers: headerAuth(),
            }).then((data)=>{
                if(data.status!= 200){throw "error"}
                return data.json()
            }).then((res)=>{
                this.getPlayer()
                return true
            })
        },
        buyCard(card: Card) {
            let lobbyID = localStorage.getItem("game")
            fetch(`${addr}/buyShop/${lobbyID}?item=${card.name}`,{
                method: 'POST',
                headers: headerAuth(),
            }).then((data)=>data.json()).then((res)=>{
                this.getPlayer()
                this.deck()
            })
        }
    }
})