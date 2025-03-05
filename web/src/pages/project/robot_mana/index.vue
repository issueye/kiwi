<template>
  <base-page title="机器人管理" desc="机器人管理">
    <template #content>
      <div class="h-full flex flex-col p-2">
        <div class="search-bar">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true">
            <el-form-item label="关键字" prop="keywords">
              <el-input v-model="queryParams.keywords" placeholder="名称/编码" clearable @keyup.enter="handleQuery" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="search" @click="handleQuery">搜索</el-button>
              <el-button icon="refresh" @click="handleResetQuery">重置</el-button>
            </el-form-item>
          </el-form>
        </div>

        <div class="grow flex flex-col">
          <div class="mb-[10px]">
            <el-button type="success" icon="plus" @click="handleAddClick">新增</el-button>
          </div>

          <div class="grow">
            <d-table :columns="columns" :table-data="tableData" :page-config="pageConfig" usePagination
              highlight-current-row stripe :loading="loading" empty-text="暂无数据">
              <template #operation="{ scope }">
                <el-button type="primary" link @click="handleScriptClick(scope.row)">脚本</el-button>
                <el-divider direction="vertical" />
                <el-button type="primary" link @click="handleEditClick(scope.row)">编辑</el-button>
                <el-divider direction="vertical" />
                <el-button type="danger" link @click="handleDeleteClick(scope.row)">删除</el-button>
              </template>
            </d-table>
          </div>
        </div>
        <Dialog v-model:visible="dialog.visible" v-model:form-data="formData" :operation-type="dialog.operationType"
          @close="handleDialogClose" />

        <ScriptDialog v-model:visible="scriptDialog.visible" v-model:data="scriptDialog.formData" />
      </div>
    </template>
  </base-page>
</template>

<script setup>
import { apiGetList, apiDeleteById } from "~/api/project/robot";
import Dialog from "./dialog.vue";
import ScriptDialog from "~/components/script_dialog.vue";

import { ElMessageBox, ElMessage } from "element-plus";
import { toast } from "~/composables/util";

import { ref, reactive, computed, onMounted } from "vue";
const queryFormRef = ref();

const loading = ref(false);

// 表格信息
const columns = [
  { prop: "name", label: "名称", attrs: { width: 230 } },
  { prop: "robot_type", label: "机器人类型", attrs: { width: 200 } },
  { prop: "webhook_url", label: "WEBHOOK", attrs: { minWidth: 200, showOverflowTooltip: true } },
  { prop: "token", label: "TOKEN", attrs: { minWidth: 200 } },
  { prop: "secret", label: "SECRET", attrs: { minWidth: 200 } },
  {
    prop: "operation",
    label: "操作",
    slot: true,
    attrs: { width: 170, fixed: "right" },
  },
];

const pageConfig = reactive({
  pageSize: 10,
  currentPage: 1,
  total: 0,
});

const tableData = ref([]);

const queryParams = reactive({
  keywords: "",
});

const dialog = reactive({
  title: "",
  visible: false,
  operationType: 0,
});

const scriptDialog = reactive({
  visible: false,
  formData: {
    name: "",
    content: "",
    type: "text/javascript",
  },
})

const formData = reactive({
  id: 0, // ID
  name: "",
  robot_type: "",
  webhook_url: "",
  token: "",
  secret: "",
  remark: "",
});

const handleDialogClose = () => {
  getData();
}

// 查询
function handleQuery() {
  getData();
}

const getData = async () => {
  loading.value = true;
  try {
    const data = await apiGetList({
      pageNum: pageConfig.currentPage,
      pageSize: pageConfig.pageSize,
      condition: queryParams,
    });
    loading.value = false;
    tableData.value = data.list;
    pageConfig.total = data.total;
  } catch (error) {
    loading.value = false;
  }
};

// 重置查询
function handleResetQuery() {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
  handleQuery();
}

const setValue = (value) => {
  formData.id = value.id;
  formData.name = value.name;
  formData.robot_type = value.robot_type;
  formData.webhook_url = value.webhook_url;
  formData.token = value.token;
  formData.secret = value.secret;
  formData.remark = value.remark;
};

const resetValue = () => {
  formData.id = undefined;
  formData.name = "";
  formData.robot_type = "";
  formData.webhook_url = "";
  formData.token = "";
  formData.secret = "";
  formData.remark = "";
};

// 新增字典
function handleAddClick() {
  dialog.visible = true;
  dialog.title = "新增角色";
  dialog.operationType = 0;
  resetValue();
}

/**
 * 编辑字典
 *
 * @param id 字典ID
 */
function handleEditClick(value) {
  dialog.visible = true;
  dialog.title = "修改角色";
  dialog.operationType = 1;
  setValue(value);
}

/**
 * 编辑脚本
 * @param {*} value 
 */
const handleScriptClick = (value) => {
  scriptDialog.formData.name = value.name;
  scriptDialog.formData.content = value.content;
  scriptDialog.visible = true;
}

/**
 * 删除数据
 */
function handleDeleteClick(value) {
  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(
    async () => {
      await apiDeleteById(value.id);
      toast("删除角色成功");
      handleQuery();
    },
    () => {
      ElMessage.info("已取消删除");
    }
  );
}

onMounted(() => {
  handleQuery();
});
</script>