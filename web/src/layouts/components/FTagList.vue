<template>
  <div class="bg-gray-100 flex items-center px-2">
    <el-tabs v-model="activeTab" type="card" @tab-remove="removeTab" style="min-width: 100px"
      @tab-change="changeTab">
      <el-tab-pane v-for="item in tabList" :label="item.title" :key="item.path" :name="item.path"
        :closable="item.path !== '/home'">
        <template #label>
          <el-dropdown trigger="contextmenu" :id="item.path" @visible-change="handleChange($event, item.path)"
            ref="dropdownRef" @command="handleCommand">
            <span :class="activeTab === item.path ? 'text-header-500' : ''">{{ item.title }}</span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item :command="{ type: 'close', data: item }" v-if="item.path !== '/home'">
                  <el-icon>
                    <Close />
                  </el-icon>
                  <span>关闭</span>
                </el-dropdown-item>
                <el-dropdown-item :command="{ type: 'closeOther', data: item }">
                  <el-icon>
                    <Operation />
                  </el-icon>
                  <span>关闭其他</span>
                </el-dropdown-item>
                <el-dropdown-item :command="{ type: 'closeAll', data: item }">
                  <el-icon>
                    <Minus />
                  </el-icon>
                  <span>关闭全部</span>
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </el-tab-pane>
    </el-tabs>

    <span class="h-[32px] bg-white rounded-sm shadow ml-auto flex items-center justify-center px-2">
      <el-dropdown @command="handleClose">
        <span class="el-dropdown-link">
          <el-icon>
            <arrow-down />
          </el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="clearOther">
              <el-icon>
                <Operation />
              </el-icon>
              <span>关闭其他</span>
            </el-dropdown-item>
            <el-dropdown-item command="clearAll">
              <el-icon>
                <Minus />
              </el-icon>
              <span>关闭全部</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </span>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useTabList } from "~/composables/useTabList.js";
import { useUserStore } from "~/store"; // 导入 Pinia 的 store

const userStore = useUserStore(); // 获取 store 实例
const { activeTab, tabList, changeTab, removeTab, removeAll,removeOther, handleClose } = useTabList();
const dropdownRef = ref();

const handleChange = (visible, path) => {
  // console.log('handleChange', e, path, 'activeTab', activeTab.value);
  if (!visible) return;

  dropdownRef.value.forEach((item) => {
    if (item.id !== path) {
      item.handleClose();
    }
  })
};

const handleCommand = ({ type, data }) => {
  switch (type) {
    case 'close':
      {
        removeTab(data.path);
        break;
      }
    case 'closeOther':
      {
        removeOther(data.path);
        break;
      }
    case 'closeAll':
      {
        removeAll();
        break;
      }
  }
  // if (type === 'close') {
  //   removeTab(data);
  // } else if (type === 'closeOther') {
  //   removeTab(data, true);
  // } else if (type === 'closeAll') {
  //   removeTab(data, false, true);
  // }
};

</script>

<style scoped>
:deep(.el-tabs__nav-scroll) {
  padding: 2px 0 2px 0;
}

:deep(.el-tabs__header) {
  border: 0 !important;
  @apply mb-0;
}

:deep(.el-tabs__nav) {
  border: 0 !important;
}

:deep(.el-tabs__item) {
  border: 0 !important;
  height: 32px;
  line-height: 32px;
  @apply bg-white mx-1 rounded-sm shadow;
}

:deep(.el-tabs__nav-next),
:deep(.el-tabs__nav-prev) {
  line-height: 32px;
  height: 32px;
}

:deep(.is-disabled) {
  cursor: not-allowed;
  /* @apply text-gray-300; */
}
</style>
