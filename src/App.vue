<script setup lang="ts">
import LeaderBoard from './components/LeaderBoard.vue';
import SnakeGame from './components/SnakeGame.vue';
import Welcome from './components/Welcome.vue';

import {Client, Session} from "@heroiclabs/nakama-js"

var client = new Client("defaultkey", "127.0.0.1", "7350");

function getDeviceId(): string {
  var deviceId = localStorage.getItem("deviceId");
  if (!deviceId) {
    deviceId = crypto.randomUUID();
    localStorage.setItem("deviceId", deviceId);
  }
  return deviceId;
}

const vars = {
  'score': '100'
}

async function createSession() {
  const session = await client.authenticateDevice(getDeviceId(), true, undefined, vars);
  return session;
}

</script>

<template>
  <!-- <div class="flex flex-row justify-evenly items-center h-screen"> -->
  <div class="grid grid-cols-3 h-screen w-screen items-center bg-heroic-light text-white">
    <Welcome msg="Snake!" />
    <SnakeGame></SnakeGame>
    <LeaderBoard></LeaderBoard>
  </div>
</template>
