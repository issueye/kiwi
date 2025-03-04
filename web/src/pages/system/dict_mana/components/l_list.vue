<template>
  <div class="flex flex-col">
    <div class="flex justify-between">
      <el-input
        class="w-full"
        v-model="dictCondition.condition"
        placeholder="检索内容"
        :suffix-icon="Search"
        @change="handleChange"
        clearable
      />
    </div>
    <div class="mt-3">
      <el-button type="success" icon="plus" @click="handleAddClick"
        >添加</el-button
      >
    </div>
    <div class="grow my-3">
      <list_item
        v-for="(item, index) in dictData"
        :key="index"
        :data="item"
        @edit="handleEditClick"
        @delete="handleDeleteClick"
        @cell="handleCellClick"
      />
    </div>
  </div>

  <Dialog
    v-model:visible="dialog.visible"
    v-model:form-data="formData"
    :operation-type="dialog.operationType"
    @close="handleDialogClose"
  />
</template>

<script setup>
import { onMounted, reactive, ref } from "vue";
import { apiDeleteById } from "~/api/dict";
import { ElMessageBox, ElMessage } from "element-plus";
import { Search } from "@element-plus/icons-vue";
import { toast } from "~/composables/util";

import Dialog from "./dict_dialog.vue";
import list_item from "./list_item.vue";
import { useDictStore } from "~/store/dict";
import { storeToRefs } from "pinia";

const dictStore = useDictStore();
const { dictCondition, dictData } = storeToRefs(dictStore);

const dialog = reactive({
  visible: false,
  operationType: 0,
});

const formData = reactive({
  id: 0,
  code: "",
  name: "",
  content_type: 2,
  remark: "",
});

const reset = () => {
  formData.id = 0;
  formData.code = "";
  formData.name = "";
  formData.content_type = 2; // 内容类型 1:Json 2:Text 3:Anko Script
  formData.remark = "";
};

const setData = (value) => {
  formData.id = value.id;
  formData.code = value.code;
  formData.name = value.name;
  formData.content_type = value.content_type;
  formData.remark = value.remark;
};

onMounted(() => {
  getData();
});

/**
 * 获取数据
 */
const getData = async () => {
  dictStore.getDictData();
};

/**
 * 添加
 */
const handleAddClick = () => {
  dialog.operationType = 0;
  reset();
  dialog.visible = true;
};

const handleEditClick = (value) => {
  dialog.operationType = 1;
  console.log('edit data', value);
  setData(value);
  console.log('formData -> index', formData);
  dialog.visible = true;
};

const handleDeleteClick = (value) => {
  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(
      async () => {
        await apiDeleteById(value.id);
        toast("删除字典成功");
        handleQuery();
      },
      () => {
        ElMessage.info("已取消删除");
      }
    )
    .finally(() => {
      getData();
    });
};

const handleCellClick = (value) => {
  dictStore.setActive(value);
};

const handleDialogClose = () => {
  getData();
};

const handleChange = () => {
  getData();
};
</script>