<template>
  <div class="min-h-screen bg-primary-500 flex items-center justify-center">
    <div
      class="w-full max-w-4xl mx-auto flex shadow-2xl rounded-lg overflow-hidden"
    >
      <!-- 左侧图片区域 -->
      <div class="hidden lg:block lg:w-2/5 relative bg-primary-100">
        <el-image
          :src="SLImage"
          class="w-full h-full object-cover"
          fit="cover"
        />
        <div
          class="absolute bottom-0 left-0 right-0 p-4 bg-gradient-to-t from-black/50 to-transparent text-white"
        >
          <h3 class="text-xl font-bold">Welcome Back</h3>
          <p class="mt-1 text-xs">欢迎使用本系统</p>
        </div>
      </div>

      <!-- 右侧登录表单区域 -->
      <div class="w-full lg:w-3/5 bg-white p-6 lg:p-8">
        <h2 class="font-bold text-2xl text-gray-800 text-center mb-6">
          欢迎回来
        </h2>

        <div
          class="flex items-center justify-center my-4 text-gray-400 space-x-2"
        >
          <span class="h-[1px] w-12 bg-gray-200"></span>
          <span class="text-sm">账号密码登录</span>
          <span class="h-[1px] w-12 bg-gray-200"></span>
        </div>

        <el-form ref="formRef" :rules="rules" :model="form" class="space-y-4">
          <el-form-item prop="username" class="w-full">
            <el-input
              v-model="form.username"
              placeholder="请输入用户名"
              size="default"
              class="login-input"
            >
              <template #prefix>
                <el-icon class="text-gray-400"><user /></el-icon>
              </template>
            </el-input>
          </el-form-item>

          <el-form-item prop="password" class="w-full">
            <el-input
              type="password"
              v-model="form.password"
              placeholder="请输入密码"
              show-password
              size="large"
              class="login-input"
            >
              <template #prefix>
                <el-icon class="text-gray-400"><lock /></el-icon>
              </template>
            </el-input>
          </el-form-item>

          <div class="flex items-center justify-between">
            <el-checkbox v-model="form.remember">记住密码</el-checkbox>
            <el-link type="primary" class="text-sm">忘记密码？</el-link>
          </div>

          <el-button
            type="primary"
            class="w-full h-12 text-lg font-medium"
            :loading="loading"
            @click="onSubmit"
          >
            登录
          </el-button>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onBeforeUnmount } from "vue";
import { useRouter } from "vue-router";
import { useUserStore } from "~/store";
import { toast } from "~/composables/util";
import SLImage from "~/assets/sl.png";

const router = useRouter();
const userStore = useUserStore();

const form = reactive({
  username: "",
  password: "",
  remember: false,
});

const rules = {
  username: [{ required: true, message: "用户名不能为空", trigger: "blur" }],
  password: [{ required: true, message: "密码不能为空", trigger: "blur" }],
};

const formRef = ref(null);
const loading = ref(false);

const onSubmit = () => {
  formRef.value.validate(async (valid) => {
    if (!valid) return false;
    loading.value = true;
    try {
      await userStore.loginUser(form);
      toast("登录成功");
      router.push("/home");
    } catch (err) {
      console.error(err);
    } finally {
      loading.value = false;
    }
  });
};

function onKeyUp(e) {
  if (e.key === "Enter") onSubmit();
}

onMounted(() => {
  document.addEventListener("keyup", onKeyUp);
});

onBeforeUnmount(() => {
  document.removeEventListener("keyup", onKeyUp);
});
</script>

<style scoped>
.login-input :deep(.el-input__wrapper) {
  padding: 8px 15px;
}

.login-input :deep(.el-input__inner) {
  height: 42px;
}

:deep(.el-button--primary) {
  background: linear-gradient(to right, #4481eb, #04befe);
  border: none;
  transition: all 0.3s ease;
}

:deep(.el-button--primary:hover) {
  transform: translateY(-1px);
  box-shadow: 0 7px 14px rgba(50, 50, 93, 0.1);
}
</style>