<template>
  <el-card shadow="hover" class="h-full">
    <template #header>
      <div class="flex justify-between items-center">
        <span class="text-lg font-bold">项目统计</span>
        <el-icon><TrendCharts /></el-icon>
      </div>
    </template>
    <div class="grid grid-cols-2 gap-4">
      <div
        v-for="(item, index) in stats"
        :key="item.label"
        class="text-center p-3 bg-gray-50 rounded"
      >
        <div class="flex items-center justify-between">
          <div class="flex-1">
            <div class="text-xl font-bold text-primary">{{ item.value }}</div>
            <div class="text-gray-500 text-xs mt-0">{{ item.label }}</div>
          </div>
          <div class="w-32 h-12" :id="'chart-' + index"></div>
        </div>
      </div>
    </div>
  </el-card>
</template>

<script setup>
import { ref, onMounted } from "vue";
import * as echarts from "echarts";
const stats = ref([
  {
    label: "总项目数",
    value: 12,
    data: [5, 8, 10, 12],
  },
  {
    label: "进行中项目",
    value: 8,
    data: [3, 6, 7, 8],
  },
  {
    label: "已完成项目",
    value: 3,
    data: [1, 2, 2, 3],
  },
  {
    label: "逾期项目",
    value: 1,
    data: [0, 0, 1, 1],
  },
]);

const generateChart = (dom, data) => {
  const chart = echarts.init(dom);
  chart.setOption({
    xAxis: { show: false },
    yAxis: { show: false },
    grid: { top: 5, bottom: 5, left: 5, right: 5 },
    series: [
      {
        type: "line",
        data: data,
        smooth: true,
        symbol: "none",
        lineStyle: {
          width: 2,
          color: "#409EFF",
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: "rgba(64,158,255,0.4)" },
            { offset: 1, color: "rgba(64,158,255,0.05)" },
          ]),
        },
      },
    ],
  });
  return chart;
};

onMounted(() => {
  stats.value.forEach((item, index) => {
    const chartDom = document.getElementById(`chart-${index}`);
    if (chartDom) {
      generateChart(chartDom, item.data);
    }
  });
});
</script>