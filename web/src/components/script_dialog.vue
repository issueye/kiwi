<template>
  <!--弹窗-->
  <el-dialog
    v-model="visible"
    :title="computedTitle"
    width="800px"
    @close="handleClose('cancel')"
    @open="handleOpen"
    :close-on-click-modal="false"
  >
    <Codemirror
      v-model:value="data.content"
      ref="cmRef"
      :options="cmOptions"
      border
      height="500"
    />
    <template #footer>
      <div class="dialog-footer">
        <el-button type="primary" @click="handleSubmitClick">确 定</el-button>
        <el-button @click="handleClose('cancel')">取 消</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { reactive, ref, toRefs, nextTick, computed } from "vue";

import "codemirror/mode/javascript/javascript.js";
import "codemirror/mode/shell/shell.js";
import Codemirror from "codemirror-editor-vue3";
import { onMounted, onUnmounted } from "vue";

const props = defineProps({
  // 是否显示
  visible: {
    type: Boolean,
    default: false,
  },
  data: {
    type: Object,
    default: () => {
      return {
        name: "",
        content: "",
        type: "text/javascript",
      };
    },
  },
});

const { visible, data } = toRefs(props);
const emits = defineEmits(["update:visible", "close"]);

const computedTitle = computed(() => {
  // 超过20个字符显示  value...value
  if (data.value.name.length > 20) {
    // return `编辑脚本 - [${data.value.name.substring(0, 10)}...${data.value.name.substring(0, 10)}]`;
    return `编辑脚本 - [${data.value.name.substring(0, 10)}...${data.value.name.substring(data.value.name.length - 10)}]`;
  }
  
  return `编辑脚本 - [${data.value.name}]`;
})

const dialog = reactive({
  title: "编辑脚本",
  loading: false,
  closeType: "cancel", // cancel 取消 ok 确定
});

const cmRef = ref();
const cmOptions = {
  mode: "text/javascript", // 语言模式
};

onMounted(() => {
  // 在加载完成之后

  nextTick(() => {
    cmRef.value?.refresh();
  });
});

onUnmounted(() => {
  cmRef.value?.destroy();
});

/**
 * 关闭弹窗
 */
const handleClose = (closeType) => {
  if (visible.value) {
    emits("close", data.value.content, closeType);
    emits("update:visible", false);
  }
};

const handleOpen = () => {
  dialog.title = `编辑脚本 - [${data.value.name}]`;
  cmOptions.mode = data.value.type;
  cmRef.value?.refresh();
};

/**
 * 提交表单按钮
 */
const handleSubmitClick = () => {
  dialog.loading = true;

  setTimeout(() => {
    dialog.loading = false;
    handleClose("ok");
  }, 500);
};
</script>

<style scoped>
</style>