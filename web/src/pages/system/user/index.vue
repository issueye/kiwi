<template>
  <base-page title="用户管理" desc="用户管理">
    <template #actions>
      <el-button type="success" icon="plus" @click="onAddClick">新增</el-button>
    </template>
    <template #content>
      <div class="h-full flex flex-col p-2">
        <div class="search-bar">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true">
            <el-form-item label="关键字" prop="keywords">
              <el-input v-model="queryParams.keywords" placeholder="名称/编码" clearable @keyup.enter="onQuery" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="search" @click="onQuery">搜索</el-button>
              <el-button icon="refresh" @click="onResetQuery">重置</el-button>
            </el-form-item>
          </el-form>
        </div>

        <div class="grow flex flex-col">
          <div class="grow">
            <d-table :columns="columns" :table-data="tableData" :page-config="pageConfig" usePagination
              highlight-current-row stripe :loading="loading" empty-text="暂无数据">
              <template #role="{ scope }">
                <el-tag size="small" effect="plain">
                  {{ getRoleName(scope.row.UserRole.role_code) }}
                </el-tag>
              </template>
              <template #operation="{ scope }">
                <el-button type="primary" link @click="onEditClick(scope.row)">编辑</el-button>
                <el-divider direction="vertical" />
                <el-button type="danger" link @click="onDeleteClick(scope.row)">删除</el-button>
              </template>
            </d-table>
          </div>
        </div>

        <!--弹窗-->
        <Dialog v-model:visible="dialog.visible" v-model:dialog="dialog" v-model:dataForm="formData" @close="onClose" />
      </div>
    </template>
  </base-page>
</template>

<script setup>
import {
  apiGetUserList,
  apiDeleteUser,
} from "~/api/user";
import Dialog from './dialog.vue';
import { storeToRefs } from 'pinia';

import { ElMessageBox, ElMessage } from "element-plus";
import { toast } from "~/composables/util";

import { ref, reactive, onMounted } from "vue";

import { useSystemStore } from '~/store/system';

const sysStore = useSystemStore()
const { roleList } = storeToRefs(sysStore);

const queryFormRef = ref();

const loading = ref(false);

/**
 * 表格信息
 */
const columns = [
  { prop: "username", label: "用户名", attrs: { width: 150 } },
  { prop: "nick_name", label: "昵称", attrs: { width: 230 } },
  { prop: "role", label: "角色", slot: true, attrs: { width: 230 } },
  { prop: "remark", label: "备注", attrs: { minWidth: 200 } },
  {
    prop: "operation",
    label: "操作",
    slot: true,
    attrs: { width: 120, fixed: "right" },
  },
];

/**
 * 分页查询
 */
const pageConfig = reactive({
  pageSize: 10,
  currentPage: 1,
  total: 0,
});

/**
 * 获取表格数据
 */
const tableData = ref([]);

/**
 * 查询条件
 */
const queryParams = reactive({
  keywords: "",
});

const dialog = reactive({
  title: "",
  visible: false,
  operationType: 0, // 0:新增 1:编辑
});

/**
 * 初始化数据
 */
onMounted(() => {
  onQuery();
  sysStore.getRoleList();
});

/**
 * 表单数据
 */
const formData = reactive({
  id: 0,
  username: "",
  nick_name: "",
  role_code: "",
  remark: "",
});

const getRoleName = (code) => {
  const role = roleList.value.find((item) => item.code === code);
  return role ? role.name : "";
};

/**
 * 查询数据
 */
const onQuery = () => {
  getData();
};

/**
 * 获取数据
 */
const getData = async () => {
  loading.value = true;
  try {
    const data = await apiGetUserList({
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

const onClose = () => {
  dialog.visible = false;
  getData();
}

/**
 * 重置查询
 */
const onResetQuery = () => {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
  getData();
};

/**
 * 新增
 */
const onAddClick = () => {
  resetValue();
  dialog.title = "新增用户";
  dialog.operationType = 0;
  dialog.visible = true;
  console.log('dialog', dialog);
};

/**
 * 编辑
 * @param value 信息
 */
const onEditClick = (value) => {
  dialog.title = "修改用户";
  dialog.operationType = 1;
  setValue(value);
  dialog.visible = true;
  console.log('dialog', dialog);
};

/**
 * 设置值
 * @param value 信息
 */
const setValue = (value) => {
  formData.id = value.id;
  formData.username = value.username;
  formData.nick_name = value.nick_name;
  formData.remark = value.remark;
  formData.role_code = value.UserRole.role_code;
};

/**
 * 重置值
 */
const resetValue = () => {
  formData.id = undefined;
  formData.username = "";
  formData.nick_name = "";
  formData.remark = "";
  formData.role_code = "";
};


/**
 * 删除数据
 * @param value 信息
 */
const onDeleteClick = (value) => {
  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(
    async () => {
      await apiDeleteUser(value.id);
      toast("删除用户成功");
      getData();
    },
    () => {
      ElMessage.info("已取消删除");
    }
  );
};
</script>
