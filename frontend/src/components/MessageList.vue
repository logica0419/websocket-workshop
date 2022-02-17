<template>
  <div class="container">
    <p v-for="message in messages" :key="message" class="message">
      {{ message }}
    </p>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";

export default defineComponent({
  props: {
    messages: {
      type: Array as PropType<string[]>,
      required: true,
    },
  },
});
</script>

<style scoped lang="scss">
.container {
  text-align: left;
  padding: 0 20px;
  width: min(80vw, 500px);
  height: min(50vh, 300px);
  overflow-y: auto;
  border: 3px solid #000000;
  border-radius: 15px;
}

.message {
  margin: 0.3em 0;
  font-size: calc(7.5px + 2vmin);
}
</style>
