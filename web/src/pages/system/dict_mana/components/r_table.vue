<template>
  <div>
    <div class="flex justify-between">
      <el-input class="w-52" placeholder="检索内容" :suffix-icon="Search" />
    </div>
    <div class="mt-3">
      <el-button type="success" icon="plus" @click="handleAddClick">添加</el-button>
    </div>
  </div>
  <div class="mt-3">
    <d-table :columns="columns" :table-data="dictDetailData" :page-config="pageConfig" usePagination
      highlight-current-row stripe :loading="loading" empty-text="暂无数据">
      <template #operation="{ scope }">
        <el-button type="primary" link @click="handleEditClick(scope.row)">编辑</el-button>
        <el-divider direction="vertical" />
        <el-dropdown @command="handleCommand">
          <span class="flex items-center text-[--el-color-primary]">
            更多<el-icon class="el-icon--right"><arrow-down /></el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item :command="{ data: scope.row, type: 'edit_content' }" class="text-[--el-color-danger]"
                v-if="activeData.content_type == 3">
                编辑内容
              </el-dropdown-item>
              <el-dropdown-item :command="{ data: scope.row, type: 'del' }"
                class="text-[--el-color-danger]">删除</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </template>
    </d-table>
  </div>
  <Dialog v-model:visible="dialog.visible" v-model:form-data="formData" :operation-type="dialog.operationType"
    @close="handleDialogClose" />
  <ScriptDialog v-model:visible="scriptDialog.visible" :data="scriptDialog.data" @close="scriptDialog.closeFn" />
</template>
<script setup>
import { ref, reactive } from "vue";
import { Search } from "@element-plus/icons-vue";
import { toast } from "~/composables/util";
import Dialog from "./detail_dialog.vue";
import { useDictStore } from "~/store/dict";
import { storeToRefs } from "pinia";
import { apiDeleteDetail } from "~/api/dict";
import ScriptDialog from "~/components/script_dialog.vue";
import { apiSaveDetail } from "~/api/dict";

const dictStore = useDictStore();
const { activeData, dictDetailData, pageConfig } = storeToRefs(dictStore);

const columns = [
  { prop: "key", label: "名称", attrs: { width: 230 } },
  { prop: "val", label: "值", attrs: { minWidth: 230, showOverflowTooltip: true } },
  { prop: "remark", label: "备注", attrs: { minWidth: 200 } },
  {
    prop: "operation",
    label: "操作",
    slot: true,
    attrs: { width: 140, fixed: "right" },
  },
];

const loading = ref(false);

const dialog = reactive({
  visible: false,
  operationType: 0,
});

const formData = reactive({
  id: 0, // ID
  code: "", // 字典编码
  key: "", // 编码
  val: "", // 名称
  remark: "", // 备注
});

const scriptDialog = reactive({
  visible: false,
  data: {
    name: "",
    title: "",
    content: "",
  },
  closeFn: async (value, closeType) => {
    if (closeType === 'cancel') {
      return;
    }
    formData.val = value;
    await apiSaveDetail(formData);
    toast("保存成功");
  }
})

const setData = (value) => {
  formData.id = value.id;
  formData.code = value.code;
  formData.key = value.key;
  formData.val = value.val;
  formData.remark = value.remark;
};

const reset = () => {
  formData.id = 0;
  formData.code = activeData.value.code;
  formData.key = "";
  formData.val = "";
  formData.remark = "";
};

const handleEditClick = (value) => {
  dialog.operationType = 1;
  setData(value);
  dialog.visible = true;
};

const handleAddClick = (value) => {
  dialog.operationType = 0;
  reset();
  dialog.visible = true;
};

const handleCommand = (value) => {
  switch (value.type) {
    case "del": {
      delData(value.data);
      break;
    }
    case "edit_content": {
      scriptDialog.data.title = "编辑内容";
      scriptDialog.data.name = value.data.key;
      scriptDialog.data.content = value.data.val;
      scriptDialog.visible = true;
      break;
    }
  }
};

const delData = (value) => {
  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(
      async () => {
        await apiDeleteDetail(value.id);
        toast("删除成功");
        dictStore.getDictData();
      },
      () => {
        ElMessage.info("已取消删除");
      }
    )
    .finally(() => { });
};

const handleDialogClose = () => {
  dictStore.getDictData();
};
</script>