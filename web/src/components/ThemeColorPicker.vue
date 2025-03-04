<template>
  <el-drawer
    v-model="drawerVisible"
    title="主题设置"
    size="300px"
    :destroy-on-close="false"
  >
    <div class="p-4">
      <div class="flex items-center justify-center mb-4">
        <el-color-picker
          v-model="themeColor"
          :predefine="predefineColors"
          show-alpha
          @change="handleColorChange"
        />
      </div>
      <div class="grid grid-cols-4 gap-2">
        <div
          v-for="(color, index) in predefineColors"
          :key="index"
          class="w-12 h-12 rounded cursor-pointer hover:ring-2 hover:ring-offset-2 hover:ring-primary-500"
          :style="{ backgroundColor: color }"
          @click="handlePresetColorClick(color)"
        ></div>
      </div>
    </div>
  </el-drawer>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { ThemeManager } from '../utils/theme'
import { useUserStore } from "~/store"; // 导入 Pinia store
import { storeToRefs } from 'pinia';

const drawerVisible = ref(false)

const userStore = useUserStore(); // 使用 Pinia store
const { themeColor } = storeToRefs(userStore); // 解构出 userStore 中的 state 和 actions

// 预定义的主题颜色
const predefineColors = [
  '#2177b8',
  '#409EFF',
  '#009688',
  '#536DFE',
  '#FF5722',
  '#FFB300',
  '#00BCD4',
  '#673AB7'
]

// 处理颜色变化
const handleColorChange = (color) => {
  if (color) {
    ThemeManager.setTheme(color)
    // 更新本地颜色值
    themeColor.value = color
  }
}

// 处理预设颜色点击
const handlePresetColorClick = (color) => {
  handleColorChange(color)
}

// 监听颜色变化
watch(themeColor, (newColor) => {
  if (newColor) {
    handleColorChange(newColor)
  }
}, { immediate: true })

// 打开抽屉的方法
const open = () => {
  drawerVisible.value = true
  // 打开时同步当前主题颜色
  // 写入到 pinia menuColor 中

  // const currentColor = localStorage.getItem('theme-primary-color')
  if (currentColor) {
    themeColor.value = currentColor
  }
}

onMounted(() => {
  // 初始化主题
  ThemeManager.initTheme()
  // 如果有保存的主题颜色，更新选择器的值
  const savedColor = localStorage.getItem('theme-primary-color')
  if (savedColor) {
    themeColor.value = savedColor
  }
})

// 暴露方法给父组件
defineExpose({
  open
})
</script>

<style scoped>
.el-color-picker {
  margin: 0 8px;
}
</style>