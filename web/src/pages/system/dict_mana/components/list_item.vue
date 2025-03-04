<template>
  <div
    class="w-full flex items-center mb-3 transform transition-all duration-300 ease-in-out hover:translate-x-1"
    @click="handleCellClick"
  >
    <div
      class="h-12 w-1 rounded-l transition-all duration-300"
      :class="{ 
        'bg-primary-500 opacity-100': isShowActive, 
        'bg-transparent opacity-0': !isShowActive 
      }"
    ></div>
    <div
      class="w-full flex border border-solid border-gray-100 justify-between px-4 h-12 items-center rounded-r-lg transition-all duration-300"
      :class="{
        'bg-primary-50 shadow-md': isShowActive,
        'hover:bg-gray-50 shadow-sm': !isShowActive
      }"
    > 
      <div class="flex items-center space-x-3">
        <el-tag size="small" effect="plain" class="min-w-[80px] text-center">{{ data.code }}</el-tag>
        <span class="font-medium text-gray-700">{{ data.name }}</span>
      </div>
      <div class="flex items-center space-x-2">
        <el-button
          class="!px-3"
          :icon="Edit"
          link
          @click="handleEditClick"
          type="primary"
        >修改</el-button>
        <el-divider direction="vertical" class="!mx-1" />
        <el-button
          class="!px-3"
          :icon="Delete"
          type="danger"
          link
          @click="handleDeleteClick"
        >删除</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, toRefs } from "vue";
import { useDictStore } from "~/store/dict";
import { storeToRefs } from "pinia";
import { Edit, Delete } from '@element-plus/icons-vue';

const props = defineProps({
  data: {
    type: Object,
    default: () => {
      return {
        id: 0, // ID
        code: "", // 编码
        name: "", // 名称
        content_type: 2, // 内容类型 1:Json 2:Text 3:Anko Script
        remark: "", // 备注
      };
    },
  },
});

const dictStore = useDictStore();
const { activeData } = storeToRefs(dictStore);

const emits = defineEmits(["edit", "delete", "cell"]);

const { data } = toRefs(props);

const isShowActive = computed(() => {
  if (!activeData.value) return false;
  return data.value.id === activeData.value.id;
});

const handleEditClick = () => {
  emits("edit", data.value);
};

const handleDeleteClick = () => {
  emits("delete", data.value);
};

const handleCellClick = () => {
  emits("cell", data.value);
};
</script>