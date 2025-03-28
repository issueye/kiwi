import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import { toast } from "~/composables/util";

import { apiGetList as apiGetProjectList } from "~/api/project";
import { apiGetList as apiGetBranchList } from "~/api/project/branch";

export const useProjectStore = defineStore(
    'project',
    () => {
        // 项目列表
        const projectList = ref([])
        // 分支列表
        const branchList = ref([])

        // 项目分页配置
        const projectPageConfig = reactive({
            pageSize: 10,
            currentPage: 1,
            total: 0,
        });

        // 项目查询参数
        const projectQryParams = reactive({
            keywords: "",
        })

        /**
         * 获取项目列表
         */
        const getProjectList = async () => {
            let params = {
                pageNum: projectPageConfig.currentPage,
                pageSize: projectPageConfig.pageSize,
                condition: projectQryParams,
            };

            const res = await apiGetProjectList(params);
            projectList.value = res.list.map((item) => {
                return {
                    ...item,
                    robots: item.project_robots.map((r) => r.id)
                }
            });
            projectPageConfig.total = res.total;
        }

        /**
         * 根据id获取项目
         * @param {*} id
         */
        const getProjectById = (id) => {
            let data = [...projectList.value];
            for (let i = 0; i < data.length; i++) {
                if (data[i].id == id) {
                    return data[i];
                }
            }

            return null;
        }

        return {
            projectList,
            projectPageConfig,
            projectQryParams,
            branchList,

            getProjectList,
            getProjectById,
        }
    },
    { persist: true }
)