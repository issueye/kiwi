import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import { toast } from "~/composables/util";

import { apiGetList as apiGetProjectList } from "~/api/project";
import {
    apiGetList as apiGetVersionList,
    apiAddData as apiAddVersionData,
    apiUpdateData as apiUpdateVersionData,
    apiDeleteById as apiDeleteVersionDataById,
} from "~/api/project/version";

import { apiGetList as apiGetBranchList } from "~/api/project/branch";

export const useProjectStore = defineStore(
    'project',
    () => {
        // 当前选中的项目
        const indexProject = ref(null)
        // 当前选中的版本
        const indexVersion = ref(null)
        // 版本列表
        const versionList = ref([])
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

        // 版本分页配置
        const versionPageConfig = reactive({
            pageSize: 10,
            currentPage: 1,
            total: 0,
        });

        // 项目查询参数
        const projectQryParams = reactive({
            keywords: "",
        })

        // 版本查询参数
        const versionQryParams = reactive({
            keywords: "",
            project_id: 0,
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
            projectList.value = res.list;
            projectPageConfig.total = res.total;
        }

        /**
         * 设置项目
         * @param {*} project 
         */
        const setProject = async (project) => {
            indexProject.value = project;

            // 切换项目时，清空版本列表
            versionList.value = [];
            // 切换项目时，清空版本
            indexVersion.value = null;

            // 切换项目时，重新获取版本列表
            versionQryParams.project_id = indexProject.value.id;
            let params = {
                pageNum: versionPageConfig.currentPage,
                pageSize: versionPageConfig.pageSize,
                condition: versionQryParams,
            };
            const res = await apiGetVersionList(params);
            versionList.value = res.list;
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

        /**
         * 获取版本列表
         */
        const getVersionList = async () => {
            versionQryParams.project_id = indexProject.value.id;
            let params = {
                pageNum: versionPageConfig.currentPage,
                pageSize: versionPageConfig.pageSize,
                condition: versionQryParams,
            };

            const res = await apiGetVersionList(params);
            versionList.value = res.list;
            versionPageConfig.total = res.total;
        }

        /**
         * 添加版本
         * @param {*} data
         */
        const addVersionData = async (data) => {
            try {
                data.project_id = indexProject.value.id;
                await apiAddVersionData(data);
                toast("添加版本信息成功");
                getVersionList();
            } catch (error) {
                toast("添加版本信息失败", "error");
            }
        }

        /**
         * 更新版本
         * @param {*} data
         */
        const updateVersionData = async (data) => {
            try {
                await apiUpdateVersionData(data);
                toast("更新版本信息成功");
                getVersionList();
            } catch (error) {
                toast("更新版本信息失败", "error");
            }
        }

        const deleteVersionById = async (id) => {
            try {
                await apiDeleteVersionDataById(id);
                toast("删除版本信息成功");
                getVersionList();
            } catch (error) {
                toast("删除版本信息失败", "error");
            }
        }

        /**
         * 获取分支列表
         */
        const getBranchList = async (id) => {
            try {
                let params = {
                    pageNum: 0,
                    pageSize: 0,
                    condition: {
                        project_id: 0,
                    },
                };

                if (id) {
                    params.condition.project_id = id;
                } else {
                    params.condition.project_id = indexProject.value.id;
                }

                const res = await apiGetBranchList(params);
                branchList.value = res.list;
            } catch (error) {
                toast("获取分支列表失败", "error");
            }
        }

        return {
            indexProject,
            indexVersion,
            versionList,
            projectList,
            projectPageConfig,
            versionPageConfig,
            projectQryParams,
            versionQryParams,
            branchList,

            getProjectList,
            getProjectById,
            getVersionList,
            setProject,
            addVersionData,
            updateVersionData,
            deleteVersionById,
            getBranchList,
        }
    },
    { persist: true }
)