<template>
  <div
      
      v-if="game.modalVisibleDice"
      class="fixed inset-0 flex items-center justify-center bg-opacity-50 z-50"
  >


      <div class="diceWrap">
        <div  id="dice-1" class="dice rolling">
          <div class="diceFace front"></div>
          <div class="diceFace up"></div>
          <div class="diceFace left"></div>
          <div class="diceFace right"></div>
          <div class="diceFace bottom"></div>
          <div class="diceFace back"></div>
        </div>
    </div>
  </div>
</template>

<script setup>

import {useGameStore} from "~/store/game.js";

const game = useGameStore()
defineProps({
  diceResult: Number, // Число, пришедшее с бэкенда
  dice2Result: Number
})

watch(()=>game.modalVisibleDice, (newValue, oldValue) => {
  if(newValue){
  setTimeout(()=>{
    setValLogic(game.diceResult)

     if(game.dice2Result){
       setTimeout(()=>{
         setValLogic(game.dice2Result)
       },2000)
     }
 },200)
}

})



const perFace = [
  [-0.1, 0.3, -1],
  [-0.1, 0.6, -0.4],
  [-0.85, -0.42, 0.73],
  [-0.8, 0.3, -0.75],
  [0.3, 0.45, 0.9],
  [-0.16, 0.6, 0.18]
];


const setValLogic = (num)=>{

    let dice = document.getElementById(`dice-1`)
    dice.classList.remove('throw','rolling')

    dice.style.transform = `rotate3d(${perFace[num - 1]}, 180deg)`;
    setInterval(() => {
      dice.classList.add('throw')
    }, 50)


}




</script>

<style scoped>
/* Контейнер для кубика */
.dice-container {
  perspective: 1000px;
  width: 100px;
  height: 100px;
  position: relative;
}

