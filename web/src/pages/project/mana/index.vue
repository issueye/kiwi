<template>
  <base-page title="项目管理" desc="项目管理" padding="0px">
    <template #content>
      <div class="flex" style="height: calc(100% - 40px);">
        <div class="h-full p-5" style="width: 600px; border-right: 1px solid #d9d9d9;">
          <Left />
        </div>
        <div class="h-full p-5" style="width: calc(100% - 680px);">
          <Right />
        </div>
      </div>
    </template>
  </base-page>
</template>

<script setup>
import Left from './left/index.vue';
import Right from './right/index.vue';

</script>