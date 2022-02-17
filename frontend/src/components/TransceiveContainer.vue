<template>
  <div class="container">
    Room:
    <select v-model="selectedRoom" class="menu">
      <option v-for="room in rooms" :key="room" :value="room">
        {{ room }}
      </option>
    </select>
  </div>
  <br />
  <message-list :messages="messages" />
  <submit-form
    :with-value="true"
    :button-name="'メッセージを送信'"
    @on-submit="onMessage" />
  <submit-form
    :with-value="false"
    :button-name="'現在時刻を送信'"
    @on-submit="onTime" />
</template>

<script lang="ts">
import { defineComponent, Ref, ref, watch } from "vue";
import MessageList from "./MessageList.vue";
import SubmitForm from "./SubmitForm.vue";
import { useWebSocket } from "../utils/ws";

export default defineComponent({
  components: { MessageList, SubmitForm },
  setup() {
    const messages: Ref<string[]> = ref([]);

    const selectedRoom = ref("A");
    const rooms = ["A", "B", "C"];

    let RWS = useWebSocket(
      (location.protocol === "https:" ? "wss" : "ws") +
        "://" +
        location.host +
        "/api/ws?room=" +
        selectedRoom.value,
      (e: MessageEvent<string>) => {
        messages.value.push(e.data);
      }
    );
    watch(selectedRoom, (val) => {
      RWS.close();
      RWS = useWebSocket(
        (location.protocol === "https:" ? "wss" : "ws") +
          "://" +
          location.host +
          "/api/ws?room=" +
          val,
        (e: MessageEvent<string>) => {
          messages.value.push(e.data);
        }
      );
    });

    const onMessage = (value: string) => {
      const mes = JSON.stringify({
        method: "message",
        args: { message: value },
      });
      RWS.send(mes);
    };

    const onTime = () => {
      const mes = JSON.stringify({
        method: "time",
      });
      RWS.send(mes);
    };

    return { messages, selectedRoom, rooms, onMessage, onTime };
  },
});
</script>

<style scoped lang="scss">
.container {
  font-size: calc(7.5px + 3vmin);
}

.menu {
  font-size: calc(7.5px + 3vmin);
  border-radius: 10px;
}
</style>
