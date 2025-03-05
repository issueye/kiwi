<template>
  <!--弹窗-->
  <el-dialog v-model="visible" :title="dialog.title" width="600px" @close="handleClose" @open="handleOpen"
    :close-on-click-modal="false">
    <el-form ref="dataFormRef" :model="formData" :rules="computedRules" label-width="auto">
      <el-form-item label="名称" prop="name">
        <el-input v-model="formData.name" placeholder="请输入名称" />
      </el-form-item>

      <el-form-item label="机器人类型" prop="robot_type">
        <el-select v-model="formData.robot_type" placeholder="请选择机器人类型">
          <el-option v-for="item in robotTypeList" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>

      <el-form-item label="Webhook地址">
        <el-input v-model="formData.webhook_url" placeholder="请输入webhook_url" />
      </el-form-item>

      <el-form-item label="TOKEN">
        <el-input v-model="formData.token" placeholder="请输入TOKEN" />
      </el-form-item>

      <el-form-item label="SECRET">
        <el-input v-model="formData.secret" placeholder="请输入SECRET" />
      </el-form-item>

      <el-form-item label="备注">
        <el-input v-model="formData.description" type="textarea" :rows="2" placeholder="请输入备注" />
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
import { apiAddData, apiUpdateData } from "~/api/project/robot";
import { apiGetDictDetails } from "~/api/dict";
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
        name: "",
        robot_type: "",
        webhook_url: "",
        token: "",
        secret: "",
        remark: "",
      };
    },
  },
});

const { visible, operationType, formData } = toRefs(props);

const emits = defineEmits(["update:visible", "open", "close"]);
const dataFormRef = ref(null);
const robotTypeList = ref([]);

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
    robot_type: [{ required: true, message: "请选择机器人类型", trigger: "blur" }],
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
      dialog.title = "新增机器人信息";
      dataFormRef.value.clearValidate();
      break;
    case 1:
      dialog.title = "修改机器人信息";
      break;
  }

  getRobotTypeList();
};

const getRobotTypeList = async () => {
  const data = await apiGetDictDetails("10000");
  robotTypeList.value = [];
  robotTypeList.value = data.list.map((e) => {
    return {
      label: e.key,
      value: e.val,
    }
  });
}

/**
 * 提交表单按钮
 */
const handleSubmitClick = () => {
  dataFormRef.value.validate(async (isValid) => {
    console.log('operationType.value', operationType.value);
    
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
    await apiAddData(formData.value);
    dialog.loading = false;
    toast("新增代码生成配置成功");
  } catch (error) { }
};

/**
 * 修改数据
 */
const editData = async () => {
  try {
    dialog.loading = true;
    await apiUpdateData(formData.value);
    dialog.loading = false;
    toast("修改代码生成配置成功");
  } catch (error) { }
};
</script>