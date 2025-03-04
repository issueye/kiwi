<template>
  <base-page title="项目概览" padding="20px">
    <template #content>
      <el-row :gutter="20" class="mb-4">
        <el-col :span="12">
          <project-stats />
        </el-col>
        <el-col :span="12">
          <recent-activities />
        </el-col>
      </el-row>

      <el-row :gutter="20">
        <el-col :span="24">
          <quick-actions @compile="handleCompile" @download="handleDownload" />
        </el-col>
      </el-row>
    </template>
  </base-page>
</template>

<script setup>
import ProjectStats from "~/components/ProjectStats.vue";
import RecentActivities from "~/components/RecentActivities.vue";
import QuickActions from "~/components/QuickActions.vue";

const handleCompile = () => {
  // 编译处理逻辑
};

const handleDownload = () => {
  // 下载处理逻辑
};
</script>

<style scoped>
.el-card {
  margin-bottom: 20px;
}
</style>
  