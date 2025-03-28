<template>
  <div class="h-full flex flex-col rounded-sm shadow">
    <div class="h-[40px] p-2 flex justify-between items-center bg-white">
      <div class="flex items-end px-3">
        <h2 class="text-xl font-bold mr-2">{{ props.title }}</h2>
        <p class="text-gray-300">{{ props.desc }}</p>
      </div>
      

      <div v-if="actionsSlot" style="margin-right: 5px;">
        <slot name="actions"></slot>
      </div>
      <slot v-else name="actions"></slot>
    </div>
    <div class=" bg-white" :style="{ height: 'calc(100% - 46px)', padding: props.padding, borderTop: '1px solid #e9e9e9' }">
      <slot name="content"></slot>
    </div>
  </div>
</template>
<script setup>
import { useSlots } from 'vue';
const props = defineProps({
  title: {
    type: String,
    default: "测试标题",
  },
  // 描述
  desc: {
    type: String,
    default: "测试描述",
  },
  // padding
  padding: {
    type: String,
    default: "20px",
  },
});

const actionsSlot = useSlots().actions;
</script>