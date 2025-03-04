<template>
    <div class="h-full flex flex-col">
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
                        <el-button type="primary" link @click="handleEditClick(scope.row)">编辑</el-button>
                        <el-divider direction="vertical" />
                        <el-button type="warning" link icon="Promotion" @click="handleCompileClick(scope.row)">编译</el-button>
                        <el-divider direction="vertical" />
                        <el-dropdown @command="handleCommand">
                            <span class="flex items-center text-[--el-color-primary]">
                                更多<el-icon class="el-icon--right"><arrow-down /></el-icon>
                            </span>
                            <template #dropdown>
                                <el-dropdown-menu>
                                    <el-dropdown-item :command="{ data: scope.row, type: 'del' }"
                                        class="text-[--el-color-danger]">删除</el-dropdown-item>
                                </el-dropdown-menu>
                            </template>
                        </el-dropdown>
                    </template>
                </d-table>
            </div>
        </div>
    </div>
    <Dialog v-model:visible="dialog.visible" v-model:form-data="formData" :operation-type="dialog.operationType"
        @close="handleDialogClose" />
    
</template>
<script setup>
import { ref, reactive, onMounted } from "vue";
import Dialog from "./dialog.vue";
import { toast } from "~/composables/util";
import { useRouter } from 'vue-router';

import { useProjectStore } from '~/store/project';
import { storeToRefs } from "pinia";

const router = useRouter();
const projectStore = useProjectStore();

const {
    versionList:tableData,
    versionPageConfig:pageConfig,
    versionQryParams:queryParams,

} = storeToRefs(projectStore);

const queryFormRef = ref(null);

const dialog = reactive({
    visible: false,
    operationType: 0, // 0:新增 1:编辑
    compileDialogVisible: false, // 编译对话框
});

const loading = ref(false);

/**
 * 表格信息
 */
const columns = [
    {
        prop: "version",
        label: "版本",
        attrs: { width: 300, showOverflowTooltip: true },
    },
    {
        prop: "branch_name",
        label: "分支名称",
        attrs: { width: 200, showOverflowTooltip: true },
    },
    {
        prop: "hash",
        label: "HASH",
        attrs: { minWidth: 200, showOverflowTooltip: true },
    },
    {
        prop: "operation",
        label: "操作",
        slot: true,
        attrs: { width: 200, fixed: "right" },
    },
];

const formData = reactive({
    id: 0, // ID
    version: "", // 版本
    branch_name: "", // 分支名称
    hash: "", // HASH
    project_id: 0, // 项目ID
    description: "", // 备注
});

onMounted(() => {
    getData();
});

/**
 * 查询
 */
const handleQuery = () => {
    getData();
};

/**
 * 重置查询
 */
const handleResetQuery = () => {
    queryFormRef.value.resetFields();
    pageConfig.value.currentPage = 1;
    handleQuery();
};

/**
 * 新增
 */
const handleAddClick = () => {
    dialog.operationType = 0;
    resetValue();
    dialog.visible = true;
};

/**
 * 修改信息
 * @param value 信息
 */
const handleEditClick = (value) => {
    dialog.operationType = 1;
    setValue(value);
    dialog.visible = true;
};

const handleCompileClick = (value) => {
    console.log('router', router);
    router.push({ path: '/project/compile', query: { version_id: value.id, id: value.project_id } });
};

const setValue = (value) => {
    console.log('value', value);
    
    formData.id = value.id;
    formData.version = value.version;
    formData.branch_name = value.branch_name;
    formData.hash = value.hash;
    formData.project_id = value.project_id;
    formData.description = value.description;
};

const resetValue = () => {
    formData.id = 0;
    formData.version = "";
    formData.branch_name = "";
    formData.hash = "";
    formData.project_id = "";
    formData.description = "";
};

/**
 * 删除
 * @param value
 */
const handleDeleteClick = async (value) => {
    loading.value = true;
    await projectStore.deleteVersionById(value.id);
    loading.value = false;
};

const handleCommand = ({ data, type }) => {
    if (type === "del") {
        handleDeleteClick(data);
    }
};

const handleDialogClose = () => {
    getData();
};

/**
 * 查询数据
 */
const getData = async () => {
    loading.value = true;
    try {
        await projectStore.getVersionList();
    } catch (error) {
        console.log(error);
    } finally {
        loading.value = false;
    }
};

</script>