.dice {
  position: absolute;
  width: 100px;
  height: 100px;
  top: calc(50% - 50px);
  left: calc(50% - 50px);
  transform-style: preserve-3d;
  transform: rotate3d(0, 0.9, 0.9, 90deg);
  transition: 0.5s cubic-bezier(0.42, 1.57, 0.62, 0.86);
}
.dice.rolling {
  animation: rotatePerFace 3s cubic-bezier(0.42, 1.57, 0.62, 0.86) infinite;
}
.dice.throw {
  animation: rotateDice 0.7s ease-in reverse, throwDice 1s linear;
}
.dice .diceFace {
  box-sizing: border-box;
  position: absolute;
  width: 100px;
  height: 100px;
  background-color: #f6f3f0;
  border: 2px solid white;
  border-radius: 20px;
  transform-style: preserve-3d;
  transition: 0.5s;
}
.dice .diceFace::before {
  position: absolute;
  content: "";
  width: 100%;
  height: 100%;
  background-color: white;
  border-radius: 20px;
  transform: translateZ(-1px);
}
.dice .diceFace::after {
  position: absolute;
  content: "";
  width: 20px;
  height: 20px;
  top: 50%;
  left: 50%;
  margin: -10px 0 0 -10px;
  background-color: #131210;
  border-radius: 100%;
  transform: translateZ(1px);
}
.dice .front {
  transform: translateZ(50px);
}
.dice .front::after {
  width: 40px;
  height: 40px;
  margin: -20px 0 0 -20px;
  background-color: #f63330;
}
.dice .up {
  transform: rotateX(90deg) translateZ(50px);
}
.dice .up::after {
  margin: -30px 0 0 -30px;
  box-shadow: 40px 40px #131210;
}
.dice .left {
  transform: rotateY(-90deg) translateZ(50px);
}
.dice .left::after {
  margin: -40px 0 0 -40px;
  box-shadow: 30px 30px #131210, 60px 60px #131210;
}
.dice .right {
  transform: rotateY(90deg) translateZ(50px);
}
.dice .right::after {
  margin: -30px 0 0 -30px;
  background-color: #f63330;
  box-shadow: 40px 0px #f63330, 0px 40px #f63330, 40px 40px #f63330;
}
.dice .bottom {
  transform: rotateX(-90deg) translateZ(50px);
}
.dice .bottom::after {
  margin: -36px 0 0 -36px;
  box-shadow: 26px 26px #131210, 52px 52px #131210, 52px 0px #131210, 0px 52px #131210;
}
.dice .back {
  transform: rotateX(180deg) translateZ(50px);
}
.dice .back::after {
  margin: -40px 0 0 -30px;
  box-shadow: 40px 0px #131210, 0px 30px #131210, 40px 30px #131210, 0px 60px #131210, 40px 60px #131210;
}
.dice.red .diceFace {
  box-sizing: border-box;
  position: absolute;
  width: 100px;
  height: 100px;
  background-color: rgba(250, 0, 0, 0.45);
  border: 2px solid rgba(255, 46, 46, 0.45);
  border-radius: 30px;
  transform-style: preserve-3d;
  transition: 0.5s;
}
.dice.red .diceFace::before {
  position: absolute;
  content: "";
  width: 100%;
  height: 100%;
  background-color: rgba(255, 46, 46, 0.45);
  border-radius: 20px;
  transform: translateZ(-1px);
}
.dice.red .diceFace::after {
  position: absolute;
  content: "";
  width: 20px;
  height: 20px;
  top: 50%;
  left: 50%;
  margin: -10px 0 0 -10px;
  background-color: white;
  border-radius: 100%;
  transform: translateZ(1px);
}
.dice.red .front {
  transform: translateZ(50px);
}
.dice.red .front::after {
  width: 40px;
  height: 40px;
  margin: -20px 0 0 -20px;
  background-color: white;
}
.dice.red .up {
  transform: rotateX(90deg) translateZ(50px);
}
.dice.red .up::after {
  margin: -30px 0 0 -30px;
  box-shadow: 40px 40px white;
}
.dice.red .left {
  transform: rotateY(-90deg) translateZ(50px);
}
.dice.red .left::after {
  margin: -40px 0 0 -40px;
  box-shadow: 30px 30px white, 60px 60px white;
}
.dice.red .right {
  transform: rotateY(90deg) translateZ(50px);
}
.dice.red .right::after {
  margin: -30px 0 0 -30px;
  background-color: white;
  box-shadow: 40px 0px white, 0px 40px white, 40px 40px white;
}
.dice.red .bottom {
  transform: rotateX(-90deg) translateZ(50px);
}
.dice.red .bottom::after {
  margin: -36px 0 0 -36px;
  box-shadow: 26px 26px white, 52px 52px white, 52px 0px white, 0px 52px white;
}
.dice.red .back {
  transform: rotateX(180deg) translateZ(50px);
}
.dice.red .back::after {
  margin: -40px 0 0 -30px;
  box-shadow: 40px 0px white, 0px 30px white, 40px 30px white, 0px 60px white, 40px 60px white;
}
.dice.blue .diceFace {
  box-sizing: border-box;
  position: absolute;
  width: 100px;
  height: 100px;
  background-color: rgba(0, 0, 255, 0.45);
  border: 2px solid rgba(51, 51, 255, 0.45);
  border-radius: 20px;
  transform-style: preserve-3d;
  transition: 0.5s;
}
.dice.blue .diceFace::before {
  position: absolute;
  content: "";
  width: 100%;
  height: 100%;
  background-color: rgba(51, 51, 255, 0.45);
  border-radius: 20px;
  transform: translateZ(-1px);
}
.dice.blue .diceFace::after {
  position: absolute;
  content: "";
  width: 20px;
  height: 20px;
  top: 50%;
  left: 50%;
  margin: -10px 0 0 -10px;
  background-color: white;
  border-radius: 100%;
  transform: translateZ(1px);
}
.dice.blue .front {
  transform: translateZ(50px);
}
.dice.blue .front::after {
  width: 40px;
  height: 40px;
  margin: -20px 0 0 -20px;
  background-color: white;
}
.dice.blue .up {
  transform: rotateX(90deg) translateZ(50px);
}
.dice.blue .up::after {
  margin: -30px 0 0 -30px;
  box-shadow: 40px 40px white;
}
.dice.blue .left {
  transform: rotateY(-90deg) translateZ(50px);
}
.dice.blue .left::after {
  margin: -40px 0 0 -40px;
  box-shadow: 30px 30px white, 60px 60px white;
}
.dice.blue .right {
  transform: rotateY(90deg) translateZ(50px);
}
.dice.blue .right::after {
  margin: -30px 0 0 -30px;
  background-color: white;
  box-shadow: 40px 0px white, 0px 40px white, 40px 40px white;
}
.dice.blue .bottom {
  transform: rotateX(-90deg) translateZ(50px);
}
.dice.blue .bottom::after {
  margin: -36px 0 0 -36px;
  box-shadow: 26px 26px white, 52px 52px white, 52px 0px white, 0px 52px white;
}
.dice.blue .back {
  transform: rotateX(180deg) translateZ(50px);
}
.dice.blue .back::after {
  margin: -40px 0 0 -30px;
  box-shadow: 40px 0px white, 0px 30px white, 40px 30px white, 0px 60px white, 40px 60px white;
}
.dice.black .diceFace {
  box-sizing: border-box;
  position: absolute;
  width: 100px;
  height: 100px;
  background-color: #111;
  border: 2px solid #2b2b2b;
  border-radius: 20px;
  transform-style: preserve-3d;
  transition: 0.5s;
}
.dice.black .diceFace::before {
  position: absolute;
  content: "";
  width: 100%;
  height: 100%;
  background-color: #2b2b2b;
  border-radius: 20px;
  transform: translateZ(-1px);
}
.dice.black .diceFace::after {
  position: absolute;
  content: "";
  width: 20px;
  height: 20px;
  top: 50%;
  left: 50%;
  margin: -10px 0 0 -10px;
  background-color: #db0;
  border-radius: 100%;
  transform: translateZ(1px);
}
.dice.black .front {
  transform: translateZ(50px);
}
.dice.black .front::after {
  width: 40px;
  height: 40px;
  margin: -20px 0 0 -20px;
  background-color: #3ef;
}
.dice.black .up {
  transform: rotateX(90deg) translateZ(50px);
}
.dice.black .up::after {
  margin: -30px 0 0 -30px;
  box-shadow: 40px 40px #db0;
}
.dice.black .left {
  transform: rotateY(-90deg) translateZ(50px);
}
.dice.black .left::after {
  margin: -40px 0 0 -40px;
  box-shadow: 30px 30px #db0, 60px 60px #db0;
}
.dice.black .right {
  transform: rotateY(90deg) translateZ(50px);
}
.dice.black .right::after {
  margin: -30px 0 0 -30px;
  background-color: #3ef;
  box-shadow: 40px 0px #3ef, 0px 40px #3ef, 40px 40px #3ef;
}
.dice.black .bottom {
  transform: rotateX(-90deg) translateZ(50px);
}
.dice.black .bottom::after {
  margin: -36px 0 0 -36px;
  box-shadow: 26px 26px #db0, 52px 52px #db0, 52px 0px #db0, 0px 52px #db0;
}
.dice.black .back {
  transform: rotateX(180deg) translateZ(50px);
}
.dice.black .back::after {
  margin: -40px 0 0 -30px;
  box-shadow: 40px 0px #db0, 0px 30px #db0, 40px 30px #db0, 0px 60px #db0, 40px 60px #db0;
}
.dice.pink .diceFace {
  box-sizing: border-box;
  position: absolute;
  width: 100px;
  height: 100px;
  background-color: #f69;
  border: 2px solid #ff99bb;
  border-radius: 40px;
  transform-style: preserve-3d;
  transition: 0.5s;
}
.dice.pink .diceFace::before {
  position: absolute;
  content: "";
  width: 100%;
  height: 100%;
  background-color: #ff99bb;
  border-radius: 20px;
  transform: translateZ(-1px);
}
.dice.pink .diceFace::after {
  position: absolute;
  content: "";
  width: 20px;
  height: 20px;
  top: 50%;
  left: 50%;
  margin: -10px 0 0 -10px;
  background-color: #ffe;
  border-radius: 100%;
  transform: translateZ(1px);
}
.dice.pink .front {
  transform: translateZ(50px);
}
.dice.pink .front::after {
  width: 40px;
  height: 40px;
  margin: -20px 0 0 -20px;
  background-color: #fe9;
}
.dice.pink .up {
  transform: rotateX(90deg) translateZ(50px);
}
.dice.pink .up::after {
  margin: -30px 0 0 -30px;
  box-shadow: 40px 40px #ffe;
}
.dice.pink .left {
  transform: rotateY(-90deg) translateZ(50px);
}
.dice.pink .left::after {
  margin: -40px 0 0 -40px;
  box-shadow: 30px 30px #ffe, 60px 60px #ffe;
}
.dice.pink .right {
  transform: rotateY(90deg) translateZ(50px);
}
.dice.pink .right::after {
  margin: -30px 0 0 -30px;
  background-color: #fe9;
  box-shadow: 40px 0px #fe9, 0px 40px #fe9, 40px 40px #fe9;
}
.dice.pink .bottom {
  transform: rotateX(-90deg) translateZ(50px);
}
.dice.pink .bottom::after {
  margin: -36px 0 0 -36px;
  box-shadow: 26px 26px #ffe, 52px 52px #ffe, 52px 0px #ffe, 0px 52px #ffe;
}
.dice.pink .back {
  transform: rotateX(180deg) translateZ(50px);
}
.dice.pink .back::after {
  margin: -40px 0 0 -30px;
  box-shadow: 40px 0px #ffe, 0px 30px #ffe, 40px 30px #ffe, 0px 60px #ffe, 40px 60px #ffe;
}

