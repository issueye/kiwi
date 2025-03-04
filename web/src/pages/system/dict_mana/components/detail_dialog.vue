<template>
  <!--弹窗-->
  <el-dialog v-model="visible" :title="dialog.title" width="700px" @close="handleClose" @open="handleOpen"
    :close-on-click-modal="false">
    <el-form ref="dataFormRef" :model="formData" :rules="computedRules" label-width="auto">
      <el-form-item label="名称" prop="key">
        <el-input v-model="formData.key" placeholder="请输入名称" />
      </el-form-item>
      <el-form-item label="值" prop="val">
        <el-input v-model="formData.val" placeholder="请输入值" type="textarea" :rows="6" />
      </el-form-item>
      <el-form-item label="备注">
        <el-input v-model="formData.remark" placeholder="请输入备注" type="textarea" :rows="2" />
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
import { apiSaveDetail } from "~/api/dict";
import { reactive, ref, toRefs, computed } from "vue";
import { useDictStore } from "~/store/dict";
import { storeToRefs } from "pinia";
import { toast } from "~/composables/util";

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
        dict_code: "", // 字典编码
        key: "", // 编码
        val: "", // 名称
        remark: "", // 备注
      };
    },
  },
});

const { visible, operationType, formData } = toRefs(props);

const emits = defineEmits(["update:visible", "close"]);
const dataFormRef = ref(null);

const dictStore = useDictStore();
const { activeData } = storeToRefs(dictStore);

const dialog = reactive({
  title: "",
  loading: false,
});

/**
 * 表单验证规则
 */
const computedRules = computed(() => {
  const rules = {
    key: [{ required: true, message: "请输入编码", trigger: "blur" }],
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
  switch (operationType.value) {
    case 0:
      dialog.title = `字典[${activeData.value.name}]添加`;
      dataFormRef.value.clearValidate();
      break;
    case 1:
      dialog.title = `字典[${activeData.value.name}]编辑`;
      break;
  }
};

/**
 * 提交表单按钮
 */
const handleSubmitClick = () => {
  dataFormRef.value.validate(async (isValid) => {
    if (isValid) {
      formData.value.dict_code = activeData.value.code;
      await apiSaveDetail(formData.value);
      toast("保存成功");
      emits("update:visible", false);
    }
  });
};
</script>