<template>
  <div
    class="flex items-center bg-primary-500 text-light-50 fixed top-0 left-0 right-0 z-[999]"
    :style="{ height: global.CARAMBOLA_HEADER_HEIGHT }"
  >
    <span
      class="w-[250px] flex justify-center items-center text-xl font-thin text-white"
    >
      <!-- <el-icon class="mr-1">
        <component :is="global.APP_ICON"></component>
      </el-icon> -->
      <el-image :src="LogoImage" class="mr-2 w-8" />
      <span class="font-semibold subpixel-antialiased tracking-widest"
        >{{ global.APP_NAME }}
      </span>
    </span>

    <!-- 这个代码创建了一个带有“刷新”功能的图标按钮，当用户将鼠标悬停在按钮上时，会显示提示文字“刷新”，并且按钮在被点击时会触发 handleRefresh 方法，通常用于刷新页面或数据。 -->
    <el-tooltip effect="dark" content="刷新" placement="bottom">
      <el-icon
        class="w-[42px] flex justify-center items-center text-xl font-thin text-white hover:bg-primary-600"
        :style="{ height: global.CARAMBOLA_HEADER_HEIGHT }"
        @click="handleRefresh"
      >
        <refresh />
      </el-icon>
    </el-tooltip>

    <!-- 菜单 -->
    <div class="grow h-full ml-5 flex items-end">
      <f-menu />
    </div>

    <!-- 这个代码创建了一个带有“全屏”功能的图标按钮，当用户将鼠标悬停在按钮上时，会显示提示文字“全屏”，并且按钮在被点击时会触发 toggle 方法，通常用于全屏显示或退出全屏。 -->
    <div class="flex items-center">
      <el-tooltip effect="dark" content="全屏" placement="bottom">
        <el-icon
          class="w-[42px] flex justify-center items-center text-xl font-thin text-white hover:bg-primary-600 mr-1"
          :style="{ height: global.CARAMBOLA_HEADER_HEIGHT }"
          @click="toggle"
        >
          <full-screen v-if="!isFullscreen" />
          <aim v-else />
        </el-icon>
      </el-tooltip>
      <el-tooltip effect="dark" content="主题" placement="bottom">
        <el-icon
          class="w-[42px] flex justify-center items-center text-xl font-thin text-white hover:bg-primary-600 mr-12"
          :style="{ height: global.CARAMBOLA_HEADER_HEIGHT }"
          @click="openThemeColorPicker"
        >
          <brush />
        </el-icon>
      </el-tooltip>
      <theme-color-picker ref="themeColorPickerRef" />
      <el-dropdown class="dropdown mr-[10px]" @command="handleCommand">
        <span class="flex items-center text-white text-light-50">
          <el-avatar class="mr-2" :size="25" :src="userStore.user.avatar" />
          {{ userStore.user.username }}
          <el-icon class="el-icon--right">
            <arrow-down />
          </el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="updateAvatar">更换头像</el-dropdown-item>
            <el-dropdown-item command="rePassword">修改密码</el-dropdown-item>
            <el-dropdown-item command="logout">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>

  <!-- 更换头像的表单抽屉 -->
  <form-drawer
    ref="avatarDrawerRef"
    title="更换头像"
    destroyOnClose
    @submit="onAvatarSubmit"
  >
    <el-form
      ref="avatarFormRef"
      :model="avatarForm"
      label-width="80px"
      size="small"
    >
      <el-form-item label="选择头像">
        <el-upload
          :file-list="avatarFileList"
          :auto-upload="false"
          :on-change="handleAvatarChange"
          list-type="picture-card"
          accept="image/*"
        >
          <i class="el-icon-plus"></i>
        </el-upload>
      </el-form-item>
    </el-form>
  </form-drawer>

  <!-- 修改密码的表单抽屉（保持不变） -->
  <form-drawer
    ref="formDrawerRef"
    title="修改密码"
    destroyOnClose
    @submit="onSubmit"
  >
    <el-form ref="formRef" :rules="rules" :model="form" label-width="80px">
      <el-form-item prop="oldpassword" label="旧密码">
        <el-input
          v-model="form.oldpassword"
          placeholder="请输入旧密码"
        ></el-input>
      </el-form-item>
      <el-form-item prop="password" label="新密码">
        <el-input
          type="password"
          v-model="form.password"
          placeholder="请输入新密码"
          show-password
        ></el-input>
      </el-form-item>
      <el-form-item prop="repassword" label="确认密码">
        <el-input
          type="password"
          v-model="form.repassword"
          placeholder="请再次输入新密码"
          show-password
        ></el-input>
      </el-form-item>
    </el-form>
  </form-drawer>
</template>

<script setup>
import FormDrawer from "~/components/FormDrawer.vue";
import { ref, reactive } from "vue";
import { useFullscreen } from "@vueuse/core";
import { useRepassword, useLogout } from "~/composables/useManager";
import { useUserStore } from "~/store";
import { toast } from "~/composables/util";
import { uploadAvatar, updateUserInfo } from "~/api/user"; // 新增的 API 函数
import { global } from "~/init/global";
import LogoImage from "~/assets/logo.ico";
import FMenu from "./FMenu.vue";
import ThemeColorPicker from "~/components/ThemeColorPicker.vue";
const userStore = useUserStore();

const { isFullscreen, toggle } = useFullscreen();

const { formDrawerRef, form, rules, formRef, onSubmit, openRePasswordForm } =
  useRepassword();

const { handleLogout } = useLogout();

// 处理用户下拉菜单的命令
const handleCommand = (command) => {
  switch (command) {
    case "updateAvatar":
      openAvatarDialog();
      break;
    case "rePassword":
      openRePasswordForm();
      break;
    case "logout":
      handleLogout();
      break;
  }
};

// 刷新页面
const handleRefresh = () => location.reload();
const themeColorPickerRef = ref(null);
const openThemeColorPicker = () => {
  themeColorPickerRef.value.open();
};
// 更换头像相关逻辑
const avatarDrawerRef = ref(null);
const avatarFormRef = ref(null);
const avatarForm = reactive({
  avatar: "",
});
const avatarFileList = ref([]);

const openAvatarDialog = () => {
  avatarDrawerRef.value.open();
};

const handleAvatarChange = (file, fileList) => {
  avatarFileList.value = fileList.slice(-1); // 只保留最新的文件
};

const onAvatarSubmit = async () => {
  if (avatarFileList.value.length === 0) {
    toast("请先选择一个头像图片", "error");
    return;
  }

  avatarDrawerRef.value.showLoading();

  try {
    // 1. 上传头像图片
    const formData = new FormData();
    formData.append("file", avatarFileList.value[0].raw);

    const uploadRes = await uploadAvatar(formData);
    const avatarUrl = uploadRes.url;

    // 2. 更新用户头像
    await updateUserInfo({
      avatar: avatarUrl,
      username: userStore.user.username,
      nick_name: userStore.user.nick_name,
    });

    // 3. 更新用户状态
    userStore.user.avatar = avatarUrl;

    // 4. 关闭抽屉并重置表单
    avatarDrawerRef.value.close();
    avatarFileList.value = [];

    toast("头像更新成功");
  } catch (error) {
    // 错误处理
    toast("头像更新失败", "error");
  } finally {
    avatarDrawerRef.value.hideLoading();
  }
};
</script>

<style scoped>
:deep(.el-dropdown-link:focus) {
  outline: none;
}
</style>
