<template>
  <el-card shadow="hover" class="h-full">
    <template #header>
      <div class="flex justify-between items-center">
        <span class="text-lg font-bold">最近操作</span>
        <el-icon><Clock /></el-icon>
      </div>
    </template>
    <el-timeline>
      <el-timeline-item
        v-for="(item, index) in activities"
        :key="index"
        :timestamp="item.time"
      >
        {{ item.content }}
      </el-timeline-item>
    </el-timeline>
  </el-card>
</template>

<script setup>
import { ref } from "vue";
const activities = ref([
  {
    time: "2024-03-15 14:30",
    content: '用户张三创建了新项目 "AI平台开发"',
  },
  {
    time: "2024-03-15 10:15",
    content: "用户李四更新了项目文档",
  },
]);
</script>