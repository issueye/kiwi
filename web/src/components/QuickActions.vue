<template>
  <el-card shadow="hover">
    <template #header>
      <div class="flex justify-between items-center">
        <span class="text-lg font-bold">快捷操作</span>
        <el-icon><Operation /></el-icon>
      </div>
    </template>
    <div class="grid grid-cols-2 gap-4">
      <el-button type="primary" icon="Upload" @click="emit('compile')">立即编译</el-button>
      <el-button type="success" icon="Download" @click="emit('download')">下载文件</el-button>
    </div>
  </el-card>
</template>

<script setup>
const emit = defineEmits(['compile', 'download'])
</script>