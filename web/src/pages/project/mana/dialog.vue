<template>
  <!--弹窗-->
  <el-dialog
    v-model="visible"
    :title="dialog.title"
    width="600px"
    @close="handleClose"
    @open="handleOpen"
    :close-on-click-modal="false"
  >
    <el-form
      ref="dataFormRef"
      :model="formData"
      :rules="computedRules"
      label-width="auto"
    >
      <el-form-item label="名称" prop="name">
        <el-input v-model="formData.name" placeholder="请输入名称" />
      </el-form-item>

      <el-form-item label="仓库地址" prop="repo_url">
        <el-input v-model="formData.repo_url" placeholder="请输入仓库地址" />
      </el-form-item>

      <el-form-item label="机器人">
        <el-select
          v-model="formData.robots"
          multiple
          collapse-tags
          placeholder="请选择机器人"
        >
          <el-option
            v-for="(item, index) in robotOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="备注">
        <el-input
          v-model="formData.description"
          type="textarea"
          :rows="3"
          placeholder="请输入备注"
        />
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
import { apiAddData, apiUpdateData } from "~/api/project";
import { apiGetList } from "~/api/project/robot";
import { reactive, ref, toRefs, computed } from "vue";

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
        name: "", // 名称
        repo_url: "", // 仓库地址
        description: "", // 备注
        robots: [], // 机器人
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

const robotOptions = ref([]);

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

const getRobotList = async () => {
  try {
    const res = await apiGetList();
    robotOptions.value = res.list.map((item) => {
      return {
        label: item.name,
        value: item.id,
      };
    });
  } catch (error) {
    return [];
  }
};

/**
 * 关闭弹窗
 */
const handleClose = () => {
  emits("close");
};

const handleOpen = () => {
  getRobotList();
  switch (operationType.value) {
    case 0:
      dialog.title = "新增项目信息";
      dataFormRef.value.clearValidate();
      break;
    case 1:
      dialog.title = "新增项目信息";
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
      emits("close");
    }
  });
};

/**
 * 新增数据
 */
const addData = async () => {
  try {
    dialog.loading = true;
    await apiAddData(formData.value);
    dialog.loading = false;
    toast("新增成功");
  } catch (error) {}
};

/**
 * 修改数据
 */
const editData = async () => {
  try {
    dialog.loading = true;
    await apiUpdateData(formData.value);
    dialog.loading = false;
    toast("修改成功");
  } catch (error) {}
};
</script>