.diceWrap {
  position: absolute;
  width: 200px;
  height: 200px;
  top: calc(50% - 100px);
  left: calc(50% - 100px);
}
.diceWrap::before {
  position: absolute;
  content: "";
  width: 70%;
  height: 10%;
  top: 90%;
  left: 15%;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 100%;
  filter: blur(10px);
}

@keyframes rotateDice {
  30% {
    transform: rotate3d(1, 1, 1, 0deg);
  }
  100% {
    transform: rotate3d(1, 1, 1, 720deg);
  }
}
@keyframes rotatePerFace {
  16% {
    transform: rotate3d(-0.1, 0.6, -0.4, 180deg);
  }
  32% {
    transform: rotate3d(-0.85, -0.42, 0.73, 180deg);
  }
  48% {
    transform: rotate3d(-0.8, 0.3, -0.75, 180deg);
  }
  64% {
    transform: rotate3d(0.3, 0.45, 0.9, 180deg);
  }
  80% {
    transform: rotate3d(-0.16, 0.6, 0.18, 180deg);
  }
  100% {
    transform: rotate3d(-0.1, 0.3, -1, 180deg);
  }
}
@keyframes throwDice {
  20% {
    margin-top: -100px;
  }
  40% {
    margin-top: 0px;
  }
  60% {
    margin-top: -30px;
  }
  80% {
    margin-top: 0px;
  }
  85% {
    margin-top: -10px;
  }
  90% {
    margin-top: 0px;
  }
  95% {
    margin-top: -3px;
  }
  100% {
    margin-top: 0px;
  }
}
body {
  background-color: #333;
}

.controller {
  position: absolute;
  width: 200px;
  padding: 20px;
  bottom: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.5);
  color: white;
  text-align: center;
  line-height: 40px;
}
</style>
