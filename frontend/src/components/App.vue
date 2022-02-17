<template>
  <div class="app">
    <img class="icon" alt="WebSocket Logo" src="../assets/websocket.svg" />
    <h1>WebSocket Client</h1>
    <transceive-container />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import TransceiveContainer from "./TransceiveContainer.vue";

export default defineComponent({
  components: { TransceiveContainer },
});
</script>

<style scoped lang="scss">
.app {
  max-height: 100vh;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  align-items: center;
  justify-content: center;
  display: flex;
  flex-direction: column;
  color: #2c3e50;
  padding-top: 50px;
  font-size: calc(5px + 2vmin);
  min-width: 400px;
}

.icon {
  width: min(50vw, 30vh);
}
</style>
