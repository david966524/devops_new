<template>
    <div>
        <div class="gva-table-box">
            <el-table ref="multipleTable" :data="tableData" style="width: 100%" tooltip-effect="dark" row-key="ID">
                <el-table-column type="selection" width="55" />
                <el-table-column align="left" label="状态" prop="color" width="150">
                    <template #default="scope">
                        <el-tag :type="scope.row.color === 'blue' ? 'success' : 'danger'" disable-transitions>{{
                            scope.row.color
                            }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="任务名" prop="name" width="250" />
                <el-table-column align="left" label="job Url" prop="url" width="500">
                    <template #default="{ row }">
                        <a :href="row.url" target="_blank">{{ row.url }}</a>
                    </template>
                </el-table-column>
                <el-table-column align="right" label="按钮组" min-width="160">
                    <template #default="scope">
                        <el-button type="primary" link icon="View" @click="showBuildHistory(scope.row)">构建历史记录</el-button>
                        <el-button type="primary" link icon="edit" @click="buildJob(scope.row)">构建新版本</el-button>
                        <!-- <el-popover v-model="scope.row.visible" placement="top" width="160">
                            <p>确定要删除吗？</p>
                            <div style="text-align: right; margin-top: 8px;">
                                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                                <el-button type="primary" @click="">确定</el-button>
                            </div>
                            <template #reference>
                                <el-button type="primary" link icon="delete"
                                    @click="scope.row.visible = true">删除</el-button>
                            </template>
                        </el-popover> -->
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <el-dialog v-model="dialogTableVisible" :before-close="closeDialog" title="" width="50%">
            <el-table :data="buildtable" :edit-config="{ editMethod: 'cell' }">
                <el-table-column property="Number" label="构建 ID" width="150" />
                <el-table-column property="URL" label="构建 Url" width="450" />
                <el-table-column fixed="right" label="Operations">
                    <template #default="scope">
                        <el-button link type="primary" icon="View" size="small"
                            @click="showBuildInfo(scope.row)">查看构建信息</el-button>
                        <el-button link type="primary" icon="View" size="small"
                            @click="showBuildConsoleOut(scope.row)">查看控制台输出</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-dialog>

        <el-dialog v-model="dialogbuildinfoVisible" :before-close="closebuildDialog" title="" width="50%">
            <div>
                <div>构建结果：
                    <el-tag :type="buildinfo.result === 'SUCCESS' ? 'success' : 'danger'" disable-transitions>{{
                        buildinfo.result
                        }}</el-tag>
                </div>
                <div>构建地址 Url：<el-link :href=buildinfo.url target="_blank" type="primary">{{ buildinfo.url }}</el-link>
                </div>
                <el-divider border-style="double" />
                <ul>
                    Comment 记录 ：
                    <li v-for="item in buildinfo.changeSet.items" :key="item">
                        <div> Commit记录： {{ item.comment }}</div>
                        <div> Commit用户： {{ item.author.fullName }}</div>
                        <el-divider border-style="dashed" />
                    </li>
                </ul>
            </div>
        </el-dialog>

        <el-dialog v-model="dialogconsoleOutVisible" :before-close="closebuildDialog" title="" width="50%">
            <div>
                <p class="break-words">{{ consoleOut }}</p>
            </div>
        </el-dialog>

        <el-dialog v-model="dialogbuildVisible" :before-close="closebuildDialog" title="" width="50%">
            <div>
                <!-- {{ parameters }} -->
                <li v-for="item in parameters" :key="item">
                    <div>
                        <p class="break-words">{{ item.description }}</p>
                    </div>
                    <el-divider border-style="dashed" />
                    <el-form :inline="true" :model="item.defaultParameterValue" class="demo-form-inline">
                        <el-form-item :label="item.defaultParameterValue.name">
                            <el-input v-model="item.defaultParameterValue.value" placeholder="" />
                        </el-form-item>
                        <el-form-item>
                            <el-button type="primary" @click="buildParameter(item.defaultParameterValue.name,item.defaultParameterValue.value)">构建新版本</el-button>
                        </el-form-item>
                    </el-form>
                </li>
            </div>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { getjobs, getBuildHistory, getBuildInfo, getConsoleOut, getBuildParameters, onlyBuildJob ,buildJobParameter} from "../../../api/jenkinsapi"
import { ref } from "vue"
import { ElMessage } from 'element-plus'

const tableData = ref([])

const getJobs = async () => {
    const data = await getjobs()
    console.log(data.data)
    tableData.value = data.data
}
const dialogTableVisible = ref(false)
const dialogbuildinfoVisible = ref(false)
const closeDialog = ref(false)
const closebuildDialog = ref(false)
const dialogconsoleOutVisible = ref(false)
const dialogbuildVisible = ref(false)
const consoleOut = ref()
getJobs()

const buildtable = ref([])
const buildinfo = ref()
const jobname = ref()
const parameters = ref()
const showBuildHistory = async (row) => {
    jobname.value = row.name
    dialogTableVisible.value = true
    const data = await getBuildHistory({ Name: row.name })
    console.log(data.data)
    buildtable.value = data.data
}

const showBuildInfo = async (row) => {

    const data = await getBuildInfo({
        Name: jobname.value,
        BuildId: row.Number
    })
    dialogbuildinfoVisible.value = true
    buildinfo.value = data.data
}

const buildJob = async (row) => {
    //Name: row.Name,
    const data = await getBuildParameters({
        Name: row.name,
    })
    //有构建参数 开启dilog  选参数构建
    if (data.data !== null) {
        console.log("参数化构建")
        dialogbuildVisible.value = true
        parameters.value = data.data
        jobname.value = row.name
        return
    }
    //没有构建参数 直接构建
    const req: any = await onlyBuildJob({
        Name: row.name
    })
    if (req.code === 0) {
        ElMessage({
            message: req.msg,
            type: 'success',
        })
    }
}

const buildParameter = async (name,value) => {

    console.log(name,value)
    const data:any = await buildJobParameter({
        name: jobname.value,
        paraName: name,
        paraValue: value
    })
    if (data.code === 0){
        closebuildDialog.value=false
        ElMessage({
            message: data.msg,
            type: 'success',
        })
    }
}

const showBuildConsoleOut = async (row) => {
    dialogconsoleOutVisible.value = true
    const data = await getConsoleOut({
        Name: jobname.value,
        BuildId: row.Number
    })
    consoleOut.value = data.data
}
</script>

<script  lang="ts">
export default {
    name: 'jenkins'
}
</script>

<style scoped>
.break-words {
    white-space: pre-wrap;
}
</style>