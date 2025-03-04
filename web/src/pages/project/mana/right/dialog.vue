<template>
  <!--弹窗-->
  <el-dialog v-model="visible" :title="dialog.title" width="700px" @close="handleClose" @open="handleOpen"
    :close-on-click-modal="false">
    <el-form ref="dataFormRef" :model="formData" :rules="computedRules" label-width="auto">
      <el-form-item label="版本号" prop="version">
        <el-input v-model="formData.version" placeholder="请输入版本号" />
      </el-form-item>

      <el-form-item label="分支名称" prop="branch_name">
        <el-select v-model="formData.branch_name" placeholder="请选择分支名称">
          <el-option v-for="item in branchList" :key="item.id" :label="item.name" :value="item.name"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="HASH" prop="hash">
        <el-input v-model="formData.hash" placeholder="请输入HASH" />
      </el-form-item>

      <el-form-item label="更新内容" prop="update_content">
        <el-input v-model="formData.update_content" type="textarea" :rows="10" placeholder="请输入更新内容" />
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button type="primary" @click="handleSubmitClick">确 定</el-button>
        <el-button @click="handleClose">取 消</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { storeToRefs } from "pinia";
import { reactive, ref, toRefs, computed } from "vue";

import { useProjectStore } from '~/store/project';

const projectStore = useProjectStore();
const { branchList } = storeToRefs(projectStore);
console.log('branchList', branchList);


const props = defineProps({
  // 操作类型 0:新增 1:编辑
  operationType: {
    type: Number,
    default: 0,
  },
  // 是否显示
  visible: {
    type: Boolean,
    default: false,
  },
  // 获取表单数据
  formData: {
    type: Object,
    default: () => {
      return {
        id: 0, // ID
        version: "", // 版本
        branch_name: "", // 分支名称
        hash: "", // HASH
        project_id: 0, // 项目ID
        description: "", // 备注
      };
    },
  },
});

const { visible, operationType, formData } = toRefs(props);

const emits = defineEmits(["update:visible", "close"]);
const dataFormRef = ref(null);

const dialog = reactive({
  title: "",
  loading: false,
});


/**
 * 表单验证规则
 */
const computedRules = computed(() => {
  const rules = {
    name: [{ required: true, message: "请输入名称", trigger: "blur" }],
    repo_url: [{ required: true, message: "请输入仓库地址", trigger: "blur" }],
  };
  return rules;
});

/**
 * 关闭弹窗
 */
const handleClose = () => {
  emits("update:visible", false);
  emits("close");
};

const handleOpen = () => {
  projectStore.getBranchList();

  switch (operationType.value) {
    case 0:
      dialog.title = "新增版本信息";
      dataFormRef.value.clearValidate();
      break;
    case 1:
      dialog.title = "新增版本信息";
      break;
  }
};

/**
 * 提交表单按钮
 */
const handleSubmitClick = () => {
  dataFormRef.value.validate(async (isValid) => {
    if (isValid) {
      switch (operationType.value) {
        case 0:
          await addData();
          break;
        case 1:
          await editData();
          break;
      }
      emits("update:visible", false);
    }
  });
};

/**
 * 新增数据
 */
const addData = async () => {
  dialog.loading = true;
  await projectStore.addVersionData(formData.value);
  dialog.loading = false;
};

/**
 * 修改数据
 */
const editData = async () => {
  dialog.loading = true;
  await projectStore.updateVersionData(formData.value);
  dialog.loading = false;
};
</script>