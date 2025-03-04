<template>
  <!--弹窗-->
  <el-dialog
    v-model="visible"
    :title="dialog.title"
    width="500px"
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
      <el-form-item label="编码" prop="code">
        <el-input v-model="formData.code" placeholder="请输入编码" />
      </el-form-item>
      <el-form-item label="名称" prop="name">
        <el-input v-model="formData.name" placeholder="请输入名称" />
      </el-form-item>
      <el-form-item label="内容类型" prop="content_type">
        <el-radio-group v-model="formData.content_type" :disabled="operationType != 0">
          <el-radio-button :value="2" label="Text" />
          <el-radio-button :value="1" label="Json" />
          <el-radio-button :value="3" label="Anko Script" />
        </el-radio-group>
      </el-form-item>
      <el-form-item label="备注">
        <el-input
          v-model="formData.remark"
          placeholder="请输入备注"
          type="textarea"
          :rows="4"
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
import { apiAddData, apiUpdateData } from "~/api/dict";
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
        code: "", // 编码
        name: "", // 名称
        content_type: 2, // 内容类型 1:Json 2:Text 3:Anko Script
        remark: "", // 备注
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
    code: [{ required: true, message: "请输入编码", trigger: "blur" }],
    name: [{ required: true, message: "请输入名称", trigger: "blur" }],
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
  console.log('formData.value', formData.value);
  
  switch (operationType.value) {
    case 0:
      dialog.title = "新增字典信息";
      dataFormRef.value.clearValidate();
      break;
    case 1:
      dialog.title = "编辑字典信息";
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
  try {
    dialog.loading = true;
    // 接口调用
    await apiAddData(formData.value);
    dialog.loading = false;
    toast("新增字典信息成功");
  } catch (error) {}
};

/**
 * 修改数据
 */
const editData = async () => {
  try {
    dialog.loading = true;
    // 接口调用
    await apiUpdateData(formData.value);
    dialog.loading = false;
    toast("修改字典信息成功");
  } catch (error) {}
};
</script>