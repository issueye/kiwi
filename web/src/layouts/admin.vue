<template>
  <div class="h-full flex flex-col" style="min-width: 856px">
    <!-- 页面头部 -->
    <div :style="{ height: global.CARAMBOLA_HEADER_HEIGHT }">
      <f-header />
    </div>
    <div
      class="flex"
      :style="{ height: `calc(100% - ${global.CARAMBOLA_HEADER_HEIGHT})` }"
    >
      <div class="h-full flex flex-col bg-gray-100 w-full overflow-y-auto">
        <div class="h-full m-3">
          <router-view />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import FHeader from "./components/FHeader.vue";
import { global } from "~/init/global";
</script>

<style scoped>
.el-aside {
  transition: all 0.3s;
}

.fade-enter-from {
  opacity: 0;
}

.fade-enter-to {
  opacity: 1;
}

.fade-leave-from {
  opacity: 1;
}

.fade-leave-to {
  opacity: 0;
}

.fade-enter-active,
.fade-leave-active {
  transition: all 0.3s;
}

.fade-enter-active {
  transition-delay: 0.3s;
}
</style>