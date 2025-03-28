<template>
  <div class="flex-block flex-layout">
    <div class="flex-auto">
      <el-table v-loading="props.loading" :data="props.tableData" style="width: 100%;" height="100%" v-bind="$attrs">
        <el-table-column v-for="column in props.columns" :key="column.prop" :prop="column.prop" :label="column.label"
          :width="column.width" v-bind="column?.attrs">
          <!-- 默认有插槽的情况 -->
          <template #default="scope">
            <slot v-if="column.slot" :name="column.prop" :scope="scope" />
            <!-- 当有告警时表格默认做出处理 -->
            <template v-if="column.tagList?.length">
              <div v-for="(tag, index) in filteredTagList(
        scope.row,
        column.tagList,
        column.prop
      )" :key="tag.value + '_' + index">
                <div class="tag" :class="tag.className">{{ tag.label }}</div>
              </div>
            </template>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 分页 -->
    <template v-if="props.usePagination">
      <div class="flex justify-between items-center flex-none">
        <div class="text-gray-500 pagination-counter">总共 {{ pageConfig.total }} 条信息</div>
        <el-pagination size="small" v-model:current-page="currentPage" class="cus-pagination" background
          layout="prev, pager, next" :page-size="pageConfig.pageSize" :total="pageConfig.total"
          @current-change="pageConfig.handleCurrentChange" />
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref } from "vue";
const props = defineProps({
  tableData: {
    type: Array,
    default: () => [],
  },
  columns: {
    type: Array,
    default: () => [],
  },
  pageConfig: {
    type: Object,
    default: () => ({
      pageSize: 5,
      total: 0,
      currentPage: 1,
    }),
  },
  loading: {
    type: Boolean,
    default: false,
  },
  usePagination: {
    type: Boolean,
    default: true,
  },
});

const filteredTagList = (scope, tagList, prop) =>
  tagList?.filter((tag) => tag?.value === scope?.[prop]);

const currentPage = ref(props.pageConfig.currentPage);

</script>
<style lang="scss" scoped>
.flex-block {
  height: 100%;
  width: 100%;
  overflow: auto;
  display: flex;
  flex-direction: column;
}

.flex-auto {
  flex: 1 1 auto;
  overflow: auto;
  position: relative;
}

.flex-none {
  flex: none;
  min-height: 32px;
}

.pagination-counter {
  font-family: PingFangSC, PingFang SC;
  font-weight: 400;
  font-size: 12px;
  color: #595959;
  line-height: 20px;
  text-align: left;
  font-style: normal;
}
</style>