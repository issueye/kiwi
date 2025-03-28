<template>
    <el-dialog v-model="dialogVisible" :title="dialog.title" width="500px" @close="onClose" @opened="onOpen">
        <el-form ref="dataFormRef" :model="dataForm" :rules="computedRules" label-width="auto">
            <el-form-item label="账户名" prop="username">
                <el-input v-model="dataForm.username" placeholder="请输入账户名称" />
            </el-form-item>

            <el-form-item label="昵称" prop="nick_name">
                <el-input v-model="dataForm.nick_name" placeholder="请输入昵称" />
            </el-form-item>

            <el-form-item label="角色" prop="role_code">
                <el-select v-model="dataForm.role_code" placeholder="请选择角色">
                    <el-option v-for="(item, index) in roleList" :key="index" :value="item.code" :label="item.name" />
                </el-select>
            </el-form-item>

            <el-form-item label="备注">
                <el-input v-model="dataForm.remark" type="textarea" placeholder="请输入备注" />
            </el-form-item>
        </el-form>

        <template #footer>
            <div class="dialog-footer">
                <el-button type="primary" @click="onSubmitClick">确 定</el-button>
                <el-button @click="onClose">取 消</el-button>
            </div>
        </template>
    </el-dialog>
</template>
<script setup>
import {
    apiAddUser,
    apiUpdateUser,
} from "~/api/user";

import { storeToRefs } from 'pinia';
import { ref, toRefs, computed } from 'vue';
import { useSystemStore } from '~/store/system';
import { toast } from "~/composables/util";

const props = defineProps({
    dialog: {
        type: Object,
        default: () => {
            return {
                title: '新增用户信息',
                operationType: 0,
            }
        },
    },

    visible: {
        type: Boolean,
        default: false,
    },

    dataForm: {
        type: Object,
        default: () => {
            return {
                id: 0,
                username: "",
                nick_name: "",
                role_code: "",
                remark: "",
            }
        }
    }
})

const { dialog, dataForm, visible } = toRefs(props)

const emits = defineEmits(['update:visible', 'close'])

const sysStore = useSystemStore()
const { roleList } = storeToRefs(sysStore);
const loading = ref(false);

const dataFormRef = ref();

const dialogVisible = computed({
    get() {
        return visible.value;
    },
    set(value) {
        emits('update:visible', value)
    }
})

const onOpen = () => {
    // console.log('onOpen -> dialog', dialog.value);
}

/**
 * 表单验证规则
 */
const computedRules = computed(() => {
    const rules = {
        username: [{ required: true, message: "请输入账户名称", trigger: "blur" }],
        nick_name: [{ required: true, message: "请输入昵称", trigger: "blur" }],
    };
    return rules;
});

/**
 * 提交表单按钮
 */
const onSubmitClick = () => {
    dataFormRef.value.validate(async (isValid) => {
        if (isValid) {
            switch (dialog.value.operationType) {
                case 0:
                    await addData();
                    break;
                case 1:
                    await editData();
                    break;
            }
        }
    });
};

/**
 * 新增数据
 */
const addData = async () => {
    try {
        loading.value = true;
        await apiAddUser(dataForm.value);
        loading.value = false;
        dialog.visible = false;
        toast("添加用户成功");
        emits('close');
    } catch (error) {
        loading.value = false;
    }
};

/**
 * 修改数据
 */
const editData = async () => {
    try {
        loading.value = true;
        await apiUpdateUser(dataForm.value);
        loading.value = false;
        dialog.visible = false;
        toast("修改用户成功");
        emits('close');
    } catch (error) {
        loading.value = false;
    }
};

// 关闭弹窗
const onClose = () => {
    dataFormRef.value.resetFields();
    dataFormRef.value.clearValidate();
    dataForm.value.id = undefined;

    emits('close')
};

</script>