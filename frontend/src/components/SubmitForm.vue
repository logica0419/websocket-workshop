<template>
  <form class="container" @submit="onSubmit">
    <input v-if="withValue" v-model="value" class="text-box" type="text" />
    <input class="button" type="submit" :value="buttonName" />
  </form>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
  props: {
    withValue: {
      type: Boolean,
      default: false,
    },
    buttonName: {
      type: String,
      required: true,
    },
  },
  emits: {
    onSubmit: (value: string) => {
      if (value === undefined) {
        return false;
      } else {
        return true;
      }
    },
  },
  setup(_props, { emit }) {
    const value = ref("");

    const onSubmit = (e: Event) => {
      e.preventDefault();
      emit("onSubmit", value.value);
      value.value = "";
    };

    return { onSubmit, value };
  },
});
</script>

<style scoped lang="scss">
.container {
  margin: 0.3em 0 0;
  display: flex;
}

.text-box {
  font-size: calc(4px + 1.5vmin);
  margin-right: 10px;
}

.button {
  font-size: calc(4px + 1.5vmin);
  border: 0;
  border-radius: 0.3em;
  background-color: #e2e2e2;
  &:hover {
    background-color: #d4d4d4;
  }
}
</style>
