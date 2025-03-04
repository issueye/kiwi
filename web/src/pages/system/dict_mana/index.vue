<template>
  <base-page title="字典管理" desc="字典管理" padding="0px">
    <template #content>
      <div class="h-full flex">
        <div style="width: 500px; border-right: 1px solid #d9d9d9">
          <div class="p-4">
            <Llist />
          </div>
        </div>
        <div style="width: calc(100% - 500px)">
          <div class="p-4">
            <!-- <RTable /> -->
            <Rtable />
          </div>
        </div>
      </div>
    </template>
  </base-page>
</template>
  
  <script setup>
import { onMounted } from "vue";
import { apiGetDictsList } from "~/api/dict";
import Llist from "./components/l_list.vue";
import Rtable from "./components/r_table.vue";

//   import LTable from "./l_table/index.vue";
//   import RTable from "./r_table/index.vue";
onMounted(() => {
  getData();
});

const getData = async () => {
  let res = await apiGetDictsList();
  console.log("res", res);
};
</script>