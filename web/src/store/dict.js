import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'

import { apiGetDictsList, apiDetails } from "~/api/dict";

export const useDictStore = defineStore(
    'dict',
    () => {
        const activeData = ref(null)

        const dictCondition = reactive({
            condition: '',
        })

        const detailCondition = reactive({
            condition: '',
        })

        const pageConfig = reactive({
            pageSize: 10,
            currentPage: 1,
            total: 0,
            handleCurrentChange: (val) => {
                pageConfig.currentPage = val;
                getDetailData();
            }
        });

        const dictData = ref([])

        const dictDetailData = ref([])

        const getDictData = async () => {
            let params = {
                pageSize: 0,
                pageNum: 0,
                condition: {
                    keywords: dictCondition.condition,
                }
            }

            const res = await apiGetDictsList(params)
            dictData.value = res.list;


            setActive(activeData.value);
        }

        const setActive = async (value) => {
            activeData.value = value;
            if (!activeData.value) return;

            getDetailData(value);
        }

        const getDetailData = async () => {
            let params = {
                pageSize: pageConfig.pageSize,
                pageNum: pageConfig.currentPage,
                condition: {
                    keywords: detailCondition.condition,
                    dict_code: activeData.value.code,
                }
            }

            const res = await apiDetails(params);
            dictDetailData.value = res.list;
            pageConfig.total = res.total;
        }

        return {
            activeData,
            dictCondition,
            detailCondition,
            dictData,
            dictDetailData,
            pageConfig,
            getDictData,
            setActive,
        }
    },
    {
        persist: true
    }
)